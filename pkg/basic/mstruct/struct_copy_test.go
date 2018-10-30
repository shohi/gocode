package mstruct_test

import (
	"log"
	"testing"
)

type MyResult struct {
	no    int
	grade []int
	title string
}

func TestStructCopyWithPointer(t *testing.T) {

	res := MyResult{
		no:    1,
		grade: []int{80, 90, 100},
		title: "math",
	}

	// copy result
	res2 := res
	res2.grade[0] = 81
	res2.title += "english"

	log.Printf("original value: %v", res)
	log.Printf("copied value: %v", res2)
}
