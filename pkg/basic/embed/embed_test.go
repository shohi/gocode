package embed_test

import (
	"log"
	"testing"
)

type Info struct {
	Name  string
	Value string
}
type Msg struct {
	Name  string
	Value string
}

type Event struct {
	Info
	Msg
}

func TestEmbededOrder(t *testing.T) {
	var event Event
	// NOTE: can't be set, err: ambiguous selector event.Name
	// event.Name = "hello"
	event.Info.Name = "hello"
	event.Msg.Value = "world"

	log.Printf("%v", event)
}
