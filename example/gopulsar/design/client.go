package design

import (
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

// Client is used for interaction with Pulsar broker,
// which supports both publishing and subscribing. All
// messages and subscriptions will reside in the same
// tenant/namespace.
type Client interface {
	// Publish messages to specified topic, an optional timeout may be applied.
	// The method returns if client receives the ack or failure occurs.
	Publish(topic string, payload []byte, opts PublishOptions) error

	// Subscribe to specified topic, a Message channel is returned.
	Subscribe(topic string, opts SubscribeOptions) (chan Message, error)

	// Close closes the underlying connection and releases resources.
	// The client should no longer be used after calling Close.
	Close() error
}

// Message abstraction used by consumer.
type Message interface {
	// Payload gets the payload of the message
	Payload() []byte

	// Ack the consumption of the message
	Ack()

	// Acknowledge the failure to process the message.
	//
	// When a message is "negatively acked" it will be marked for redelivery after
	// some fixed delay. The delay is configurable when constructing Client with
	// ClientOptions.MsgNackRedeliveryDelay.
	//
	// This call is not blocking.
	Nack()

	// Get message redelivery count, redelivery count maintain in pulsar broker.
	// When client nack acknowledge messages, broker will dispatch message again
	// with message redelivery count in CommandMessage defined.
	//
	// Message redelivery increases monotonically in a broker, when topic switches
	// ownership to a another broker redelivery count will be recalculated.
	RedeliveryCount() uint32
}

// ClientOptions is used to configure a customized Pulsar client.
type ClientOptions struct {
	// Configure the service URL for the Pulsar service.
	// This parameter is required
	URL string

	// Timeout for the establishment of a TCP connection (default: 30 seconds)
	ConnectionTimeout time.Duration

	// Client name. All the internally created producers and consumers
	// will take the name as prefix.
	Name string

	// Pulsar tenant
	Tenant string

	// Pulsar namespace
	Namespace string

	// PubMaxPendingMessages set the max size of the queue
	// holding the messages pending to receive an acknowledgment
	// from the broker.
	PubMaxPendingMessages int

	// Sets the size of the consumer receive queue.
	SubReceiveQueueSize int

	// The delay after which to redeliver the message which failed to be
	// processed.
	MsgNackRedeliveryDelay time.Duration
}

// PublishOptions is used to configure publishing.
type PublishOptions struct {
	// Timeout sets the timeout for a publish operation. If the ACK
	// is not receieved within given timeout, publish will return error.
	Timeout time.Duration
}

// SubscribeOptions is used to configure subscription.
type SubscribeOptions struct {
	// Select the subscription type to be used when subscribing to the topic.
	// Default is `Exclusive`
	Type pulsar.SubscriptionType

	// Specify the subscription name, which is required when subscribing.
	SubscriptionName string
}
