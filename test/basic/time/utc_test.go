package time_test

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestTime_format(t *testing.T) {
	tm := time.Now()

	str := tm.UTC().Format("2006-01-02 15:04:05.000Z")

	fmt.Printf("hello: %v", str)

	fmt.Println(strings.ContainsAny("http: Server closed", " \n,"))
}
