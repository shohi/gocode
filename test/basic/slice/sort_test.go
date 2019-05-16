package slice_test

import (
	"log"
	"sort"
	"testing"
)

type MyStruct struct {
	Age   int
	Value string
}

type Objects []MyStruct

func (o Objects) Len() int {
	return len(o)
}
func (o Objects) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}
func (o Objects) Less(i, j int) bool {
	return o[i].Age > o[j].Age
}

func TestSlice_Sort(t *testing.T) {
	objs := make(Objects, 3)
	objs[0] = MyStruct{80, "age-80"}
	objs[1] = MyStruct{90, "age-90"}
	objs[2] = MyStruct{100, "age-100"}

	sort.Sort(objs)

	for _, v := range objs {
		log.Printf("obj: [%v]", v)
	}

}
