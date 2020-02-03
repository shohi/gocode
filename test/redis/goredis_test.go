package redis_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func newMiniredisForTest() *miniredis.Miniredis {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	return s
}

func TestNewMiniredis(t *testing.T) {
	s := newMiniredisForTest()
	log.Printf("redis server address: %v", s.Addr())
}

func TestGoRedis_Set(t *testing.T) {
	s := newMiniredisForTest()
	defer s.Close()

	client := redis.NewClient(&redis.Options{
		Addr:     s.Addr(),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	// client.Set("goredis", "hello", 5*time.Second)

	err = client.Set("key", "value", 10*time.Second).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	// time.Sleep(10 * time.Minute)
}

func TestGoRedis_Publish(t *testing.T) {
	assert := assert.New(t)

	s := newMiniredisForTest()
	defer s.Close()

	log.Printf("miniredis addr: %v\n", s.Addr())

	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	defer client.Close()

	errCh := make(chan error, 2)

	go func() {
		client2 := redis.NewClient(&redis.Options{
			Addr: s.Addr(),
		})
		sub := client2.Subscribe("foo")
		defer sub.Close()

		assert.NotNil(sub)

		// result, err := sub.ReceiveTimeout(100 * time.Millisecond)
		// skip ACK
		result, err := sub.Receive()

		result, err = sub.Receive()
		log.Println(result, err)

		if err != nil {
			errCh <- err
			return
		}

		msg, ok := result.(*redis.Message)
		assert.True(ok)
		log.Printf("message: %v\n", msg.Payload)
	}()

	err := client.Publish("foo", 100).Err()
	log.Println("published")
	assert.Nil(err)

	select {
	case <-time.After(150 * time.Millisecond):
		// do nothing
	case err := <-errCh:
		assert.Nil(err)
	}
}
