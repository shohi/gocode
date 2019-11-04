package client

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/shohi/gocode/example/grpc/streaming/pb"
)

func PrintLists(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.List(context.Background(), r)
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp==> pt.name: [%s], pt.value: [%d]", resp.Pt.Name, resp.Pt.Value)
	}

	return nil
}

func PrintRecord(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.Record(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n < 6; n++ {
		err := stream.Send(r)
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
		r.Pt.Value++
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	log.Printf("resp==> pt.name: [%s], pt.value: [%d]", resp.Pt.Name, resp.Pt.Value)

	return nil
}

func PrintRoute(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.Route(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n <= 6; n++ {
		err = stream.Send(r)
		if err != nil {
			return err
		}

		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp==> pt.name: [%s], pt.value: [%d]", resp.Pt.Name, resp.Pt.Value)

		r.Pt.Value++

		time.Sleep(1 * time.Second)
	}

	stream.CloseSend()

	return nil
}
