package cmcst

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	pulsar "github.com/Comcast/pulsar-client-go"
	"github.com/stretchr/testify/assert"
)

func TestPulsarClient_Comcast(t *testing.T) {
	assert := assert.New(t)

	client, err := pulsar.NewClient(pulsar.ClientConfig{
		Addr:        "pulsar://localhost:6650",
		DialTimeout: 3 * time.Second,
	})
	client.Close()

	fmt.Println("client")

	if err != nil {
		panic(err)
	}

	// FIXME: library doesn't work with pulsar v2.5.2, out of date
	p, err := client.NewProducer(context.Background(), "hello", "producer_name")

	fmt.Println("producer")
	assert.Nil(err)

	ret, err := p.Send(context.Background(), []byte("hello"))
	assert.Nil(err)
	log.Println(ret)

}
