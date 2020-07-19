package gopulsar

import (
	"fmt"
	"testing"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func init() {
	// show timestamp
	formatter := &logrus.TextFormatter{
		TimestampFormat: time.RFC3339Nano,
		FullTimestamp:   true,
	}

	logrus.SetFormatter(formatter)
}

func TestCreateClient_InvalidURL(t *testing.T) {
	assert := assert.New(t)
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://localhost:7760",
		ConnectionTimeout: 1 * time.Second,
		OperationTimeout:  1 * time.Second,
	})

	assert.Nil(err)
	assert.NotNil(client)

	fmt.Printf("=====> create client done\n")

	p, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "unreached-01",
	})
	assert.Nil(err)
	assert.NotNil(p)
}
