package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

var typeEntry = reflect.TypeOf(Entry{})
var ptrTypeEntry = reflect.TypeOf(&Entry{})

type Entry struct {
	Key   string
	Value string
}

func newEntryValue1() *Entry {
	val := reflect.New(typeEntry).Interface()

	return val.(*Entry)
}

func newEntryValue2() *Entry {
	val := reflect.New(ptrTypeEntry.Elem()).Interface()

	return val.(*Entry)
}

func BenchmarkReflect(b *testing.B) {
	var val *Entry
	cases := []struct {
		name string
		fn   func() *Entry
	}{

		{"value-type", newEntryValue1},
		{"ptr-type", newEntryValue2},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				val = c.fn()
			}
		})

	}

	_ = val

}

func TestReflect_Value(t *testing.T) {
	var typ = reflect.TypeOf(Entry{})
	var v1 = reflect.New(typ).Interface()
	var v2 = reflect.New(typ).Elem().Interface()
	fmt.Printf("type - origin: %v v1: %v v2: %v, value - v1: %v, v2: %v\n",
		typ, reflect.TypeOf(v1), reflect.TypeOf(v2), v1, v2)

	var typ2 = reflect.TypeOf(&Entry{})
	var v3 = reflect.New(typ2).Interface()
	var v4 = reflect.New(typ2.Elem()).Interface()
	var v5 = reflect.New(typ2).Elem().Interface()

	fmt.Printf("type2: %v, v3: %v, v4: %v, v5: %v\n",
		typ2, v3, v4, v5)

	var typ3 = reflect.TypeOf(1)
	fmt.Printf("type3: %v\n", typ3)
}
