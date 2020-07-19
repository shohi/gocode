package gopulsar

import (
	"fmt"
	"testing"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/stretchr/testify/assert"
)

func TestSubscription_DifferentMode(t *testing.T) {
	assert := assert.New(t)

	topic := "resub-diff"
	subName := "my-sub"
	deleteTopic(topic)

	// 1. subscribe with Exclusive mode first
	client := newTestClient(defaultClientOptions())
	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:             "s1",
		Topic:            topic,
		SubscriptionName: subName,
		Type:             pulsar.Exclusive,
	})
	assert.Nil(err)
	s.Close()

	// 2. resubscribe with Shared mode
	s, err = client.Subscribe(pulsar.ConsumerOptions{
		Name:             "s1",
		Topic:            topic,
		SubscriptionName: subName,
		Type:             pulsar.Shared,
	})
	assert.Nil(err)
	if s != nil {
		s.Close()
	}

	// 3. subscribe multiple times
	s2, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:             "s2",
		Topic:            topic,
		SubscriptionName: subName,
		Type:             pulsar.Shared,
	})
	assert.Nil(err)
	if s2 != nil {
		defer s2.Close()
	}

	s3, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:             "s3",
		Topic:            topic,
		SubscriptionName: subName,
		Type:             pulsar.Shared,
	})
	assert.Nil(err)
	if s3 != nil {
		defer s3.Close()
	}
	// publish message
	publishOneMessage(topic, "msg-00", 10, false)

	// consuming
	go func() {
		for m := range s2.Chan() {
			fmt.Printf("=====> %v, message %v\n", m.Consumer.Name(), string(m.Payload()))
			s2.Ack(m)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for m := range s3.Chan() {
			fmt.Printf("=====> %v, message %v\n", m.Consumer.Name(), string(m.Payload()))
			s3.Ack(m)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(1 * time.Minute)
}

func TestSubscription_Newly(t *testing.T) {
	assert := assert.New(t)

	// clear topic
	topic := "subs-newly"
	deleteTopic(topic)

	client := newTestClient(defaultClientOptions())

	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:             "s0",
		Topic:            topic,
		SubscriptionName: "my-sub",
		Type:             pulsar.Exclusive,
	})
	assert.Nil(err)
	s.Close()

	// publish messages
	go publishMsg(client, PublishOptions{
		topic:    topic,
		interval: 2 * time.Second,
	}, "hello")

	// keep subscribing
	ticker := time.NewTicker(10 * time.Second)
	count := 1
	for range ticker.C {
		ns, err := client.Subscribe(pulsar.ConsumerOptions{
			Name:             fmt.Sprintf("s%d", count),
			SubscriptionName: fmt.Sprintf("my-sub-%d", count),
			Topic:            topic,
		})
		assert.Nil(err)

		go func(s pulsar.Consumer) {
			timeout := 10 * time.Second
			for {
				start := time.Now()
				select {
				case <-time.After(timeout):
					//
				case m := <-s.Chan():
					fmt.Printf("%v receives msg - %v\n", m.Consumer.Name(), string(m.Payload()))
					s.Ack(m)
				}

				timeout = timeout - time.Now().Sub(start)
				if timeout <= 0 {
					return
				}
			}
		}(ns)

		count++
		if count == 5 {
			break
		}
	}

	time.Sleep(5 * time.Minute)
}

func TestSubscription_OtherSubscription(t *testing.T) {
	assert := assert.New(t)

	topic := "other-topic"
	subName := "my-sub-0"
	prepareMessageForShared(topic, subName, "hello", 10)

	// consuming
	client := newTestClient(defaultClientOptions())
	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:             "s1",
		Topic:            topic,
		SubscriptionName: "my-sub-1",
		Type:             pulsar.Exclusive,
	})
	assert.Nil(err)
	defer s.Close()

	go func() {
		for m := range s.Chan() {
			fmt.Printf("%v receives msg - %v\n", m.Consumer.Name(), string(m.Payload()))
			s.Ack(m)
		}
	}()

	time.Sleep(3 * time.Minute)
}

func TestSubscription_ConsumeLate(t *testing.T) {
	assert := assert.New(t)

	topic := "consume-late"
	subName := "my-sub"
	prepareMessageForShared(topic, subName, "msg-00", 10)

	client := newTestClient(defaultClientOptions())
	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:             "s1",
		Topic:            topic,
		SubscriptionName: subName + "1",
		Type:             pulsar.Failover,
	})

	assert.Nil(err)
	defer s.Close()

	go func() {
		for m := range s.Chan() {
			fmt.Printf("%v receives msg - %v\n", m.Consumer.Name(), string(m.Payload()))
			s.Ack(m)
		}
	}()

	time.Sleep(5 * time.Minute)
}

func TestSubscription_Delete(t *testing.T) {
	assert := assert.New(t)

	topic := "subs-delete"
	subName := "my-sub"
	prepareMessageForShared(topic, subName, "msg-00", 10)

	client := newTestClient(defaultClientOptions())
	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: subName,
		Type:             pulsar.Shared,
	})
	assert.Nil(err)

	if err != nil {
		panic(err)
	}

	go func() {

	}()
}
