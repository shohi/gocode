package basic

import (
	"log"
	"strings"
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	aa := time.Now()
	time.Sleep(1 * time.Second)
	log.Println(aa)
	bb := time.Since(aa)
	bb = bb - (bb % time.Second)
	log.Println(bb)
}

func TestDuration(t *testing.T) {
	d := 1000 * time.Second
	log.Println(strings.ToUpper(d.String()))
}

func TestUnixTimestamp(t *testing.T) {
	log.Println(time.Now().Unix())
}

func TestTimeString(t *testing.T) {
	log.Println(time.Now().String())
}

// Duration must come with unit
func TestParseDuration(t *testing.T) {
	durationStr := "10s"
	log.Println(time.ParseDuration(durationStr))

	durationStr = "10"
	log.Println(time.ParseDuration(durationStr))
}
