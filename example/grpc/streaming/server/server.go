package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/shohi/gocode/example/grpc/streaming/pb"
	"google.golang.org/grpc"
)

type Config struct {
	Port int
}

type StreamService struct {
}

func (s *StreamService) List(r *pb.StreamRequest, stream pb.StreamService_ListServer) error {
	log.Printf("Req==> %v", r)
	for n := 0; n <= 6; n++ {
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  r.Pt.Name,
				Value: r.Pt.Value + int32(n),
			},
		})
		time.Sleep(1 * time.Second)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *StreamService) Record(stream pb.StreamService_RecordServer) error {
	var maxVal int32 = -1
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.StreamResponse{Pt: &pb.StreamPoint{Name: "gRPC Stream Server: Record", Value: maxVal + 1}})
		}
		if err != nil {
			return err
		}

		if maxVal == -1 || maxVal < r.Pt.Value {
			maxVal = r.Pt.Value
		}

		log.Printf("stream.Req==> pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}
}

func (s *StreamService) Route(stream pb.StreamService_RouteServer) error {
	n := 0
	for {
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  "gPRC Stream Client: Route",
				Value: int32(n),
			},
		})
		if err != nil {
			return err
		}

		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		n++
		log.Printf("stream.Req==> pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)

		time.Sleep(1 * time.Second)
	}
}

func Run(conf Config) {
	server := grpc.NewServer()
	pb.RegisterStreamServiceServer(server, &StreamService{})

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	log.Printf("grpc server: %v", fmt.Sprintf("localhost:%v", conf.Port))

	server.Serve(l)
}
