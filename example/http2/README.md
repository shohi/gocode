# HTTP2 Example

## Usage

```
# 1. create self-signed CA cert and private key for test
openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt


# 2. run http2 server with TLS & test
go run tls/server.go
go run tls/client.go

# 3. run http2 server without TLS & test
go run h2c/server.go
go run h2c/client.go

```

## Reference

1. Go Http2和h2c, <https://colobu.com/2018/09/06/Go-http2-%E5%92%8C-h2c/>

2. HTTP/2 Adventure in the Go World, <https://posener.github.io/http2/>
