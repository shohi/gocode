package http

import (
	"context"
	"log"
	"net"
	"sync"
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

func TestFastHTTP_Server(t *testing.T) {
	s := fasthttp.Server{}
	var wg sync.WaitGroup

	addr := "0.0.0.0:0"
	ln, err := net.Listen("tcp4", addr)

	if err != nil {
		log.Fatalf("error in net.Listen: %s", err)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := s.Serve(ln)
		t.Fatalf("err: %v", err)
	}()

	time.Sleep(100 * time.Millisecond)

	log.Printf("address: %v", ln.Addr())
}

func TestFastHTTP_Client(t *testing.T) {
	url := "http://0.0.0.1:9002"

	client := fasthttp.Client{}
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	respCh := make(chan error, 1)
	var startTime = time.Now()
	go func() {
		req.SetRequestURI(url)
		err := client.Do(req, resp)
		respCh <- err
	}()

	var err error
	select {
	case <-ctx.Done():
		err = ctx.Err()
	case err = <-respCh:
	}

	log.Printf("===> err: %v, duration: %v", err, time.Since(startTime))

}

func TestFastHTTP_Headers(t *testing.T) {
	req := fasthttp.AcquireRequest()
	req.Header.Set("k1", "v1")
	req.Header.Set("k2", "v2")

	req.Header.VisitAll(func(key, value []byte) {
		log.Printf("%v: %v", string(key), string(value))
	})
}
