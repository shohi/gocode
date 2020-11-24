package redis_test

import (
	"log"
	"testing"

	"github.com/go-redis/redis"
)

func TestClient_GetNil(t *testing.T) {
	s := newMiniredisForTest()
	defer s.Close()

	client := redis.NewClient(&redis.Options{Addr: s.Addr()})

	// If key does not exist, `redis.Nil` will be returned.
	content, err := client.Get("Hello").Bytes()

	log.Printf("content = nil: [%v], err = nil: %v\n", content, err == nil)
}
