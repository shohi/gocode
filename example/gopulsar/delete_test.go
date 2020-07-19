package gopulsar

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/stretchr/testify/assert"
)

func TestMessageDeletion_Reader(t *testing.T) {
	assert := assert.New(t)

	// 0. populate messages
	topic := "msg-deleted"
	subName := "my-sub"
	prepareMessageForShared(topic, subName, "msg-00", 10)

	// 1. consuming all
	client := newTestClient(defaultClientOptions())
	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:             "s",
		Topic:            topic,
		SubscriptionName: subName,
	})
	assert.Nil(err)

	go func() {
		for m := range s.Chan() {
			s.Ack(m)
			fmt.Printf("====> consumer received %v, acked\n",
				string(m.Payload()))
		}
	}()

	// 2. reader keeps reading
	go func(topic string) {
		ticker := time.NewTicker(10 * time.Second)
		count := 0
		for range ticker.C {
			shouldContinue := readTopic("reader", topic, pulsar.EarliestMessageID(), true, count)
			if !shouldContinue {
				fmt.Printf("======> no message for reader, exit\n")
				return
			}
			count++
		}
	}(topic)

	// 3. close consumer explicitly
	time.Sleep(10 * time.Second)
	s.Close()
	log.Printf("=====> closed active consumer")

	// 4. delete subscription
	log.Println("=====> start deleting subscription after 20s.......")
	time.Sleep(20 * time.Second)
	log.Println("=====> deleting subscription.......")
	deleteSubscription(topic, subName)
	log.Println("=====> deleting subscription done .......")

	// 4. delete topic
	log.Println("=====> start deleting topic after 20s.......")
	time.Sleep(20 * time.Second)
	log.Println("=====> delete topic.......")
	deleteTopic(topic)
	log.Println("=====> deleting topic done .......")

	time.Sleep(5 * time.Minute)
}

func TestMessageDeletion_DeleteTopic(t *testing.T) {
	assert := assert.New(t)

	// 0. populate messages
	topic := "msg-deleted"
	subName := "my-sub"
	prepareMessageForShared(topic, subName, "msg-00", 10)

	// 1. consuming all
	client := newTestClient(defaultClientOptions())
	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:             "s",
		Topic:            topic,
		SubscriptionName: subName,
	})
	assert.Nil(err)

	go func() {
		for m := range s.Chan() {
			s.Ack(m)
			fmt.Printf("====> consumer received %v, acked\n",
				string(m.Payload()))
		}
	}()

	// 3. close consumer explicitly
	time.Sleep(10 * time.Second)
	readTopic("reader0", topic, pulsar.EarliestMessageID(), true, 0)

	s.Close()
	log.Printf("=====> closed active consumer")

	readTopic("reader1", topic, pulsar.EarliestMessageID(), true, 1)

	// 4. delete subscription
	log.Println("=====> start deleting subscription after 20s.......")
	time.Sleep(20 * time.Second)
	log.Println("=====> deleting subscription.......")
	deleteSubscription(topic, subName)
	log.Println("=====> deleting subscription done .......")

	readTopic("reader2", topic, pulsar.EarliestMessageID(), true, 2)

	// 4. delete topic
	log.Println("=====> start deleting topic after 20s.......")
	time.Sleep(20 * time.Second)
	log.Println("=====> delete topic.......")
	deleteTopic(topic)
	log.Println("=====> deleting topic done .......")

	readTopic("reader3", topic, pulsar.EarliestMessageID(), true, 4)

	time.Sleep(5 * time.Minute)
}
