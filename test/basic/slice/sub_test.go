package slice_test

import (
	"log"
	"testing"
)

func TestGetLargerSliceFromSmallerOne(t *testing.T) {

	s := []int{1, 2, 3, 4, 5, 6}
	ss := s[:0]
	ss = append(ss, 10)

	sss := ss[0:cap(s)]

	log.Printf("==> ss: {%v}, sss: {%v}", ss, sss)
}

func TestSubslice_ExtendLen(t *testing.T) {
	arr := []int{1, 2, 3, 4}

	// NOTE: last index is exclusive
	log.Println(arr[1:len(arr)])
}
