/*
This program generates the protobuf and SteamLanguage files from the SteamKit data.
*/
package main

import (
	"encoding/csv"
	"flag"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var clean = flag.Bool("clean", true, "Clean output before generating")
var steamLang = flag.Bool("steamlang", true, "Generate Steam language files")
var protobufs = flag.Bool("proto", true, "Generate protobufs")
var verbose = flag.Bool("verbose", false, "Verbose output")

func main() {
	validInput := false
	flag.Parse()

	if *clean {
		cleanOutput()
		validInput = true
	}

	if *steamLang {
		buildSteamLanguage()
		validInput = true
	}

	if *protobufs {
		buildProto()
		validInput = true
	}

	if !validInput {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func cleanOutput() {
	print("# Cleaning")
	os.RemoveAll("../protocol/protobuf")

	os.Remove("../protocol/steamlang/enums.go")
	os.Remove("../protocol/steamlang/messages.go")
}

func buildSteamLanguage() {
	print("# Building Steam Language")
	execute("dotnet", "run", "-c", "release", "-p", "./GoSteamLanguageGenerator", "./SteamKit", "../protocol/steamlang")
	execute("gofmt", "-w", "../protocol/steamlang/enums.go", "../protocol/steamlang/messages.go")
}

func buildProto() {
	print("# Building Protobufs")

	buildProtobufs("protos.csv", "./Protobufs", "../", "./protocol/protobuf")
}
func buildProtobufs(srcCsv string, srcBaseDir string, outDir string, packagePrefix string) {
	file, err := os.Open(srcCsv)
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	tmpDir, err := ioutil.TempDir("", "generator-")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmpDir)

	err = os.MkdirAll(outDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	groupedRecords := make(map[string][][]string)

	print("# Preprocessing to: " + tmpDir)
	for _, r := range records[1:] {
		srcDir := r[0]
		srcFile := r[1]
		pkg := r[2]

		fullTmpPath := filepath.Join(tmpDir, srcDir, srcFile)
		fullSrcPath := filepath.Join(srcBaseDir, srcDir, srcFile)

		print("# Preprocessing: " + pkg + ", " + srcFile)
		preprocessProto(packagePrefix, pkg, fullSrcPath, fullTmpPath)

		if _, ok := groupedRecords[pkg]; !ok {
			groupedRecords[pkg] = make([][]string, 0)
		}

		groupedRecords[pkg] = append(groupedRecords[pkg], r)
	}

	for pkg, packageRecords := range groupedRecords {
		args := []string{"--go_out=" + outDir}

		print("# Building: " + pkg)
		for _, r := range packageRecords {
			srcDir := r[0]
			srcFile := r[1]

			filePath := filepath.Join(tmpDir, srcDir, srcFile)

			// protoc doesn't seem to have an issue with the same include path
			// being included many times
			args = append(args, "-I="+filepath.Join(tmpDir, srcDir))
			args = append(args, filePath)
		}

		execute("protoc", args...)
	}
}

// This regexp is used to fixup package references in the protobuf files.
// Any issues with this should be caught at generation time but this is definitely fragile.
// TODO: Use proper protobuf parser instead of fixup regexp.
var protobufPackageFixup = regexp.MustCompile(`((?:(?:repeated|optional|required)[\t\f ]+|(?:[\t\f ]+\())\.)`)

func preprocessProto(pkgPrefix, pkg, srcFile, outFile string) {
	err := os.MkdirAll(filepath.Dir(outFile), os.ModePerm)
	if err != nil {
		panic(err)
	}

	input, err := ioutil.ReadFile(srcFile)
	if err != nil {
		panic(err)
	}

	out, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	_, err = out.WriteString("syntax = \"proto2\";\n")
	if err != nil {
		panic(err)
	}

	_, err = out.WriteString("package " + pkg + ";\n")
	if err != nil {
		panic(err)
	}

	_, err = out.WriteString("option go_package = \"" + pkgPrefix + "/" + pkg + "\";\n")
	if err != nil {
		panic(err)
	}

	output := protobufPackageFixup.ReplaceAll(input, []byte("${1}"+pkg+"."))
	_, err = out.Write(output)
	if err != nil {
		panic(err)
	}
}

func print(text string) { os.Stdout.WriteString(text + "\n") }

func printerr(text string) { os.Stderr.WriteString(text + "\n") }

// This writer appends a "> " after every newline so that the outpout appears quoted.
type quotedWriter struct {
	w       io.Writer
	started bool
}

func newQuotedWriter(w io.Writer) *quotedWriter {
	return &quotedWriter{w, false}
}

func (w *quotedWriter) Write(p []byte) (n int, err error) {
	if !w.started {
		_, err = w.w.Write([]byte("> "))
		if err != nil {
			return n, err
		}
		w.started = true
	}

	for i, c := range p {
		if c == '\n' {
			nw, err := w.w.Write(p[n : i+1])
			n += nw
			if err != nil {
				return n, err
			}

			_, err = w.w.Write([]byte("> "))
			if err != nil {
				return n, err
			}
		}
	}
	if n != len(p) {
		nw, err := w.w.Write(p[n:])
		n += nw
		return n, err
	}
	return
}

func execute(command string, args ...string) {
	if *verbose {
		print(command + " " + strings.Join(args, " "))
	}
	cmd := exec.Command(command, args...)
	cmd.Stdout = newQuotedWriter(os.Stdout)
	cmd.Stderr = newQuotedWriter(os.Stderr)
	err := cmd.Run()
	if err != nil {
		printerr(err.Error())
		os.Exit(1)
	}
}
