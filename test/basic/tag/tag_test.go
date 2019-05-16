package tag_teset

import (
	"log"
	"reflect"
	"testing"
)

type MyType struct {
	A int `my:"hello"` // <key, value> in tag must keep in the format of [key="value"]
}

func printField() {
	var mt MyType
	vv := reflect.TypeOf(mt)
	// v := vv.Elem()
	v := vv
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := field.Tag
		log.Printf("field ==> %v, tag ==> %v, [my]tag ==> %v", field, tag, tag.Get("my"))
	}
}

func TestTag(t *testing.T) {
	printField()
}
