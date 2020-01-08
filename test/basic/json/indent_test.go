package json_test

import (
	"encoding/json"
	"log"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestJSON_indent_v1(t *testing.T) {
	data := struct {
		Name  string
		Value string
	}{"hello", "world"}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	// NOTE: jsoniter does not support other prefix except empty string.
	// And only space indent is allowed.
	var prefix = ""

	// 4-spaces
	var indent = strings.Repeat(" ", 4)
	bs, _ := json.MarshalIndent(&data, prefix, indent)

	log.Printf("json string:\n%v\n", string(bs))
}

func TestJSON_indent_v2(t *testing.T) {
	data := struct {
		Name  string
		Value string
	}{"hello", "world"}

	var prefix = "@"

	// 4-spaces
	var indent = strings.Repeat("\t", 2)
	bs, _ := json.MarshalIndent(&data, prefix, indent)

	log.Printf("json string:\n%v\n", string(bs))
}
