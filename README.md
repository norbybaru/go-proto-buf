GO PROTOCOL BUFFER

# Introduction to Protocol Buffers

## Prerequisites

* **recommended** `direnv`, to allow all Go-based binaries to be local to this folder and not installed globally. For more details you can refer to [this post](https://mariocarrion.com/2020/11/20/golang-go-tool-direnv.html).

## Tools

Install the following tools:

* **required** Protocol Buffers Compiler, `protoc` (to date version `25.0`):
    * Homebrew: `brew install protobuf`
* **required** `buf` CLI for linting and compiling:
    * `go install github.com/bufbuild/buf/cmd/buf@v1.28.0`
* **required** Protocol Buffer Plugin for Go:
    * `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1`
* **recommended** Code Formatting, `clang-format`, you can use `find . -name '*.proto' | xargs clang-format -i`
    * Homebrew: `brew install clang-format` (to date version `17.0.5`):

Command:

Help command
```bash
make help
```

Initialize Buf
```bash
buf mod init
```
