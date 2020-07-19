package gopulsar

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/stretchr/testify/assert"
)

func TestCompaction_Producer(t *testing.T) {
	assert := assert.New(t)

	client := newTestClient(defaultClientOptions())

	compactTopic := "compacted-topic-1"

	p, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic:              compactTopic,
		Name:               "compact-p-1",
		DisableBatching:    true,
		MaxPendingMessages: 1,
	})
	assert.Nil(err)
	assert.NotNil(p)

	var payload [64]byte
	var ctx = context.Background()

	for i := 0; i < 10; i++ {
		_, err = p.Send(ctx, &pulsar.ProducerMessage{
			Payload: payload[:],
			Key:     "pulsar",
		})
		fmt.Printf("=====> publish message: %v\n", i)
		assert.Nil(err)
		time.Sleep(30 * time.Millisecond)
	}
}

func TestCompaction_Consumer(t *testing.T) {
	assert := assert.New(t)

	client := newTestClient(defaultClientOptions())

	compactTopic := "compacted-topic-1"

	c, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:             compactTopic,
		Name:              "compact-s-1",
		SubscriptionName:  "compact-sub-1",
		ReceiverQueueSize: 1,
		ReadCompacted:     true,
	})
	assert.Nil(err)
	assert.NotNil(c)

	for i := 0; i < 50; i++ {
		msg, err := c.Receive(context.Background())
		assert.Nil(err)
		fmt.Printf("===> receive mssage: %v - %v\n", i, msg.PublishTime())
	}
}
