package basic

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

	// hh.vals = []string{"ef", "gh"}
	hh.vals[0] = "ef"
	hh.name = "hh"

	log.Printf("original object: %v \n copied object: %v", h, hh)

}
