package slice_test

import (
	"log"
	"strings"
	"testing"
)

func TestInitialization(t *testing.T) {
	s := make([]int, 10)

	log.Println(strings.Repeat("*", 10)+"Before ===>", len(s))
	for k, v := range s {
		log.Printf("%d ===> %v", k, v)
	}

	//
	s = append(s, 10)
	log.Println(strings.Repeat("*", 10)+"after ===> ", len(s))
	for k, v := range s {
		log.Printf("%d ===> %v", k, v)
	}
}

func TestInitWithNil(t *testing.T) {
	var bs []byte
	bs = nil
	log.Printf("byte slice: %v", bs)
}

func TestSliceCapAfterTruncate(t *testing.T) {
	bs := []byte("hello")
	log.Printf("before: len - [%d], cap - [%d]", len(bs), cap(bs))

	bs = bs[:0]
	log.Printf("after: len - [%d], cap - [%d]", len(bs), cap(bs))
}

func TestSliceMutation(t *testing.T) {
	bs := []byte("hello")
	nb := bs[:3]

	log.Printf("before: content - [%v]", string(bs))

	nb[0] = 'w'

	log.Printf("after: content - [%v]", string(bs))
	nbb := bs[:100]
	log.Printf("long: content - [%v], len - [%d], cap - [%d]", string(nbb), len(nbb), cap(nbb))
}

func TestSliceAssign(t *testing.T) {
	type dataType struct {
		val int
	}

	// case 1 - use equal
	// NOTE: change will propagate
	aa := []dataType{dataType{val: 1}, dataType{val: 2}}

	bb := aa
	bb[0].val = 10
	log.Printf("use equal: a => %v, b => %v", aa, bb)

	// case 2 - use append
	// NOTE: change does not propagate
	var cc []dataType
	cc = append(cc, aa...)
	cc[0].val = 20

	log.Printf("use append: a => %v, c => %v", aa, cc)
}

func TestSliceWithElementChange(t *testing.T) {
	type dataType struct {
		val int
	}

	aa := []dataType{dataType{val: 1}, dataType{val: 2}}

	log.Printf("last aa: %v", aa)

	for _, d := range aa {
		d.val += 100
	}

	// NOTE: aa not changed
	log.Printf("current aa: %v", aa)

	// direct change
	aa[0].val += 200
	log.Printf("direct aa: %v", aa)
}
