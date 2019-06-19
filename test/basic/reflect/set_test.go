package reflect_test

import (
	"log"
	"reflect"
	"testing"
)

func fill(ret interface{}, val interface{}) {
	k := reflect.TypeOf(ret).Kind()
	rv := reflect.ValueOf(ret).Elem()

	log.Printf("#### type: %v", k)
	log.Printf("#### value: %v", rv.Kind())
	log.Printf("#### rcv: %v", reflect.TypeOf(val).Kind())

	reflect.ValueOf(ret).Elem().Set(reflect.ValueOf(val))
}

func TestReflect_Fill(t *testing.T) {

	var s string
	fill(&s, "hello")

	log.Printf("===> fill string: %v", s)
}
