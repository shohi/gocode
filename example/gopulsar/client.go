package gopulsar

type Config struct {
}

type Client interface {
	Publish(topic string, data []byte) error
	Subscribe(topic string) (Consumer, error)
	Close() error
}

type Message interface {
	Ack()
	Nack()
	Payload()
	RedeliveryCount() uint32
}

type Consumer interface {
	// Unsubscribe the consumer
	Unsubscribe() error

	// Chan returns a channel to consume messages from
	Chan() <-chan Message

	// Close the consumer and stop the broker to push more messages
	Close()
}
