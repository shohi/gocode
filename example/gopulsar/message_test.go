package gopulsar

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/stretchr/testify/assert"
)

func TestMessagePolicy_Default(t *testing.T) {
	ctx := context.Background()

	client := newTestClient(defaultClientOptions())
	defer client.Close()

	producer := createProducer(client, "topic-msg")
	defer producer.Close()

	msgID, err := producer.Send(ctx, &pulsar.ProducerMessage{
		Payload: []byte("hello"),
	})

	fmt.Println(msgID, err)
}

func TestAck_Repeat(t *testing.T) {
	assert := assert.New(t)

	client := newTestClient(defaultClientOptions())
	defer client.Close()

	topic := "topic-ack-multiple"

	// Publishing message
	go publishMsg(client, PublishOptions{
		topic:    topic,
		interval: 100 * time.Millisecond,
	}, "msg")

	// Consumer
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic: topic,
	})
	assert.Nil(err)

	// ctx := context.Background()
	// tctx, cancel := context.WithTimeout()

	for m := range consumer.Chan() {
		fmt.Printf("====> received msg: %v\n",
			string(m.Payload()))
	}
}

func TestAck_Multiple(t *testing.T) {
	assert := assert.New(t)
	ts := time.Now().Unix()

	client := newTestClient(defaultClientOptions())
	defer client.Close()

	topic := fmt.Sprintf("topic-ack-multiple-%d", ts)
	// topic := "topic-ack-multiple"
	ctx := context.Background()

	// Consumer
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:               topic,
		SubscriptionName:    "my-sub",
		NackRedeliveryDelay: 10 * time.Millisecond,
		Type:                pulsar.Shared,
		// ReceiverQueueSize:   1,
	})
	assert.Nil(err)
	defer consumer.Close()

	// Publishing message
	producer := createProducer(client, topic)
	defer producer.Close()

	_, err = producer.Send(ctx, &pulsar.ProducerMessage{
		Payload: []byte("hello"),
	})
	assert.Nil(err)

	m, err := consumer.Receive(context.Background())
	assert.Nil(err)

	go func() {
		consumer.Nack(m)
		ticker := time.NewTicker(100 * time.Millisecond)
		cnt := 0
		for range ticker.C {
			consumer.Nack(m)
			cnt++
			if cnt > 10 {
				consumer.Ack(m)
			}
		}
	}()

	// consumer
	go func() {
		cnt := 0
		for m := range consumer.Chan() {
			fmt.Printf("====> 1 received msg: %v, topic: %v, cnt: %d\n",
				m.ID(),
				string(m.Payload()),
				cnt)
			cnt++
			consumer.Nack(m)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(10 * time.Minute)
}

func TestMessageDelivery_Ack(t *testing.T) {
	assert := assert.New(t)

	// 0. topic
	topic := "message-ack"
	deleteTopic(topic)

	// 1. create a consumer first to keep messages persistent
	client := newTestClient(defaultClientOptions())
	defer client.Close()

	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "message-ack",
		SubscriptionName: "sub-1",
		Type:             pulsar.Exclusive,
	})
	assert.Nil(err)
	s.Close()

	// 2. publish message
	p, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic:              topic,
		DisableBatching:    true,
		MaxPendingMessages: 1,
	})
	assert.Nil(err)

	_, err = p.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte("hello"),
	})
	assert.Nil(err)
	p.Close()

	// 3. create two different consumers
	s2, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "message-ack",
		SubscriptionName: "sub-2",
		Type:             pulsar.Exclusive,
	})
	assert.Nil(err)
	defer s2.Close()

	s3, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "message-ack",
		SubscriptionName: "sub-3",
		Type:             pulsar.Exclusive,
	})
	assert.Nil(err)
	defer s3.Close()

	// 4. consuming
	go func() {
		select {
		case <-time.After(2 * time.Second):
			assert.Fail("s2 should receive message")
		case m := <-s2.Chan():
			fmt.Printf("s2 received message: %v\n", string(m.Payload()))
			s2.Ack(m)
		}

	}()

	go func() {
		select {
		case <-time.After(2 * time.Second):
			assert.Fail("s3 should receive message")
		case m := <-s3.Chan():
			fmt.Printf("s3 received message: %v\n", string(m.Payload()))
			s3.Ack(m)
		}
	}()

	// previous subscription resubscribe
	s, err = client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "message-ack",
		SubscriptionName: "sub-1",
		Type:             pulsar.Exclusive,
	})
	assert.Nil(err)
	defer s.Close()

	go func() {
		select {
		case <-time.After(2 * time.Second):
			assert.Fail("s1 should receive message")
		case m := <-s.Chan():
			fmt.Printf("s1 received message: %v\n", string(m.Payload()))
			s.Ack(m)
		}
	}()

	// 4. newly subscription should not receive the message

	// wait for broker deleting the message
	time.Sleep(1 * time.Second)

	s4, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "message-ack",
		SubscriptionName: "sub-4",
		Type:             pulsar.Exclusive,
	})
	assert.Nil(err)
	defer s4.Close()

	select {
	case <-time.After(2 * time.Second):
		// do nothing
	case m := <-s4.Chan():
		fmt.Printf("s4 should not receive message - %v\n", string(m.Payload()))
		s4.Ack(m)
	}

}

func TestMessageDelivery_NoSubscribe(t *testing.T) {
	assert := assert.New(t)

	// 0. topic
	topic := "message-ack"
	deleteTopic(topic)
	var s, s2 pulsar.Consumer
	var err error

	// 1. create a consumer first to keep messages persistent
	client := newTestClient(defaultClientOptions())
	defer client.Close()

	s, err = client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "message-ack",
		SubscriptionName: "sub-1",
		Type:             pulsar.Exclusive,
	})
	assert.Nil(err)
	s.Close()

	/*
		s2, err = client.Subscribe(pulsar.ConsumerOptions{
			Topic:            "message-ack",
			SubscriptionName: "sub-2",
			Type:             pulsar.Exclusive,
		})
		assert.Nil(err)
		s2.Close()
	*/

	// 2. publish message
	p, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic:              topic,
		DisableBatching:    true,
		MaxPendingMessages: 1,
	})
	assert.Nil(err)

	_, err = p.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte("hello"),
	})
	assert.Nil(err)
	p.Close()

	// 3. create two different consumers
	s2, err = client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "message-ack",
		SubscriptionName: "sub-2",
		Type:             pulsar.Exclusive,
	})
	assert.Nil(err)
	defer s2.Close()

	// 4. consuming
	go func() {
		select {
		case <-time.After(2 * time.Second):
			assert.Fail("s2 should receive message")
		case m := <-s2.Chan():
			fmt.Printf("s2 received message: %v\n", string(m.Payload()))
			s2.Ack(m)
		}

	}()

	// previous subscription resubscribe
	s, err = client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "message-ack",
		SubscriptionName: "sub-1",
		Type:             pulsar.Exclusive,
	})
	assert.Nil(err)
	defer s.Close()

	go func() {
		select {
		case <-time.After(2 * time.Second):
			assert.Fail("s1 should receive message")
		case m := <-s.Chan():
			fmt.Printf("s1 received message: %v\n", string(m.Payload()))
			s.Ack(m)
		}
	}()

	time.Sleep(30 * time.Second)
}
