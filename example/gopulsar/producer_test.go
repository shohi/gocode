package gopulsar

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/stretchr/testify/assert"
)

func TestProducer_Tenant(t *testing.T) {
	assert := assert.New(t)

	client := newTestClient(defaultClientOptions())
	defer client.Close()

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "persistent://tenant/hello/topic-1",
	})
	assert.Nil(err)
	assert.NotNil(producer)

	if producer != nil {
		producer.Close()
	}
}

// copied from `pulsar-client-go` example folder

func TestProducer_Key(t *testing.T) {
	assert := assert.New(t)

	client := newTestClient(defaultClientOptions())
	defer client.Close()

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "topic-1",
	})
	assert.Nil(err)
	defer producer.Close()

	ctx := context.Background()

	var key string
	for i := 0; i < 10; i++ {
		if i < 5 {
			key = "hello"
		} else {
			key = "world"
		}

		msgId, err := producer.Send(ctx, &pulsar.ProducerMessage{
			Key:     key,
			Payload: []byte(fmt.Sprintf("hello-%d", i)),
		})

		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Published message: ", msgId)
		}
	}
	consumer1, err := createConsumer(client, pulsar.ConsumerOptions{
		Name:             "1",
		Topic:            "topic-1",
		SubscriptionName: "my-sub",
		Type:             pulsar.KeyShared,
	})

	consumer2, err := createConsumer(client, pulsar.ConsumerOptions{
		Name:             "2",
		Topic:            "topic-1",
		SubscriptionName: "my-sub",
		Type:             pulsar.KeyShared,
	})

	assert.NotNil(consumer1)
	assert.NotNil(consumer2)

	time.Sleep(10 * time.Second)
}
