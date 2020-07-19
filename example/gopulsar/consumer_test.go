package gopulsar

import (
	"log"
	"testing"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/apache/pulsar-client-go/pulsar"
)

func init() {
	// logrus.SetLevel(logrus.DebugLevel)
	logrus.SetLevel(logrus.InfoLevel)
}

// copied from `pulsar-client-go` example folder

func TestSubscription_SameName(t *testing.T) {
	client := newTestClient(defaultClientOptions())
	defer client.Close()

	consumer1, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:             "consumer",
		Topic:            "topic-1",
		SubscriptionName: "my-sub",
		Type:             pulsar.Failover,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer1.Close()

	consumer2, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:             "consumer",
		Topic:            "topic-1",
		SubscriptionName: "my-sub",
		Type:             pulsar.Failover,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer2.Close()
}

func TestSubscription_Exclusive(t *testing.T) {
	client := newTestClient(defaultClientOptions())
	defer client.Close()

	consumer1, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "topic-1",
		SubscriptionName: "my-sub",
		Type:             pulsar.Exclusive,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer1.Close()

	// In `Exclusive` mode, an other consumer can't be created.
	consumer2, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "topic-1",
		SubscriptionName: "my-sub",
		Type:             pulsar.Exclusive,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer2.Close()
}

func TestSubscription_Failover(t *testing.T) {
	client := newTestClient(defaultClientOptions())
	defer client.Close()

	consumer1, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:             "consumer-1",
		Topic:            "topic-failover",
		SubscriptionName: "my-sub",
		// Type:                pulsar.Shared,
		Type:                pulsar.Failover,
		NackRedeliveryDelay: 10 * time.Millisecond,
		// ReceiverQueueSize:   50,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer1.Close()
	go consumeMsg(consumer1, "consumer-1")

	// In `Exclusive` mode, an other consumer can't be created.
	consumer2, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:             "consumer-2",
		Topic:            "topic-failover",
		SubscriptionName: "my-sub",
		Type:             pulsar.Failover,
		// Type:                pulsar.Shared,
		NackRedeliveryDelay: 10 * time.Millisecond,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer2.Close()
	go consumeMsg(consumer2, "consumer-2")

	go publishMsg(client, PublishOptions{
		topic:    "topic-failover",
		interval: 50 * time.Millisecond,
	}, "hello")

	time.Sleep(3 * time.Second)
}

func TestConsumer_NegativeQueueSize(t *testing.T) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650",
	})

	if err != nil {
		t.FailNow()
	}
	defer client.Close()

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:             "topic",
		SubscriptionName:  "my-sub",
		ReceiverQueueSize: -1,
	})
	defer func() {
		if consumer != nil {
			consumer.Close()
		}
	}()

	if err != nil {
		t.FailNow()
	}
}

func TestSubscription_Shared_Persistent(t *testing.T) {
	client := newTestClient(defaultClientOptions())
	defer client.Close()

	consumer1, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:             "consumer-1",
		Topic:            "shared-topic-1",
		SubscriptionName: "my-sub",
		Type:             pulsar.Shared,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer1.Close()
	go consumeMsg(consumer1, "consumer-1")

	consumer2, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:             "consumer-2",
		Topic:            "shared-topic-1",
		SubscriptionName: "my-sub",
		Type:             pulsar.Shared,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer2.Close()
	go consumeMsg(consumer2, "consumer-2")

	go publishMsg(client, PublishOptions{
		topic:    "shared-topic-1",
		interval: 50 * time.Millisecond,
	}, "hello")

	time.Sleep(3 * time.Second)
}

func TestSubscription_Shared_NoPersistent(t *testing.T) {
	client := newTestClient(defaultClientOptions())
	defer client.Close()

	consumer1, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:             "consumer-1",
		Topic:            "non-persistent://public/default/shared-np-topic-1",
		SubscriptionName: "my-sub",
		Type:             pulsar.Shared,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer1.Close()
	go consumeMsg(consumer1, "consumer-1")

	consumer2, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:             "consumer-2",
		Topic:            "non-persistent://public/default/shared-np-topic-1",
		SubscriptionName: "my-sub",
		Type:             pulsar.Shared,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer2.Close()
	go consumeMsg(consumer2, "consumer-2")

	go publishMsg(client, PublishOptions{
		topic:    "non-persistent://public/default/shared-np-topic-1",
		interval: 50 * time.Millisecond,
	}, "hello")

	time.Sleep(3 * time.Second)
}

func TestConsumer_Compacted(t *testing.T) {
	// TODO

}

// TODO
func TestConsumer_CloseAndMsgChan(t *testing.T) {

}
