package nats

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/nats-io/go-nats"
	stan "github.com/nats-io/go-nats-streaming"
	"github.com/stretchr/testify/assert"
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

func createStanConn(opts stanOptions) (stan.Conn, error) {
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

func TestStan_Connection_Sub(t *testing.T) {
	assert := assert.New(t)

	opts := defaultStanOptions()
	opts.clientID = "test-sub"
	conn, err := createStanConn(opts)
	assert.Nil(err)

	conn.Subscribe("foo", func(m *stan.Msg) {
		log.Printf("=====> msg: %v", m)
	})

	time.Sleep(10 * time.Minute)
}

func TestStan_Connection_Pub(t *testing.T) {
	assert := assert.New(t)

	opts := defaultStanOptions()
	opts.clientID = "test-pub"
	conn, err := createStanConn(opts)
	assert.Nil(err)

	assert.Nil(err)
	log.Printf("connection, err: %v", err)

	// if stan server is down and restarted during the test
	// publish will fail forever and lostHandler will be called.
	for k := 0; k < 100; k++ {
		start := time.Now()
		msg := fmt.Sprintf("hello-%v", k)
		err = conn.Publish("foo", []byte(msg))
		log.Printf("publish-%v, elapsed: %v, err: %v", k, time.Since(start), err)
		time.Sleep(1 * time.Second)
	}

	time.Sleep(10 * time.Minute)
}
