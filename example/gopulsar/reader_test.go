package gopulsar

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/stretchr/testify/assert"
)

func TestReader_HasNext(t *testing.T) {
	assert := assert.New(t)

	topic := "reader-topic"
	subName := "my-sub-0"
	prepareMessageForShared(topic, subName, "hello", 10)

	// read topic
	readTopic("reader", topic, pulsar.EarliestMessageID(), true, 0)

	// consume topic
	client := newTestClient(defaultClientOptions())
	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:             "s1",
		Topic:            topic,
		SubscriptionName: subName,
		Type:             pulsar.Shared,
	})
	assert.Nil(err)
	assert.NotNil(s)

	var firstUnacked pulsar.Message
	go func() {
		count := 0
		for m := range s.Chan() {
			fmt.Printf("consumer new message: %v\n",
				string(m.Payload()))

			if count%10 == 4 {
				if firstUnacked == nil {
					fmt.Printf("========> init unacked, index: %v\n", count)
					firstUnacked = m
				}
			} else {
				s.Ack(m)
			}
			count++
			if count == 10 {
				s.Close()
				fmt.Printf("======> consumer closed")
			}
		}
	}()

	// wait for consuming
	time.Sleep(3 * time.Second)

	fmt.Printf("=====> unacked: %v\n", firstUnacked.ID())

	// new consumer
	s2, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:  "s1",
		Topic: topic,
		// SubscriptionName: subName,
		SubscriptionName:            subName + "new",
		Type:                        pulsar.Shared,
		SubscriptionInitialPosition: pulsar.SubscriptionPositionEarliest,
	})
	assert.Nil(err)
	go func() {
		fmt.Printf(strings.Repeat("\n", 3))

		for m := range s2.Chan() {
			fmt.Printf("consumer 2 received message: %v\n", string(m.Payload()))
		}

	}()

	time.Sleep(1 * time.Second)

	// another reader
	// pulsar.LatestMessageID() can't get the last message.
	readTopic("reader2", topic, firstUnacked.ID(), true, 1)
}
