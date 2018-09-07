package mreflect

import (
	"log"
	"testing"

	"github.com/jinzhu/copier"
)

type helloCP struct {
	name string
	vals []string
}

func TestCopierSlice(t *testing.T) {
	h := helloCP{
		name: "h",
		vals: []string{"ab", "cd"},
	}
	var hh helloCP
	copier.Copy(&hh, &h)
	copier.Copy(&hh.vals[0], &h.vals[0])

	// hh.vals = []string{"ef", "gh"}
	hh.vals[0] = "ef"
	hh.name = "abcdefgh"

	log.Printf("original object: %v \n copied object: %v", h, hh)
}

func TestShadowCopy(t *testing.T) {
	h := helloCP{
		name: "bob",
		vals: []string{"paris", "newyork"},
	}

	h1 := &h
	h2 := *h1
	h2.name = "tom"

	log.Printf("h1: %v, h2: %v", *h1, h2)
}
