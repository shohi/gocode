package gopulsar

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/Shopify/toxiproxy/toxics"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/stretchr/testify/assert"
)

func TestConsumer_Proxy(t *testing.T) {
	assert := assert.New(t)
	opts := ToxicOptions{
		Stream: "upstream",
		Toxic: &toxics.LatencyToxic{
			Latency: 1 * 1000, // 1s
		},
	}

	proxy, client := newTestClientViaProxy(
		defaultClientOptions(), opts)

	defer func() {
		fmt.Println("===> closing proxy")
		proxy.Stop()
		fmt.Println("===> closed proxy")

		fmt.Println("===> closing client")
		_, err := client.CreateProducer(pulsar.ProducerOptions{
			Name:  "test",
			Topic: "test-aa",
		})

		fmt.Printf("=======> create producer, err: %v", err)

		client.Close()
		fmt.Println("===> closed client")
	}()

	fmt.Printf("======> new proxy client\n")
	proxyTopic := "proxy-topic-1"
	p, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic:              proxyTopic,
		Name:               "proxy-p-1",
		DisableBatching:    true,
		MaxPendingMessages: 1,
	})
	fmt.Printf("======> create producer \n")
	assert.Nil(err)
	assert.NotNil(p)

	fmt.Printf("======> end \n")
}

// Case Design
// 1. one producer publishes messages
// 2. two consumers subscribes same topic
// - one directly connects to the broker
// - one connects to proxy

// Case 1 - Ack Failure log
func TestAckFailure(t *testing.T) {
	assert := assert.New(t)

	topic := "ack-fail-test"

	// clean up topic firstly
	deleteTopic(topic)

	clientOpts := ClientOptions{
		DialTimeout: 2 * time.Second,
		OpTimeout:   500 * time.Millisecond,
	}

	txOpts := ToxicOptions{
		Name:   "latency-up",
		Stream: "up",
		Toxic: &toxics.LatencyToxic{
			Latency: 50, // millisecond
		},
	}
	proxy, client := newTestClientViaProxy(clientOpts, txOpts)
	defer proxy.Stop()
	defer client.Close()

	fmt.Printf("========> start scribe\n\n")

	start := time.Now()
	s1, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:                "c",
		Topic:               topic,
		SubscriptionName:    "my-sub",
		Type:                pulsar.Failover,
		ReceiverQueueSize:   1,
		NackRedeliveryDelay: 0,
	})
	fmt.Printf("======> subscribe elapse: %v\n", time.Now().Sub(start))
	assert.Nil(err)
	assert.NotNil(s1)
	defer s1.Close()

	nTxOpts := ToxicOptions{
		Name:   "latency-up",
		Stream: "up",
		Toxic: &toxics.TimeoutToxic{
			Timeout: 1000, // millisecond
		},
	}
	go func() {
		<-time.After(3 * time.Second)
		updateProxyToxic(proxy, nTxOpts)
		fmt.Printf("======> update toxic: %v\n", nTxOpts.Name)
	}()

	/*
		go func() {
			<-time.After(6 * time.Second)
			removeProxyToxic(proxy, nTxOpts)
			fmt.Printf("======> update toxic: %v\n", nTxOpts.Name)
		}()
	*/

	go func() {
		msgCh := s1.Chan()
		counter := 0
		for msg := range msgCh {
			counter++
			s1.Ack(msg)
			fmt.Printf("====> s1, messasge: %v, counter: %v, acking\n",
				string(msg.Payload()), counter)
		}
	}()

	client1 := newTestClient(clientOpts)
	p1, err := client1.CreateProducer(pulsar.ProducerOptions{
		Topic:              topic,
		MaxPendingMessages: 1,
		DisableBatching:    true,
	})
	assert.Nil(err)
	assert.NotNil(p1)
	defer p1.Close()
	defer client1.Close()

	// Publish message
	go publishWithTicker(p1, 2*time.Second)
	// publishOnce(p1, "msg-00")

	time.Sleep(5 * time.Minute)
}

// Case X - Ack Failure X
func TestAckXXX(t *testing.T) {
	assert := assert.New(t)

	topic := "ack-fail-test"

	// clean up topic firstly
	deleteTopic(topic)

	clientOpts := ClientOptions{
		DialTimeout: 2 * time.Second,
		OpTimeout:   500 * time.Millisecond,
	}

	txOpts := ToxicOptions{
		Name:   "latency-up",
		Stream: "up",
		Toxic: &toxics.LatencyToxic{
			Latency: 50, // millisecond
		},
	}
	proxy, client := newTestClientViaProxy(clientOpts, txOpts)
	defer proxy.Stop()
	defer client.Close()

	fmt.Printf("========> start scribe\n\n")

	start := time.Now()
	s1, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:                "c",
		Topic:               topic,
		SubscriptionName:    "my-sub",
		Type:                pulsar.Failover,
		ReceiverQueueSize:   1,
		NackRedeliveryDelay: 0,
	})
	fmt.Printf("======> subscribe elapse: %v\n", time.Now().Sub(start))
	assert.Nil(err)
	assert.NotNil(s1)
	defer s1.Close()

	nTxOpts := ToxicOptions{
		Name:   "latency-up",
		Stream: "up",
		Toxic: &toxics.TimeoutToxic{
			Timeout: 1000, // millisecond
		},
	}
	go func() {
		<-time.After(3 * time.Second)
		updateProxyToxic(proxy, nTxOpts)
		fmt.Printf("======> update toxic: %v\n", nTxOpts.Name)
	}()

	/*
		go func() {
			<-time.After(6 * time.Second)
			removeProxyToxic(proxy, nTxOpts)
			fmt.Printf("======> update toxic: %v\n", nTxOpts.Name)
		}()
	*/

	go func() {
		msgCh := s1.Chan()
		counter := 0
		for msg := range msgCh {
			counter++
			s1.Ack(msg)
			fmt.Printf("====> s1, messasge: %v, counter: %v, acking\n",
				string(msg.Payload()), counter)
		}
	}()

	client1 := newTestClient(clientOpts)
	p1, err := client1.CreateProducer(pulsar.ProducerOptions{
		Topic:              topic,
		MaxPendingMessages: 1,
		DisableBatching:    true,
	})
	assert.Nil(err)
	assert.NotNil(p1)
	defer p1.Close()
	defer client1.Close()

	s99, err := client1.Subscribe(pulsar.ConsumerOptions{
		Name:                "b",
		Topic:               topic,
		SubscriptionName:    "my-sub",
		Type:                pulsar.Failover,
		ReceiverQueueSize:   1,
		NackRedeliveryDelay: 0,
	})
	assert.Nil(err)
	assert.NotNil(s99)

	// consume
	go func() {
		msgCh := s99.Chan()
		counter := 0
		for msg := range msgCh {
			fmt.Printf("====> s99, messasge: %v, counter: %v\n",
				string(msg.Payload()), counter)
			counter++
		}
	}()

	// Publish message
	go publishWithTicker(p1, 2*time.Second)
	// publishOnce(p1, "msg-00")

	time.Sleep(5 * time.Minute)
}

// Case Design: Ack Failure - 确保在重连之前, 发送Ack消息
// - Publish
// 1. 用正常的client保证connection正常 -- 使用latency
// 2. 确保一条消息被存储 -- PublishOne Message
//
// - Consumer
// 1. 使用proxy, 正确设置DialTimeout/OpTimeout
//    保证connection正常 -- 使用upstream latency.
// 2. 接受消息, 发Ack
// 3. 在Latency期间, 将proxy停掉
// 4. 确保Ack错误后, 重启proxy让连接恢复

func TestAckFailure_OneConsumer(t *testing.T) {
	assert := assert.New(t)

	topic := "ack-fail-one-consumer"

	// clean up topic firstly
	deleteTopic(topic)

	// 1. normal client - producer
	createSubscription(topic, "my-sub", pulsar.Shared)
	publishOneMessage(topic, "msg-00", 10, false)

	// 2. proxy client - consumer
	clientOpts := ClientOptions{
		DialTimeout: 10 * time.Second,
		OpTimeout:   5 * time.Second,
	}
	txOpts := ToxicOptions{
		Name:   "latency-up",
		Stream: "up",
		Toxic: &toxics.LatencyToxic{
			Latency: 2 * 1000, // millisecond
		},
	}
	proxy, client := newTestClientViaProxy(clientOpts, txOpts)
	defer proxy.Stop()
	defer client.Close()

	fmt.Printf("========> start subscribe\n\n")

	start := time.Now()
	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:                "s",
		Topic:               topic,
		SubscriptionName:    "my-sub",
		Type:                pulsar.Shared,
		ReceiverQueueSize:   4,
		NackRedeliveryDelay: 10 * time.Millisecond,
	})

	fmt.Printf("======> subscribe elapse: %v\n", time.Now().Sub(start))
	assert.Nil(err)
	defer s.Close()

	// Main Operation
	closeCh := make(chan struct{}, 1)
	restartCh := make(chan struct{}, 1)

	// stop proxy
	go func() {
		<-closeCh
		time.Sleep(100 * time.Millisecond)
		log.Println("=====> stopping proxy...")
		proxy.Stop()

		log.Println("=====> stopped proxy, will restart 15s later...")
		time.AfterFunc(15*time.Second, func() {
			restartCh <- struct{}{}
		})
	}()

	// restart proxy
	go func() {
		<-restartCh

		log.Println("=====> restarting proxy...")
		proxy.Toxics.ResetToxics()
		err := proxy.Start()
		if err != nil {
			panic(err)
		}
		log.Println("=====> restarted proxy...")
	}()

	msgCh := s.Chan()
	counter := 0
	done := false
	for msg := range msgCh {
		counter++

		// only close once
		if !done {
			closeCh <- struct{}{}
			done = true
		}
		/*
			if counter == 3 {
				if !done {
					proxy.Stop()
					time.AfterFunc(15*time.Second, func() {
						restartCh <- struct{}{}
					})
					done = true
				}
			}
		*/

		s.Ack(msg)
		fmt.Printf("====> s, message: %v, counter: %v, redeliver: %v, acking\n",
			string(msg.Payload()), counter, msg.RedeliveryCount())
	}

	time.Sleep(5 * time.Minute)
}

func TestNackFailure_OneConsumer(t *testing.T) {
	assert := assert.New(t)

	topic := "ack-fail-one-consumer"

	// clean up topic firstly
	deleteTopic(topic)

	// 1. normal client - producer
	createSubscription(topic, "my-sub", pulsar.Shared)
	publishOneMessage(topic, "msg-00", 1, false)

	// 2. proxy client - consumer
	clientOpts := ClientOptions{
		DialTimeout: 10 * time.Second,
		OpTimeout:   5 * time.Second,
	}
	txOpts := ToxicOptions{
		Name:   "latency-up",
		Stream: "up",
		Toxic: &toxics.LatencyToxic{
			Latency: 2 * 1000, // millisecond
		},
	}
	proxy, client := newTestClientViaProxy(clientOpts, txOpts)
	defer proxy.Stop()
	defer client.Close()

	fmt.Printf("========> start subscribe\n\n")

	start := time.Now()
	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:                "s",
		Topic:               topic,
		SubscriptionName:    "my-sub",
		Type:                pulsar.Shared,
		ReceiverQueueSize:   1,
		NackRedeliveryDelay: 10 * time.Millisecond,
	})

	fmt.Printf("======> subscribe elapse: %v\n", time.Now().Sub(start))
	assert.Nil(err)
	defer s.Close()

	// Main Operation
	closeCh := make(chan struct{}, 1)
	restartCh := make(chan struct{}, 1)

	// stop proxy
	go func() {
		<-closeCh
		time.Sleep(100 * time.Millisecond)
		log.Println("=====> stopping proxy...")
		proxy.Stop()

		log.Println("=====> stopped proxy, will restart 15s later...")
		time.AfterFunc(15*time.Second, func() {
			restartCh <- struct{}{}
		})
	}()

	// restart proxy
	go func() {
		<-restartCh

		log.Println("=====> restarting proxy...")
		proxy.Toxics.ResetToxics()
		err := proxy.Start()
		if err != nil {
			panic(err)
		}
		log.Println("=====> restarted proxy...")
	}()

	msgCh := s.Chan()
	counter := 0
	done := false
	lastTime := time.Now()

	for msg := range msgCh {
		counter++

		// only close once
		if !done {
			closeCh <- struct{}{}
			done = true
		}
		/*
			if counter == 3 {
				if !done {
					proxy.Stop()
					time.AfterFunc(15*time.Second, func() {
						restartCh <- struct{}{}
					})
					done = true
				}
			}
		*/

		start := time.Now()
		gap := time.Now().Sub(lastTime)
		lastTime = time.Now()

		s.Nack(msg)
		fmt.Printf("====> s, messsge: %v, counter: %v, redeliver: %v, elapse: %v, last: %v, nacking\n",
			string(msg.Payload()),
			counter,
			msg.RedeliveryCount(),
			time.Now().Sub(start),
			gap,
		)
	}

	time.Sleep(5 * time.Minute)
}

func TestNackFailure_MultipleClient(t *testing.T) {
	assert := assert.New(t)

	topic := "ack-fail-one-consumer"
	subName := "my-sub"

	// setup
	prepareMessageForShared(topic, subName, "msg-00", 10)

	// 2. proxy client - consumer
	clientOpts := ClientOptions{
		DialTimeout: 10 * time.Second,
		OpTimeout:   5 * time.Second,
	}
	txOpts := ToxicOptions{
		Name:   "latency-up",
		Stream: "up",
		Toxic: &toxics.LatencyToxic{
			Latency: 2 * 1000, // millisecond
		},
	}
	proxy, client := newTestClientViaProxy(clientOpts, txOpts)
	defer proxy.Stop()
	defer client.Close()

	fmt.Printf("========> start subscribe\n\n")

	start := time.Now()
	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:                "proxy_consumer",
		Topic:               topic,
		SubscriptionName:    subName,
		Type:                pulsar.Shared,
		ReceiverQueueSize:   4,
		NackRedeliveryDelay: 10 * time.Millisecond,
	})

	fmt.Printf("======> subscribe elapse: %v\n", time.Now().Sub(start))
	assert.Nil(err)
	defer s.Close()

	// Main Operation
	closeCh := make(chan struct{}, 1)
	restartCh := make(chan struct{}, 1)

	// stop proxy
	go func() {
		<-closeCh
		time.Sleep(100 * time.Millisecond)
		log.Println("=====> stopping proxy...")
		proxy.Stop()

		log.Println("=====> stopped proxy, will restart 15s later...")
		time.AfterFunc(15*time.Second, func() {
			restartCh <- struct{}{}
		})
	}()

	// restart proxy
	go func() {
		<-restartCh

		log.Println("=====> restarting proxy...")
		proxy.Toxics.ResetToxics()
		err := proxy.Start()
		if err != nil {
			panic(err)
		}

		// create a new consumer
		var newConsumer = true
		// var newConsumer = false
		go func() {
			if !newConsumer {
				return
			}
			nc := newTestClient(defaultClientOptions())
			ns, err := nc.Subscribe(pulsar.ConsumerOptions{
				Name:                "direct_consumer",
				Topic:               topic,
				SubscriptionName:    subName,
				Type:                pulsar.Shared,
				ReceiverQueueSize:   1,
				MessageChannel:      make(chan pulsar.ConsumerMessage, 1),
				NackRedeliveryDelay: 10 * time.Millisecond,
			})
			if err != nil {
				panic(ns)
			}

			last := time.Now()
			for msg := range ns.Chan() {
				gap := time.Now().Sub(last)
				last = time.Now()

				fmt.Printf("====> ns, messsge: %v, redeliver: %v, elapse: %v\n",
					string(msg.Payload()), msg.RedeliveryCount(), gap,
				)
				ns.Nack(msg)
				time.Sleep(1 * time.Second)
			}
		}()

		log.Println("=====> restarted proxy...")
	}()

	msgCh := s.Chan()
	counter := 0
	lastTime := time.Now()

	var done = false
	var triggerClose = make(chan struct{})
	go func() {
		<-triggerClose
		fmt.Printf("====> trigger closing after 1s ...\n")
		time.AfterFunc(1*time.Second, func() {
			closeCh <- struct{}{}
		})
	}()

	for msg := range msgCh {
		counter++

		// only close once
		if !done {
			close(triggerClose)
			done = true
		}

		start := time.Now()
		gap := time.Now().Sub(lastTime)
		lastTime = time.Now()

		s.Nack(msg)
		fmt.Printf("====> s, messsge: %v, counter: %v, redeliver: %v, elapse: %v, last: %v, nacking\n",
			string(msg.Payload()),
			counter,
			msg.RedeliveryCount(),
			time.Now().Sub(start),
			gap,
		)
		time.Sleep(1 * time.Second)
	}

	time.Sleep(5 * time.Minute)
}

func TestNack_Order(t *testing.T) {
	assert := assert.New(t)

	topic := "nack-one-consumer"
	subName := "my-sub"

	// setup
	prepareMessageForShared(topic, subName, "msg-00", 10)

	// normal susbscribe
	client := newTestClient(defaultClientOptions())
	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:                "direct_consumer",
		Topic:               topic,
		SubscriptionName:    subName,
		Type:                pulsar.Shared,
		ReceiverQueueSize:   1,
		MessageChannel:      make(chan pulsar.ConsumerMessage),
		NackRedeliveryDelay: 10 * time.Millisecond,
	})
	assert.Nil(err)

	// consuming
	go func() {
		last := time.Now()
		count := 1
		for msg := range s.Chan() {
			start := time.Now()
			if count%3 == 0 {
				msg.Consumer.Nack(msg)
				printNewMsg(msg, &last, fmt.Sprintf("[%v]nacking: %v",
					count, time.Now().Sub(start)))
			} else {
				printNewMsg(msg, &last, fmt.Sprintf("[%v]new message", count))
			}
			count++

			time.Sleep(10 * time.Second)
		}
	}()

	time.Sleep(5 * time.Minute)
}

func TestNoAckNack_OneConsumer(t *testing.T) {
	assert := assert.New(t)

	topic := "no-ack-nack"
	subName := "my-sub"

	// setup
	prepareMessageForShared(topic, subName, "msg-00", 10)

	// normal susbscribe
	client := newTestClient(defaultClientOptions())
	s, err := client.Subscribe(pulsar.ConsumerOptions{
		Name:                "s1",
		Topic:               topic,
		SubscriptionName:    subName,
		Type:                pulsar.Shared,
		ReceiverQueueSize:   1,
		MessageChannel:      make(chan pulsar.ConsumerMessage),
		NackRedeliveryDelay: 10 * time.Millisecond,
	})
	assert.Nil(err)

	// consuming
	go func() {
		last := time.Now()
		count := 1
		for msg := range s.Chan() {
			printNewMsg(msg, &last, fmt.Sprintf("[%v]new message", count))
			time.Sleep(2 * time.Second)
			count++
			if count == 10 {
				go func() {
					s2 := newSharedSubscription("s1", topic, subName)
					fmt.Printf("======> new subscription start....\n")
					for m := range s2.Chan() {
						printNewMsg(m, &last, fmt.Sprintf("[s2]new message"))
					}
				}()
			}
		}
	}()

	time.Sleep(5 * time.Minute)
}

func TestConsume_ExistingTopic(t *testing.T) {
	topic := "no-ack-nack"
	subName := "my-sub"

	s := newSharedSubscription("draining", topic, subName)

	go func() {
		last := time.Now()
		for m := range s.Chan() {
			printNewMsg(m, &last, fmt.Sprintf("[draining]new message"))
			m.Ack(m)
		}
	}()

	time.Sleep(5 * time.Minute)
}
