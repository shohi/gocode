package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/shohi/gocode/example/grpc/plain/pb"
)

type Config struct {
	Addr string
}

func main() {
	conf := Config{Addr: ":9001"}
	conn, err := grpc.Dial(conf.Addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp==> %s", resp.GetResponse())
}
