# README

使用Protobuf序列化样例

## Prerequisite

```bash
# install protobuf compiler
brew install protobuf

# install golang proto compiler plugin
# https://github.com/protocolbuffers/protobuf-go
go get google.golang.org/protobuf/cmd/protoc-gen-go
```

## Example

```bash
# compile proto to Golang-specific codes
protoc -I=$PWD --go_out=$PWD $PWD/message.proto

# Mashall/Unmarshal


```


## Links
1. Protocol Buffer Basics: Go, <https://developers.google.com/protocol-buffers/docs/gotutorial>

2. Protobuf终极教程, <https://colobu.com/2019/10/03/protobuf-ultimate-tutorial-in-go/>
