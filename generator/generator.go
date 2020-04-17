/*
This program generates the protobuf and SteamLanguage files from the SteamKit data.
*/
package main

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var printCommands = true

func main() {
	args := strings.Join(os.Args[1:], " ")

	found := false
	if strings.Contains(args, "clean") {
		clean()
		found = true
	}
	if strings.Contains(args, "steamlang") {
		buildSteamLanguage()
		found = true
	}
	if strings.Contains(args, "proto") {
		buildProto()
		found = true
	}

	if !found {
		os.Stderr.WriteString("Invalid target!\nAvailable targets: clean, proto, steamlang\n")
		os.Exit(1)
	}
}

func clean() {
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

	print("# Preprocessing to: " + tmpDir)
	for _, r := range records[1:] {
		srcDir := r[0]
		srcFile := r[1]
		pkg := r[2]

		fullTmpPath := filepath.Join(tmpDir, srcDir, srcFile)
		fullSrcPath := filepath.Join(srcBaseDir, srcDir, srcFile)

		print("# Preprocessing: " + pkg + ", " + srcFile)
		preprocessProto(packagePrefix+"/"+pkg, fullSrcPath, fullTmpPath)
	}

	for _, r := range records[1:] {
		srcDir := r[0]
		srcFile := r[1]
		pkg := r[2]

		print("# Building: " + pkg + ", " + srcFile)
		compileProto(tmpDir, srcDir, srcFile, outDir)
	}
}

func compileProto(srcBase, srcSubdir, proto, outDir string) {
	err := os.MkdirAll(outDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	srcFile := filepath.Join(srcBase, srcSubdir, proto)

	execute("protoc", "--go_out="+outDir, "-I="+filepath.Join(srcBase, srcSubdir), "-I="+filepath.Join(srcBase, "steam"), srcFile)
}

func preprocessProto(pkg, srcFile, outFile string) {
	err := os.MkdirAll(filepath.Dir(outFile), os.ModePerm)
	if err != nil {
		panic(err)
	}

	in, err := os.Open(srcFile)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	_, err = io.WriteString(out, "syntax = \"proto2\";\n")
	if err != nil {
		panic(err)
	}

	_, err = io.WriteString(out, "option go_package = \""+pkg+"\";\n")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(out, in)
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
		nw, err := w.w.Write(p[n:len(p)])
		n += nw
		return n, err
	}
	return
}

func execute(command string, args ...string) {
	if printCommands {
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
