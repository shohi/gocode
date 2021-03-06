package time_test

import (
	"log"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDuration(t *testing.T) {
	d := 1000 * time.Second
	log.Println(strings.ToUpper(d.String()))
}

func TestParseDuration(t *testing.T) {
	durationStr := "10s"
	log.Println(time.ParseDuration(durationStr))

	// 100 year
	durationStr = "876000h"
	log.Println(time.ParseDuration(durationStr))

	durationStr = "10"
	log.Println(time.ParseDuration(durationStr))
}

func TestDurationZeroValue(t *testing.T) {
	var d time.Duration
	log.Printf("zero value of duration is %v", d)
}

func TestDurationFromFloat64(t *testing.T) {
	//
	aa := 17.58
	log.Printf("duration: %v", int64(aa))
	log.Printf("duration: %v", time.Duration(aa)*time.Second)
}

func TestDurationConversion(t *testing.T) {
	start := time.Now()
	time.Sleep(100 * time.Millisecond)
	end := time.Now()
	log.Printf("duration: %v, %v, %d", end.Sub(start), end.Sub(start)/time.Millisecond, end.Sub(start)/time.Millisecond)
}

func TestDurationDivide(t *testing.T) {
	t1 := 1 * time.Second
	t2 := 1 * time.Millisecond

	log.Printf("ms ==> %v", float64(t1/t2))
	log.Printf("ms ==> %v", float64(t1)/float64(t2))
}

func TestDuration_Parse_LeadingSpace(t *testing.T) {
	durStr := "     10s"
	d, err := time.ParseDuration(durStr)
	assert.NotNil(t, err)
	log.Printf("====> duration: %v, error: %v", d, err)
}

func TestDuration_format(t *testing.T) {
	d := 10 * time.Second
	log.Printf("duration: %v", d)
}

func TestDurationSub(t *testing.T) {
	startT := time.Now()
	time.Since(startT)
}

func TestDuration_Cast(t *testing.T) {
	var d float64 = 1000.01
	log.Printf("duration: %v", time.Duration(d))
}
