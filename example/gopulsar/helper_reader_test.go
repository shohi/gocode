package gopulsar

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

// If no message read, return false; otherwise true.
func readTopic(name, topic string, startID pulsar.MessageID, inclusive bool, iter int) bool {
	client := newTestClient(defaultClientOptions())
	defer client.Close()

	// StartMessageID: pulsar.EarliestMessageID(),
	reader, err := client.CreateReader(pulsar.ReaderOptions{
		Name:                    name,
		Topic:                   topic,
		StartMessageID:          startID,
		StartMessageIDInclusive: inclusive,
	})

	if err != nil {
		panic(err)
	}
	defer reader.Close()

	hasMsg := false
	for reader.HasNext() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		msg, err := reader.Next(ctx)
		if err != nil {
			fmt.Printf("===> iter: %v, reader: %v, err: %v\n",
				name, iter, err)

			cancel()
			hasMsg = false
			break
		}

		fmt.Printf("===> iter: %v, reader: %v, message: %v\n",
			iter, name, string(msg.Payload()))

		cancel()
		hasMsg = true
	}

	return hasMsg
}
