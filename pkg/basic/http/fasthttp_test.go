package http

import (
	"log"
	"testing"
	"time"

	"github.com/valyala/fasthttp"
)

type stat struct {
	size    int64
	latency int64
}

func TestFastHTTP(t *testing.T) {
	pureURL := "http://localhost:9002/pulsar/common/1"

	client := fasthttp.Client{
		MaxConnsPerHost: 100,
	}

	queue := make(chan *stat, 2048)

	for i := 0; i < 2; i++ {
		go func() {

			for k := 0; k < 10000; k++ {
				req := fasthttp.AcquireRequest()
				resp := fasthttp.AcquireResponse()

				req.SetRequestURI(pureURL)

				startTime := time.Now()
				err := client.Do(req, resp)

				if err == nil {
					latency := int64(time.Since(startTime) / time.Millisecond)
					body := make([]byte, len(resp.Body()))
					copy(body, resp.Body())
					queue <- &stat{int64(len(body)), latency}
				}

				fasthttp.ReleaseRequest(req)
				fasthttp.ReleaseResponse(resp)
			}
		}()
	}

	go func() {
		var batchSize int64 = 100
		var latency int64
		var count int64
		seq := 0
		for {
			select {
			case t := <-queue:
				if count == batchSize {
					log.Printf("==seq: %d, avg latency: %d ms", seq, latency/batchSize)

					latency = 0
					count = 0
					seq++
				}

				count++
				latency += t.latency
			}
		}
	}()

	time.Sleep(100 * time.Second)
}
