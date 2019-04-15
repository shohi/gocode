package main

import (
	"log"
	"testing"
)

func init() {
	testMain = true
	log.Printf("init in test")
}

func TestMain(t *testing.T) {
	log.Printf("test, testMain: %v", testMain)
}
