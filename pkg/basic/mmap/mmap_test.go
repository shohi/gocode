package mmap

import (
	"log"
	"testing"
)

func TestGetFromNilMap(t *testing.T) {
	var m map[string]string

	val, ok := m["hello"]

	log.Printf("get value from nil -- [ok]: %v", ok)
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

func TestTraverseNilMap(t *testing.T) {
	var m map[string]int
	for k, v := range m {
		log.Printf("key ==> %v, value ==> %v", k, v)
	}
}

func TestKeyExistence(t *testing.T) {
	m := map[string]string{
		"a": "a",
		"b": "b",
		"c": "",
	}

	val, ok := m["d"]
	log.Printf("value: %v, existence: %v", val, ok)

	val, ok = m["c"]
	log.Printf("value: %v, existence: %v", val, ok)
}
