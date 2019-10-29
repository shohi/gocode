package nats_test

import (
	"fmt"
	"testing"
	"time"

	stan "github.com/nats-io/nats-streaming.go"
	nats "github.com/nats-io/nats.go"
)

func TestMessenger_STAN_Connect(t *testing.T) {
	t.Skip()
	// assert := assert.New(t)
	nc, _ := nats.Connect("nats://localhost:4223")
	for k := 0; k < 10; k++ {
		stan.Connect("no-exist-cluster-cluster",
			"test-client",
			// stan.NatsURL("nats://localhost:4224"),
			stan.NatsConn(nc),
			stan.ConnectWait(1*time.Second))
		// assert.NotNil(err)
	}

	fmt.Printf("=================> Connection Over\n")

	time.Sleep(10 * time.Minute)
}
