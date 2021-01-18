package json

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestJsonMarshal_Context(t *testing.T) {
	var ctx context.Context = context.WithValue(context.Background(), "key", "value")
	fmt.Printf("context type: %v, elem: %v\n",
		reflect.TypeOf(ctx), reflect.TypeOf(ctx).Elem())

	data, _ := json.Marshal(ctx)
	// Output: ?
	fmt.Printf("content: [%v]\n", string(data))
}

func TestType_Marshaler(t *testing.T) {
	var m json.Marshaler = nil
	typ1 := reflect.TypeOf(m)
	fmt.Printf("m: %+v\n", typ1)

	typ := reflect.TypeOf((*json.Marshaler)(nil))
	fmt.Printf("type: %+v\n", typ)

	marshalerType := typ.Elem()
	fmt.Printf("marshaler type: %+v\n", marshalerType)
}
