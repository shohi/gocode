package gopulsar

import (
	"fmt"
	"testing"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/stretchr/testify/assert"
)

func TestConsumer_MultipleClient(t *testing.T) {
	assert := assert.New(t)

	// clear subscription and populate
	topic := "multiple_client"
	subName := "my-sub"

	deleteTopic(topic)
	createSubscription(topic, subName, pulsar.Shared)
	publishOneMessage(topic, "msg-00", 20, false)

	c1 := newTestClient(defaultClientOptions())
	defer c1.Close()
	msgCh1 := make(chan pulsar.ConsumerMessage)
	s1, err := c1.Subscribe(pulsar.ConsumerOptions{
		Name:                "s1",
		Topic:               topic,
		Type:                pulsar.Shared,
		SubscriptionName:    subName,
		MessageChannel:      msgCh1,
		ReceiverQueueSize:   1,
		NackRedeliveryDelay: 100 * time.Millisecond,
	})
	assert.Nil(err)
	defer s1.Close()

	// 1. consumer 1
	go func() {
		last := time.Now()
		for msg := range s1.Chan() {
			gap := time.Now().Sub(last)
			last = time.Now()

			time.Sleep(5 * time.Second)

			s1.Ack(msg)

			fmt.Printf("====> s1, messsge: %v, redeliver: %v, elapse: %v\n",
				string(msg.Payload()), msg.RedeliveryCount(), gap,
			)
		}

	}()

	c2 := newTestClient(defaultClientOptions())
	defer c2.Close()

	s2, err := c2.Subscribe(pulsar.ConsumerOptions{
		Name:                "s2",
		Topic:               topic,
		Type:                pulsar.Shared,
		SubscriptionName:    subName,
		ReceiverQueueSize:   1,
		NackRedeliveryDelay: 100 * time.Millisecond,
	})
	assert.Nil(err)
	defer s2.Close()

	// 2. consumer 2
	go func() {
		last := time.Now()
		for msg := range s2.Chan() {
			gap := time.Now().Sub(last)
			last = time.Now()

			s2.Ack(msg)

			fmt.Printf("====> s2, messsge: %v, redeliver: %v, elapse: %v\n",
				string(msg.Payload()), msg.RedeliveryCount(), gap,
			)
		}
	}()

	time.Sleep(5 * time.Minute)
}
