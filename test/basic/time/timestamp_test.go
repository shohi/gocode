package time_test

import (
	"log"
	"testing"
	"time"
)

func TestTime_UnixTimestamp(t *testing.T) {
	tt := time.Now()

	log.Printf("unix timestamp: %v", tt.Unix())
}
