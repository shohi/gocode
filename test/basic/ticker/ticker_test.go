package ticker_test

import (
	"log"
	"testing"
	"time"
)

func TestTickerStop_Behaviour(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	count := 0

	for {
		select {
		case t := <-ticker.C:
			count++
			log.Printf("ticker - %v - %v\n", count, t)
			// NOTE: after 5 times, panic occurs.
			// INFO: fatal error: all goroutines are asleep - deadlock!
			if count == 5 {
				ticker.Stop()
			}
		}
	}
}

func TestTickerStop_Mulitple(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			tt := <-ticker.C
			log.Printf("ticker - %v\n", tt)
		}
	}()

	time.Sleep(3 * time.Second)

	// NOTE: ticker supports mulitple `Stop` calls.
	ticker.Stop()
	ticker.Stop()
}
