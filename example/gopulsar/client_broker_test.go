package gopulsar

import (
	"fmt"
	"testing"
	"time"

	"github.com/Shopify/toxiproxy/toxics"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/stretchr/testify/assert"
)

// Case Design
// - 1. disable client side ping - change default ping interval
// - 2. check broker ping/pong commands

func TestBrokerDisconnect(t *testing.T) {
	assert := assert.New(t)

	// common
	topic := "broker-ping"
	subName := "my-sub"

	client := newTestClient(defaultClientOptions())
	defer client.Close()

	_, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:             topic,
		SubscriptionName:  subName,
		MessageChannel:    make(chan pulsar.ConsumerMessage),
		ReceiverQueueSize: 1,
	})
	assert.Nil(err)

	time.Sleep(10 * time.Minute)
}

func TestBrokerReconnect(t *testing.T) {

}

// Case Design: After consumer gets reconnected, messages will be redelivered
// messages.
// 1. run proxy client
// 2. disable client side ping
// 3. after connected, add latency toxic and wait for broker closing socket
// 4. after client detected, start other subscriber
// 5. check message redelivery
func TestBrokerDispatch_ConsumerReconnect(t *testing.T) {
	assert := assert.New(t)

	topic := "broker-close"
	subName := "my-sub"

	// clear and populate
	prepareMessageForShared(topic, subName, "msg-00", 10)

	proxy, client := newTestClientViaProxy(
		defaultClientOptions(),
		nopToxicOptions(),
	)
	defer client.Close()
	defer proxy.Stop()

	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:              "s1",
		Topic:             topic,
		Type:              pulsar.Shared,
		SubscriptionName:  subName,
		MessageChannel:    make(chan pulsar.ConsumerMessage),
		ReceiverQueueSize: 1,
	})
	assert.Nil(err)
	assert.NotNil(s)
	defer s.Close()

	// consuming
	go func() {
		last := time.Now()
		for m := range s.Chan() {
			printNewMsg(m, &last, "proxy consumer acking after 2m")
			time.Sleep(2 * time.Minute)
			s.Ack(m)
		}
	}()

	txOpts := ToxicOptions{
		Name:   "latency-up",
		Stream: "up",
		Toxic: &toxics.TimeoutToxic{
			Timeout: 40 * 1000, // millisecond
		},
	}
	updateProxyToxic(proxy, txOpts)
	// TODO:
	// 1. remove proxy
	go func() {
		fmt.Printf("=======> proxy toxic removing after 30s ......\n")
		time.Sleep(30 * time.Second)
		removeProxyToxic(proxy, txOpts)
		fmt.Printf("=======> proxy toxic removed......\n")
	}()
	// 2. check redelivery

	// normal client
	nc := newTestClient(defaultClientOptions())
	defer nc.Close()

	ns, err := nc.Subscribe(pulsar.ConsumerOptions{
		Name:              "s2",
		Topic:             topic,
		Type:              pulsar.Shared,
		SubscriptionName:  subName,
		MessageChannel:    make(chan pulsar.ConsumerMessage),
		ReceiverQueueSize: 1,
	})
	assert.Nil(err)

	go func() {
		last := time.Now()
		for m := range ns.Chan() {
			printNewMsg(m, &last, "normal consumer acking after 10s")
			time.Sleep(10 * time.Second)
			s.Ack(m)
		}
	}()

	time.Sleep(10 * time.Minute)
}

// Case Design: during liveness check, broker will not redeliver
// messages.
// 1. run proxy client
// 2. disable client side ping
// 3. after connected, add latency toxic and wait for broker closing socket
// 4. after client detected, start other subscriber
// 5. check message redelivery
func TestBrokerDispatch_PingTimeout(t *testing.T) {
	assert := assert.New(t)

	topic := "proxy-close"
	subName := "my-sub"

	// clear and populate
	prepareMessageForShared(topic, subName, "msg-00", 10)

	proxy, client := newTestClientViaProxy(
		defaultClientOptions(),
		nopToxicOptions(),
	)
	defer client.Close()
	defer proxy.Stop()

	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:              "s1",
		Topic:             topic,
		Type:              pulsar.Shared,
		SubscriptionName:  subName,
		MessageChannel:    make(chan pulsar.ConsumerMessage, 4),
		ReceiverQueueSize: 4,
	})
	assert.Nil(err)
	assert.NotNil(s)
	defer s.Close()

	// consuming
	done := false
	triggerToxicCh := make(chan struct{})
	go func() {
		last := time.Now()
		for m := range s.Chan() {
			if !done {
				close(triggerToxicCh)
				done = true
			}
			fmt.Printf("====> s1, msg chan capacity: %v\n", len(s.Chan()))
			printNewMsg(m, &last, "proxy consumer acking after 2m")
			time.Sleep(2 * time.Minute)
			s.Ack(m)
		}
	}()

	// add toxic
	txOpts := ToxicOptions{
		Name:   "latency-up",
		Stream: "up",
		Toxic: &toxics.TimeoutToxic{
			Timeout: 40 * 1000, // millisecond
		},
	}
	triggerDelCh := make(chan struct{})
	go func() {
		<-triggerToxicCh
		updateProxyToxic(proxy, txOpts)
		close(triggerDelCh)
	}()

	// remove proxy
	go func() {
		<-triggerDelCh
		fmt.Printf("=======> proxy toxic removing after 2m ......\n")
		time.Sleep(2 * time.Minute)
		removeProxyToxic(proxy, txOpts)
		fmt.Printf("=======> proxy toxic removed......\n")
	}()
	// 2. check redelivery

	// normal client
	nc := newTestClient(defaultClientOptions())
	defer nc.Close()

	ns, err := nc.Subscribe(pulsar.ConsumerOptions{
		Name:              "s2",
		Topic:             topic,
		Type:              pulsar.Shared,
		SubscriptionName:  subName,
		MessageChannel:    make(chan pulsar.ConsumerMessage, 2),
		ReceiverQueueSize: 2,
	})
	assert.Nil(err)

	go func() {
		last := time.Now()
		for m := range ns.Chan() {
			printNewMsg(m, &last, "normal consumer acking after 10s")
			time.Sleep(5 * time.Second)
			s.Ack(m)
		}
	}()

	time.Sleep(10 * time.Minute)
}

// Case Design: during liveness check, broker will not redeliver
// messages.
// 1. run proxy client
// 2. disable client side ping
// 3. after connected, add latency toxic and wait for broker closing socket
// 4. after client detected, start other subscriber
// 5. check message redelivery
func TestBroker_ClientSocketRelinked(t *testing.T) {
	assert := assert.New(t)

	topic := "socket-relink"
	subName := "my-sub"

	// clear and populate
	prepareMessageForShared(topic, subName, "msg-00", 10)

	proxy, client := newTestClientViaProxy(
		defaultClientOptions(),
		nopToxicOptions(),
	)
	defer client.Close()
	defer proxy.Stop()

	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:              "s1",
		Topic:             topic,
		Type:              pulsar.Shared,
		SubscriptionName:  subName,
		MessageChannel:    make(chan pulsar.ConsumerMessage, 4),
		ReceiverQueueSize: 4,
	})
	assert.Nil(err)
	assert.NotNil(s)
	defer s.Close()

	// consuming
	done := false
	triggerToxicCh := make(chan struct{})
	go func() {
		last := time.Now()
		for m := range s.Chan() {
			if !done {
				close(triggerToxicCh)
				done = true
			}
			fmt.Printf("====> s1, msg chan capacity: %v\n", len(s.Chan()))
			printNewMsg(m, &last, "proxy consumer acking after 1m")
			time.Sleep(1 * time.Minute)
			s.Ack(m)
		}
	}()

	// add toxic
	txOpts := ToxicOptions{
		Name:   "latency-up",
		Stream: "up",
		Toxic: &toxics.LatencyToxic{
			Latency: 40 * 1000, // millisecond
		},
	}
	triggerDelCh := make(chan struct{})
	go func() {
		<-triggerToxicCh
		updateProxyToxic(proxy, txOpts)
		close(triggerDelCh)
	}()

	// remove proxy
	go func() {
		<-triggerDelCh
		fmt.Printf("=======> proxy toxic removing after 2m - %v......\n",
			time.Now().Add(2*time.Minute),
		)
		time.Sleep(2 * time.Minute)
		removeProxyToxic(proxy, txOpts)
		fmt.Printf("=======> proxy toxic removed - %v......\n",
			time.Now(),
		)
	}()
	// 2. check redelivery

	// normal client
	nc := newTestClient(defaultClientOptions())
	defer nc.Close()

	ns, err := nc.Subscribe(pulsar.ConsumerOptions{
		Name:              "s2",
		Topic:             topic,
		Type:              pulsar.Shared,
		SubscriptionName:  subName,
		MessageChannel:    make(chan pulsar.ConsumerMessage, 2),
		ReceiverQueueSize: 2,
	})
	assert.Nil(err)

	go func() {
		last := time.Now()
		for m := range ns.Chan() {
			printNewMsg(m, &last, "normal consumer acking after 10s")
			time.Sleep(5 * time.Second)
			ns.Ack(m)
		}
	}()

	time.Sleep(10 * time.Minute)
}

// Case Design: client-side close the connection during ping timeout
// 1. run proxy client
// 2. shorten client side ping timeout
// 3. after connected, add latency toxic and wait for client closing socket
// 4. after client detected, start other subscriber
// 5. check message redelivery
func TestBroker_ClientPingTimeout(t *testing.T) {
	assert := assert.New(t)

	topic := "socket-relink"
	subName := "my-sub"

	// clear and populate
	prepareMessageForShared(topic, subName, "msg-00", 10)

	proxy, client := newTestClientViaProxy(
		defaultClientOptions(),
		nopToxicOptions(),
	)
	defer client.Close()
	defer proxy.Stop()

	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:              "s1",
		Topic:             topic,
		Type:              pulsar.Shared,
		SubscriptionName:  subName,
		MessageChannel:    make(chan pulsar.ConsumerMessage, 4),
		ReceiverQueueSize: 4,
	})
	assert.Nil(err)
	assert.NotNil(s)
	defer s.Close()

	// consuming
	done := false
	triggerToxicCh := make(chan struct{})
	go func() {
		last := time.Now()
		for m := range s.Chan() {
			if !done {
				close(triggerToxicCh)
				done = true
			}
			fmt.Printf("====> s1, msg chan capacity: %v\n", len(s.Chan()))
			printNewMsg(m, &last, "proxy consumer acking after 1m")
			time.Sleep(1 * time.Minute)
			s.Ack(m)
		}
	}()

	// add toxic
	txOpts := ToxicOptions{
		Name:   "latency-up",
		Stream: "up",
		Toxic: &toxics.LatencyToxic{
			Latency: 40 * 1000, // millisecond
		},
	}
	triggerDelCh := make(chan struct{})
	go func() {
		<-triggerToxicCh
		updateProxyToxic(proxy, txOpts)
		close(triggerDelCh)
	}()

	// remove proxy
	go func() {
		<-triggerDelCh
		fmt.Printf("=======> proxy toxic removing after 5m - %v......\n",
			time.Now().Add(5*time.Minute),
		)
		time.Sleep(5 * time.Minute)
		removeProxyToxic(proxy, txOpts)
		fmt.Printf("=======> proxy toxic removed......\n")
	}()
	// 2. check redelivery

	// normal client
	nc := newTestClient(defaultClientOptions())
	defer nc.Close()

	ns, err := nc.Subscribe(pulsar.ConsumerOptions{
		Name:              "s2",
		Topic:             topic,
		Type:              pulsar.Shared,
		SubscriptionName:  subName,
		MessageChannel:    make(chan pulsar.ConsumerMessage, 2),
		ReceiverQueueSize: 2,
	})
	assert.Nil(err)

	go func() {
		last := time.Now()
		for m := range ns.Chan() {
			printNewMsg(m, &last, "normal consumer acking after 10s")
			time.Sleep(5 * time.Second)
			ns.Ack(m)
		}
	}()

	time.Sleep(10 * time.Minute)
}

// Case Design: test broker timeout
// 1. run proxy client
// 2. disable client side ping
// 3. after connected, add latency toxic and wait for borker closing socket
func TestBroker_PingTimeout(t *testing.T) {
	assert := assert.New(t)

	topic := "borker-ping-timeout"
	subName := "my-sub"

	proxy, client := newTestClientViaProxy(
		defaultClientOptions(),
		nopToxicOptions(),
	)
	defer client.Close()
	defer proxy.Stop()

	// create a consumer
	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:              "s1",
		Topic:             topic,
		Type:              pulsar.Shared,
		SubscriptionName:  subName,
		MessageChannel:    make(chan pulsar.ConsumerMessage, 1),
		ReceiverQueueSize: 1,
	})
	assert.Nil(err)
	assert.NotNil(s)
	defer s.Close()

	// add toxic
	txOpts := ToxicOptions{
		Name:   "latency-up",
		Stream: "up",
		Toxic: &toxics.LatencyToxic{
			Latency: 40 * 1000, // millisecond
		},
	}
	go func() {
		fmt.Printf("========> add proxy toxic - %v\n", time.Now())
		time.Sleep(1 * time.Second)
		updateProxyToxic(proxy, txOpts)
	}()

	// check ping timeout

	time.Sleep(10 * time.Minute)
}

// Case Design: nack previous messages after reconnected
// messages.
// 1. run proxy client
// 2. disable client side ping
// 3. after connected, add latency toxic and wait for broker closing socket
// 4. after client detected, start other subscriber
// 5. check message redelivery
func TestClient_NackAndReconnect(t *testing.T) {
	assert := assert.New(t)

	topic := "nack-reconnect"
	subName := "my-sub"

	// clear and populate
	prepareMessageForShared(topic, subName, "msg-00", 10)

	proxy, client := newTestClientViaProxy(
		defaultClientOptions(),
		nopToxicOptions(),
	)
	defer client.Close()
	defer proxy.Stop()

	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:              "s1",
		Topic:             topic,
		Type:              pulsar.Shared,
		SubscriptionName:  subName,
		MessageChannel:    make(chan pulsar.ConsumerMessage, 4),
		ReceiverQueueSize: 4,
	})
	assert.Nil(err)
	assert.NotNil(s)
	defer s.Close()

	// consuming
	done := false
	triggerToxicCh := make(chan struct{})
	go func() {
		last := time.Now()
		for m := range s.Chan() {
			if !done {
				close(triggerToxicCh)
				done = true
			}
			fmt.Printf("====> s1, msg chan capacity: %v\n", len(s.Chan()))
			printNewMsg(m, &last, "proxy consumer acking after 2m")
			time.Sleep(1 * time.Minute)
			if m.RedeliveryCount() == 0 {
				s.Nack(m)
			}
		}
	}()

	// add toxic
	txOpts := ToxicOptions{
		Name:   "latency-up",
		Stream: "up",
		Toxic: &toxics.TimeoutToxic{
			Timeout: 40 * 1000, // millisecond
		},
	}
	triggerDelCh := make(chan struct{})
	go func() {
		<-triggerToxicCh
		updateProxyToxic(proxy, txOpts)
		close(triggerDelCh)
	}()

	// remove proxy
	go func() {
		<-triggerDelCh
		fmt.Printf("=======> proxy toxic removing after 2m ......\n")
		time.Sleep(2 * time.Minute)
		removeProxyToxic(proxy, txOpts)
		fmt.Printf("=======> proxy toxic removed......\n")
	}()
	// 2. check redelivery

	// normal client
	nc := newTestClient(defaultClientOptions())
	defer nc.Close()

	ns, err := nc.Subscribe(pulsar.ConsumerOptions{
		Name:              "s2",
		Topic:             topic,
		Type:              pulsar.Shared,
		SubscriptionName:  subName,
		MessageChannel:    make(chan pulsar.ConsumerMessage, 2),
		ReceiverQueueSize: 2,
	})
	assert.Nil(err)

	go func() {
		last := time.Now()
		for m := range ns.Chan() {
			printNewMsg(m, &last, "normal consumer acking after 10s")
			time.Sleep(5 * time.Second)
			s.Ack(m)
		}
	}()

	time.Sleep(10 * time.Minute)
}
