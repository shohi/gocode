package client

import (
	"log"
	"strings"

	"github.com/shohi/gocode/example/grpc/streaming/pb"
	"google.golang.org/grpc"
)

func Run(conf Config) {
	conn, err := grpc.Dial(conf.Addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	defer conn.Close()

	rc := pb.NewStreamServiceClient(conn)

	if strings.Contains(conf.Service, "list") {
		err = PrintLists(rc, &pb.StreamRequest{Pt: &pb.StreamPoint{
			Name:  "gRPC Stream Client: List",
			Value: 1000,
		}})
		if err != nil {
			log.Fatalf("printLists.err: %v", err)
		}
	}

	if strings.Contains(conf.Service, "record") {
		err = PrintRecord(rc, &pb.StreamRequest{Pt: &pb.StreamPoint{
			Name:  "gRPC Stream Client: Record",
			Value: 2000,
		}})
		if err != nil {
			log.Fatalf("printRecord.err: %v", err)
		}
	}

	if strings.Contains(conf.Service, "route") {
		err = PrintRoute(rc, &pb.StreamRequest{Pt: &pb.StreamPoint{
			Name:  "gRPC Stream Client: Route",
			Value: 3000,
		}})
		if err != nil {
			log.Fatalf("printRoute.err: %v", err)
		}
	}
}
