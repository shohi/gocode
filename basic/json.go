package main

import "fmt"
import "encoding/json"

type hello struct {
	A int    `json:"fieldA"`
	B string `json:"fieldB"`
}

func marshal() {
	bs, _ := json.Marshal(hello{10, "hello"})
	// fmt.Println(string(bs))

	bs, _ = json.MarshalIndent(hello{10, "hello"}, "", "  ")
	fmt.Println(string(bs))
}

func rawmessage() {
	aa := json.RawMessage("hello")
	fmt.Println(aa)
	fmt.Println(string(aa))
}

func main() {
	rawmessage()
}
