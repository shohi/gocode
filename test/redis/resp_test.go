package redis_test

import (
	"log"
	"testing"

	"github.com/bsm/redeo/resp"
	"github.com/shohi/gocode/pkg/serde"
)

func TestRESP_Serialize(t *testing.T) {
	c := resp.NewCommand("PUBLISH", []byte("foo"), []byte("bar"))
	data := serde.SerializeRESP(c)
	log.Printf("content: [%v]", string(data))
	log.Printf("content-q: [%q]", string(data))
}

func TestRESP_Deserialize(t *testing.T) {
	c := resp.NewCommand("PUBLISH", []byte("foo"), []byte("bar"))
	data := serde.SerializeRESP(c)

	cmd, err := serde.DeseializeRESP(data)
	log.Printf("cmd: %v,err: %v ", cmd, err)
}
