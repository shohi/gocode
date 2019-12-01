package godb

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

func TestGoRedis(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
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
