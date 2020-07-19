package gopulsar

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/stretchr/testify/assert"
)

func TestProducer_NonPersistent(t *testing.T) {
	assert := assert.New(t)

	client := newTestClient(defaultClientOptions())
	defer client.Close()

	p, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "non-persistent://public/default/topic-1",
	})
	assert.Nil(err)
	defer p.Close()

	for k := 0; k < 10; k++ {
		_, err := p.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: []byte(fmt.Sprintf("np-%d", k)),
		})
		if err != nil {
			fmt.Printf("===> send message - %v, err: %v\n", k, err)
		}
	}
}

func TestConsumer_NonPersistent(t *testing.T) {
	assert := assert.New(t)

	client := newTestClient(defaultClientOptions())
	defer client.Close()

	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "non-persistent://public/default/topic-1",
		SubscriptionName: "my-np-sub",
	})
	assert.Nil(err)
	defer s.Close()

	go func() {
		ctx := context.Background()
		for {
			msg, err := s.Receive(ctx)
			if err != nil {
				fmt.Printf("===> receive message error, err: %v\n", err)
			} else {
				fmt.Printf("====> receive message: %v\n", string(msg.Payload()))
			}
		}
	}()

	time.Sleep(10 * time.Second)
}

func TestNonPersistent_Nack(t *testing.T) {
	assert := assert.New(t)

	client := newTestClient(defaultClientOptions())
	defer client.Close()

	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic: "non-persistent://public/default/np-topic-1",
		// Topic:               "persistent://public/default/np-topic-1",
		SubscriptionName:    "my-np-sub",
		ReceiverQueueSize:   1,
		NackRedeliveryDelay: time.Millisecond,
	})
	assert.Nil(err)
	defer s.Close()

	go func() {
		ctx := context.Background()
		for {
			msg, _ := s.Receive(ctx)
			if msg != nil {
				fmt.Printf("===> received msg: %v\n", string(msg.Payload()))
				s.Nack(msg)
			}
		}
	}()

	p, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "non-persistent://public/default/np-topic-1",
		// Topic:           "persistent://public/default/np-topic-1",
		DisableBatching: true,
	})
	defer p.Close()

	go func() {
		ctx := context.Background()
		cnt := 0
		for {
			_, _ = p.Send(ctx, &pulsar.ProducerMessage{
				Payload: []byte(fmt.Sprintf("msg-%d", cnt)),
			})
			cnt++
			time.Sleep(100 * time.Millisecond)
		}

	}()

	time.Sleep(5 * time.Second)
}

func TestPersistent_NoSub_Producer(t *testing.T) {
	assert := assert.New(t)

	client := newTestClient(defaultClientOptions())
	defer client.Close()

	topic := "topic-no-sub"
	p, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic:              topic,
		Name:               "no-sub-p",
		DisableBatching:    true,
		MaxPendingMessages: 1,
	})
	assert.Nil(err)
	defer p.Close()

	ticker := time.NewTicker(100 * time.Millisecond)
	ctx := context.Background()

	cnt := 0
	for range ticker.C {
		_, err = p.Send(ctx, &pulsar.ProducerMessage{
			Payload: []byte("test"),
		})
		assert.Nil(err)
		cnt++
		fmt.Printf("===> publish message %v\n", cnt)
		if cnt == 100 {
			break
		}
	}

	ticker.Stop()
}

func TestPersistent_NoSub_Consumer(t *testing.T) {
	assert := assert.New(t)

	client := newTestClient(defaultClientOptions())
	defer client.Close()

	topic := "topic-no-sub"
	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:                        "no-sub-s",
		Topic:                       topic,
		SubscriptionName:            "my-sub",
		Type:                        pulsar.Shared,
		ReceiverQueueSize:           1,
		SubscriptionInitialPosition: pulsar.SubscriptionPositionLatest,
	})
	assert.Nil(err)
	defer s.Close()

	bgCtx := context.Background()
	ctx, cancel := context.WithTimeout(bgCtx, 10*time.Second)
	defer cancel()

	msgCh := s.Chan()
	cnt := 0
	for {
		select {
		case msg := <-msgCh:
			fmt.Printf("======> receive message: %v\n",
				msg.PublishTime())
			cnt++
		case <-ctx.Done():
			goto outer
		}
	}
outer:
	//
}

func TestPersistent_NoSub_Reader(t *testing.T) {
	assert := assert.New(t)

	client := newTestClient(defaultClientOptions())
	defer client.Close()

	topic := "topic-no-sub"
	r, err := client.CreateReader(pulsar.ReaderOptions{
		Name:              "no-sub-r",
		Topic:             topic,
		ReceiverQueueSize: 1,
		StartMessageID:    pulsar.EarliestMessageID(),
	})
	assert.Nil(err)
	defer r.Close()

	ctx := context.Background()
	for r.HasNext() {
		msg, err := r.Next(ctx)
		assert.Nil(err)

		fmt.Printf("======> read message: %v\n",
			msg.PublishTime())
	}
}
