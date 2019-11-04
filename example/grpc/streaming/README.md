# Streaming gRPC


## Usage

```
# genernated RPC related go source code
cd pb/
protoc --go_out=plugins=grpc:. *.proto

# run server & client
cd cmd/
go run server/main.go
go run client/main.go

```
