package date_test

import (
	"log"
	"testing"

	"github.com/jinzhu/now"
)

func TestNow_Parse(t *testing.T) {
	// Not supported
	timeStr := "2006-01-02 15:04:05,000Z"
	now.TimeFormats = append(now.TimeFormats, timeStr)
	tm, err := now.Parse("2008-01-02 15:04:05,006Z")

	log.Printf("time: %v, err: %v", tm, err)

	// timeStr2 := "2006-01-02 15:04:05.003Z"

	// log.Printf("format: %v", time.Now().UTC().Format(timeStr2))
	// tm2, err2 := time.Parse(timeStr2, timeStr2)
	// log.Printf("time: %v, err: %v, format: %v", tm2, err2, time.Now().Format(timeStr2))

}
