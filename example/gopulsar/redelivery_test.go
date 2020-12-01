package gopulsar

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func TestMessage_Redelivery(t *testing.T) {
	client := newTestClient(defaultClientOptions())
	defer client.Close()

	topic := "test-redelivery"
	deleteTopic(topic)

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:                "consumer-1",
		Topic:               topic,
		SubscriptionName:    "my-sub",
		Type:                pulsar.Shared,
		NackRedeliveryDelay: 50 * time.Millisecond,
	})
	if err != nil {
		log.Fatal(err)
	}

	// publish msg
	publishOneMessage(topic, "hello", 1, false)

	// consumer
	for {
		m, _ := consumer.Receive(context.Background())
		fmt.Printf("====> msg: %v, redeliver: %v\n",
			string(m.Payload()), m.RedeliveryCount())
		consumer.Nack(m)
		time.Sleep(2 * time.Second)

	}
	// for m := range consumer.Chan() {
	//}
}
