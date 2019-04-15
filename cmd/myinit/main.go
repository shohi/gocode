package main

import "log"

var testMain = false

func init() {
	log.Printf("init in main")
}

func main() {
	log.Printf("run main.go, testMain: %v", testMain)
}
