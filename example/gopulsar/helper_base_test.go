package gopulsar

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

var globalID int64
var globalRand *rand.Rand

func init() {
	globalRand = rand.New(rand.NewSource(99))
}

type ClientOptions struct {
	DialTimeout time.Duration
	OpTimeout   time.Duration
	MaxConns    int
}

func defaultClientOptions() ClientOptions {
	return ClientOptions{
		DialTimeout: 10 * time.Second,
		OpTimeout:   10 * time.Second,
		MaxConns:    10, // If set to 1, then there are some interesting happened for multiple clients.
	}
}

func newTestClient(opts ClientOptions) pulsar.Client {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		// URL:                     "pulsar://localhost:6650",
		URL:                     "pulsar://127.0.0.1:6650",
		ConnectionTimeout:       opts.DialTimeout,
		OperationTimeout:        opts.OpTimeout,
		MaxConnectionsPerBroker: opts.MaxConns,
	})

	if err != nil {
		panic(err)
	}

	return client
}

func createProducer(client pulsar.Client, topic string) pulsar.Producer {
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic:           topic,
		DisableBatching: true,
	})
	if err != nil {
		panic(err)
	}

	return producer
}

func createConsumer(client pulsar.Client,
	opts pulsar.ConsumerOptions) (pulsar.Consumer, error) {

	consumer, err := client.Subscribe(opts)

	// id := atomic.AddInt64(&globalID, 1)

	if err == nil {
		go consumeMsg(consumer, opts.Name)
	}

	return consumer, err
}

func consumeMsgAndAck(c pulsar.Consumer, name string, ackRate float64) {
	// TODO

}

func consumeMsg(c pulsar.Consumer, name string) {
	msgCh := c.Chan()
	cnt := 1
	for m := range msgCh {
		if cnt%10 == 0 {
			fmt.Printf("%v ====> nack, %v-%v\n",
				name,
				m.ID(),
				string(m.Payload()),
			)

			c.Nack(m.Message)
		} else {
			fmt.Printf("%v - received message: %v\n",
				name,
				string(m.Payload()))
			c.Ack(m.Message)
		}

		cnt++

		if cnt >= 20 && name == "consumer-1" {
			c.Close()
			return
		}
	}
}

type PublishOptions struct {
	topic    string
	interval time.Duration
}

func publishMsg(c pulsar.Client, opts PublishOptions, msg string) {
	p := createProducer(c, opts.topic)

	ticker := time.NewTicker(opts.interval)
	cnt := 0
	for range ticker.C {
		msg := fmt.Sprintf("%v-%d", msg, cnt)
		p.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: []byte(msg),
		})
		fmt.Printf("====> published msg: %v\n", msg)
		cnt++
	}
}

func publishOnce(p pulsar.Producer, msg string) {
	ctx := context.Background()

	_, err := p.Send(ctx, &pulsar.ProducerMessage{
		Payload: []byte(msg),
	})
	if err != nil {
		panic(err)
	}
}

func publishWithTicker(p pulsar.Producer, interval time.Duration) {
	ctx := context.Background()
	ticker := time.NewTicker(interval)
	counter := 0
	for range ticker.C {
		msg := fmt.Sprintf("msg-%d", counter)
		start := time.Now()
		_, err := p.Send(ctx, &pulsar.ProducerMessage{
			Payload: []byte(msg),
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("====> publishing messasge: %v, elapsed: %v\n",
			counter, time.Now().Sub(start))
		counter++
	}
}

func publishOneMessage(topic, msg string, count int, skip bool) {
	if skip {
		return
	}

	defer func() {
		fmt.Printf("\n\n\n=======> successful published %v message\n\n\n\n", count)
	}()

	client := newTestClient(defaultClientOptions())
	defer client.Close()

	p, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic:              topic,
		MaxPendingMessages: 1,
		DisableBatching:    true,
	})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	// Publish message
	for k := 0; k < count; k++ {
		publishOnce(p, fmt.Sprintf("%v - %d", msg, k))
	}
}

func createSubscription(topic, subName string, typ pulsar.SubscriptionType) {
	client := newTestClient(defaultClientOptions())
	defer client.Close()

	_, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:             topic,
		SubscriptionName:  subName,
		Type:              typ,
		ReceiverQueueSize: 1,
	})

	if err != nil {
		panic(err)
	}
}

func consumeSubscription(topic, subName string, typ pulsar.SubscriptionType) {

}

func prepareMessageForShared(topic, subName, msg string, count int) {
	// clean up topic firstly
	deleteTopic(topic)

	// first subscribe to keep message
	// and then publish given count of messsages
	createSubscription(topic, subName, pulsar.Shared)
	publishOneMessage(topic, msg, count, false)
}

func printNewMsg(msg pulsar.ConsumerMessage, last *time.Time, info string) {
	gap := time.Now().Sub(*last)
	*last = time.Now()

	msgFormat := "====> %v, messsge: %v, " +
		"redeliver: %v, " +
		"now: %v, " +
		"elapse: %v, " +
		"last: %v, " +
		info + "\n"

	fmt.Printf(
		msgFormat,
		msg.Consumer.Name(),
		string(msg.Payload()),
		msg.RedeliveryCount(),
		time.Now(),
		gap,
		*last,
	)
}

func newSharedSubscription(name string, topic, subName string) pulsar.Consumer {
	client := newTestClient(defaultClientOptions())
	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:                name,
		Topic:               topic,
		SubscriptionName:    subName,
		Type:                pulsar.Shared,
		ReceiverQueueSize:   1,
		MessageChannel:      make(chan pulsar.ConsumerMessage),
		NackRedeliveryDelay: 10 * time.Millisecond,
	})
	if err != nil {
		panic(err)
	}

	return s
}
