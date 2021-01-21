package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONMarshal_Tags(t *testing.T) {
	type AttrType map[string]string

	var a struct {
		ID         string   `json:"id"`
		Setting    string   `json:"-"`
		Commits    int      `json:",string"`
		Attr       AttrType `json:",omitempty"`
		pullRquest int
	}

	a.ID = "<1>"
	a.Setting = "default"
	a.Commits = 100
	a.pullRquest = 3

	data, err := json.Marshal(a)
	fmt.Printf("json: %v, err: %v\n", string(data), err)

	// Output:
	// {"id":"\u003c1\u003e","Commits":"100"}
}
