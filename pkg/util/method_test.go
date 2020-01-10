package util

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hashicorp/memberlist"
)

type MyTest struct {
	val string
}

func (m *MyTest) Update(v string) {
	m.val = v
}

func (m MyTest) String() string {
	return m.val
}

func TestExtractPublicMethods(t *testing.T) {
	var m memberlist.Memberlist
	sigs := ExtractPublicMethods(&m)

	for _, sig := range sigs {
		fmt.Printf("%v\n", sig)
	}
}

func extractField(v interface{}) {
	vv := reflect.ValueOf(v).Elem()
	t := vv.Type()
	if t.Kind() != reflect.Struct {
		panic("should be struct")
	}

	fmt.Println(vv.NumField())
	for i, l := 0, vv.NumField(); i < l; i++ {
		field := vv.Field(i)
		fmt.Println(field)
	}
}

func TestExtractField(t *testing.T) {
	extractField(&MyTest{})
}

func TestSignature_FromFunc(t *testing.T) {
	sig := funcToSignature(extractField)
	fmt.Printf("signature: %+v\n", sig)
}

func TestSignature_FromMethod(t *testing.T) {
	var m memberlist.Memberlist
	vType := reflect.TypeOf(&m)
	fmt.Println("method: ", vType.NumMethod())

	sig := methodToSignature(vType.Method(0))
	fmt.Printf("signature: %+v\n", sig)
}
