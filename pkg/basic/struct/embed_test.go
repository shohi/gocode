package struct_test

import (
	"log"
	"testing"
)

func TestEmbeddedStruct_Init(t *testing.T) {
	type MyStruct struct {
		opts struct {
			name string
		}
	}

	myS := MyStruct{
		opts: struct{ name string }{name: "hello"},
	}

	log.Printf("mystruct: %v", myS)

}
