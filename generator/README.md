# Generator for steamlang and protobuf

We generate Go code from SteamKit protocol descriptors, namely `steamlang` files and protocol buffer files.

## Dependencies
1. Get SteamKit and Protobufs submodule: `git submodule update --init --recursive`.
2. Install [`protoc`](https://developers.google.com/protocol-buffers/docs/downloads), the protocol buffer compiler.
3. Install `proco-gen-go`: `go get google.golang.org/protobuf/protoc-gen-go/`. Match this with the tag in the root go.mod.
4. Install the .NET Core SDK (3.1 or later).

## Execute generator

Execute `go run generator.go clean proto steamlang` to clean build files, then build protocol buffer files and then build steamlang files.
