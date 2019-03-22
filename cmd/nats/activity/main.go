package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/go-nats"
	stan "github.com/nats-io/go-nats-streaming"
)

type stanOptions struct {
	clusterID   string
	clientID    string
	lostHandler stan.ConnectionLostHandler
}

func defaultStanOptions() stanOptions {
	return stanOptions{
		clusterID: "test-cluster",
		clientID:  "test-client",
		lostHandler: func(_ stan.Conn, err error) {
			log.Printf("======> connection lost, err: %v", err)
		},
	}
}

func newStanConn(opts stanOptions) (stan.Conn, error) {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		return nil, err
	}

	conn, err := stan.Connect(
		opts.clusterID,
		opts.clientID,
		stan.NatsConn(nc),
		stan.ConnectWait(2*time.Second),
		stan.Pings(2, 10),
		stan.SetConnectionLostHandler(opts.lostHandler),
	)

	return conn, err
}

func main() {
	channels := 50
	pause := 1 * time.Second

	signalCh := make(chan struct{}, 2*channels)

	for k := 0; k < channels; k++ {
		go publishPerChannel(k, pause, signalCh)
	}

	<-signalCh
}

func publishPerChannel(index int,
	pause time.Duration,
	signalCh chan<- struct{},
) {

	lostHandler := func(_ stan.Conn, err error) {
		log.Printf("======> connection lost, index: %v, err: %v", index, err)
		signalCh <- struct{}{}
	}
	opts := defaultStanOptions()
	opts.clientID = fmt.Sprintf("test-pub-%06d", index)
	opts.lostHandler = lostHandler
	conn, err := newStanConn(opts)

	if err != nil {
		log.Printf("======> connection err, index: %v, err: %v", index, err)
		signalCh <- struct{}{}
		return
	}
	subject := fmt.Sprintf("foo-%v", index)
	msg := make([]byte, 1024)
	deadline := time.Now().Add(24 * time.Hour)

	for {
		if time.Now().After(deadline) {
			signalCh <- struct{}{}
			break
		}

		err = conn.Publish(subject, msg)
		if err != nil {
			log.Printf("======> publish err, index: %v, err: %v", index, err)
			signalCh <- struct{}{}
			return
		}
		time.Sleep(pause)
	}

}
