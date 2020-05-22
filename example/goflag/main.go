package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var interval time.Duration
	flag.DurationVar(&interval, "interval", 0, "sync interval")

	flag.Parse()

	fmt.Printf("interval: %v\n", interval)
}
