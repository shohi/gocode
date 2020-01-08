package sort_test

import (
	"fmt"
	"sort"
	"testing"
)

// refer, https://yourbasic.org/golang/how-to-sort-in-go/

type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func TestSortInterface(t *testing.T) {
	family := []Person{
		{"Alice", 23},
		{"Eve", 2},
		{"Bob", 25},
	}
	sort.Sort(ByAge(family))
	fmt.Println(family) // [{Eve 2} {Alice 23} {Bob 25}]
}

func TestSortInterface_V2(t *testing.T) {
	family := []Person{
		{"Alice", 23},
		{"Eve", 2},
		{"Bob", 25},
	}
	sort.Slice(family, func(i, j int) bool {
		return family[i].Name < family[j].Name
	})

	fmt.Println(family)
	//
	// [{Alice 23} {Bob 25} {Eve 2}]
}
