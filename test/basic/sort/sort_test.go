package sort_test

import (
	"fmt"
	"log"
	"sort"
	"testing"
)

func TestSortInDesc(t *testing.T) {
	a := []int{1, 2, 4, 3}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	log.Printf("sorted slice: %v", a)
}

func TestSortInAsc(t *testing.T) {
	ints := []int{1, 2, 5, 6}
	sort.Ints(ints)
	fmt.Println(ints)
	idx := sort.Search(len(ints), func(i int) bool { return (ints[i] >= 10) })
	fmt.Println(idx)
}
