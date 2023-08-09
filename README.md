# protoc-gen-go-json
Protobuf compiler plugin to generate Go JSON Marshal


## Install

```bash
go get github.com/The-night-elves/protoc-gen-go-json
```

Also requires [protoc](https://github.com/google/protobuf) and [protoc-gen-go](https://github.com/golang/protobuf) to be installed.

## Usage

```bash
protoc --go_out=. --plugin protoc-gen-go="${GOBIN}/protoc-gen-go" \
    --plugin=protoc-gen-go-json="${GOBIN}/protoc-gen-go-json" \
    --go-json_out=. --go-json_opt=config=FileNameSuffix=.json.go,config=EncodeMethodName=MarshalJSON \
    proto/{your protobuf name}.proto
```

### Options

The generator supports the following options which can be specified in the `--go-json_opt` parameter:
- FileNameSuffix string output file name suffix, default is `.json.go`
- EncodeMethodName string encode method name, default is `MarshalJSON`
- ImportWriter string import writer, default golang standard import `bytes`
  - warn ImportWriter need implement method `WriteByte(byte), WriteString(string), Write([]byte)`
- NewWriter string new writer, default bytes.`Buffer`, expr `var buf bytes.Buffer`, protogen.GoImportPath(`bytes`).Ident(`Buffer`)
- WriteBytes string write bytes, write bytes method name, default `buf.Bytes()`