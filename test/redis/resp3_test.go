package redis_test

import (
	"log"
	"testing"

	"github.com/shohi/gocode/pkg/serde"
	"github.com/stretchr/testify/assert"
)

func TestRESP3_Serialize(t *testing.T) {
	data := serde.SerializeRawRESP3("PUBLISH", "foo", "bar")
	/*
		var buf bytes.Buffer
		w := resp3.NewWriter(&buf)
		w.WriteCommand("PUBLISH", "foo", "bar")
		w.Flush()
	*/

	log.Printf("content: [%v]", string(data))
	log.Printf("content-q: [%q]", string(data))
}

func TestRESP3_Deserialize(t *testing.T) {
	assert := assert.New(t)
	data := serde.SerializeRawRESP3("PUBLISH", "foo", "bar")
	cmd, err := serde.DeserializeRESP3(data)

	assert.Nil(err)

	log.Printf("cmd: %v", cmd.SmartResult())
}
