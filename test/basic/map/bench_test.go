package map_test

import (
	"testing"
)

func addKV(m map[int]interface{}, cnt int) {
	for i := 0; i < cnt; i++ {
		m[i] = i
	}
}

var globalVal map[int]interface{}

func BenchmarkMap_Alloc(b *testing.B) {
	m := make(map[int]interface{}, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		addKV(m, 10)
	}

	globalVal = m
}
