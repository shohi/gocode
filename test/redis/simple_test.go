package redis_test

import (
	"fmt"
	"testing"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func TestMember_set(t *testing.T) {
	c := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	ok, err := c.HSet("test", "k1", "v1").Result()

	fmt.Printf("=======> hset, ok: %v, err: %v\n", ok, err)
}

func TestRedisOptions_Reuse(t *testing.T) {
	assert := assert.New(t)

	s1 := newMiniredisForTest()
	defer s1.Close()
	err := s1.Set("server", "s1")
	assert.Nil(err)

	s2 := newMiniredisForTest()
	defer s2.Close()
	err = s2.Set("server", "s2")
	assert.Nil(err)

	conf := redis.Options{
		Addr: s1.Addr(),
	}

	client := redis.NewClient(&conf)
	val, err := client.Get("server").Result()
	assert.Nil(err)
	assert.Equal("s1", val)

	// another client
	conf2 := conf
	conf2.Addr = s2.Addr()

	client = redis.NewClient(&conf)
	val, err = client.Get("server").Result()
	assert.Nil(err)
	assert.NotEqual("s2", val)
	assert.Equal("s1", val)
}
