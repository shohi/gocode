package basic

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"

	"golang.org/x/net/context"
)

type key string

func TestContext(t *testing.T) {
	keyStr := key("key")
	valueCtx := context.WithValue(context.Background(), keyStr, "value")

	dlCtx, dlCancelFunc := context.WithTimeout(valueCtx, 10*time.Second)
	defer dlCancelFunc()

	ctx, doCancelFunc := context.WithCancel(dlCtx)
	defer doCancelFunc()

	log.Println(ctx)
}

func TestContextParallel(t *testing.T) {
	ctx, doCancelFunc := context.WithCancel(context.Background())
	defer doCancelFunc()

	testFileURL := "https://speed.hetzner.de/1GB.bin"
	client := &http.Client{}

	go func(c context.Context) {

		select {
		case <-c.Done():
			log.Println("exit by context Done")
		default:
			log.Println(testFileURL)
			log.Println(client)

			resp, err := client.Get(testFileURL)
			if err != nil {
				return
			}
			if resp != nil {
				defer resp.Body.Close()
				io.Copy(ioutil.Discard, resp.Body)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
}

func TestDoCancelMultiTimes(t *testing.T) {
	_, doCancelFunc := context.WithCancel(context.Background())

	// cancel function can be called multiple times
	doCancelFunc()
	doCancelFunc()
}
