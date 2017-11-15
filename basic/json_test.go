package basic

import (
	"encoding/json"
	"log"
	"testing"
)

type hello struct {
	A int    `json:"fieldA"`
	B string `json:"fieldB"`
}

func TestMarshal(t *testing.T) {
	bs, _ := json.Marshal(hello{10, "hello"})
	// log.Println(string(bs))

	bs, _ = json.MarshalIndent(hello{10, "hello"}, "", "  ")
	log.Println(string(bs))
}

func TestRawMessag(t *testing.T) {
	aa := json.RawMessage("hello")
	log.Println(aa)
	log.Println(string(aa))
}
