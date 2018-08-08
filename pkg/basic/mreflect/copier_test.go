package mreflect

import (
	"log"
	"testing"

	"github.com/jinzhu/copier"
)

type Contact struct {
	email string
	phone int
}

type Person struct {
	name      string
	age       int
	interests []string
	contacts  []*Contact
}

func TestCopier_Pointer(t *testing.T) {
	type result struct {
		value *string
	}

	val := "old"

	oldRes := &result{value: &val}
	newRes := &result{}
	copier.Copy(newRes, oldRes)

	val2 := "mid"
	newRes.value = &val2

	log.Printf("old: %v \n new: %v", *oldRes.value, *newRes.value)
}

func TestCopier_Slice(t *testing.T) {
	type result struct {
		vals []string
	}

	oldRes := &result{vals: []string{"hello", "world"}}
	newRes := &result{}

	copier.Copy(newRes, oldRes)
	newRes.vals = []string{"hi", "wo"}

	log.Printf("old: %v \n new: %v", oldRes, newRes)
}
