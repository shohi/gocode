package mmap

import (
	"log"
	"testing"
)

func TestGetFromNilMap(t *testing.T) {
	var m map[string]string

	val := m["hello"]

	log.Printf("m is empty: %v, val is \"\": %v", m == nil, val == "")
}

func TestMapKeyWithUint(t *testing.T) {
	m := make(map[uint]bool, 10)
	m[uint(10)] = true

	for k, v := range m {
		log.Printf("%d => %v", k, v)
	}

	log.Printf("non exist key: %v, value: %v", uint(20), m[uint(20)])
}
