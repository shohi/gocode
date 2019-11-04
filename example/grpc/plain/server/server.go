package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/shohi/gocode/example/grpc/plain/pb"
	"google.golang.org/grpc"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	log.Printf("Req==> %v", r.Request)
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

type Config struct {
	Port int
}

func Run(conf Config) {
	server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server, &SearchService{})

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	log.Printf("grpc server: %v", fmt.Sprintf("localhost:%v", conf.Port))

	server.Serve(l)
}
