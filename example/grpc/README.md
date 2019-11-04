# gRPC

## Requirements

```bash
# 1. Protocol Buffer 3
brew install protobuf

# 3. gPRC - The Go language implementation of gRPC. HTTP/2 based RPC
# if `golang.org` is unaccessible, add following replacement in project's go.mod:
#     replace google.golang.org/grpc => github.com/grpc/grpc-go v1.23.1
go get -u google.golang.org/grpc


# 3. install generator for go code
go get -u github.com/golang/protobuf/protoc-gen-go

```

## Reference

1. go-grpc-example, <https://book.eddycjy.com/golang/grpc/install.html>
