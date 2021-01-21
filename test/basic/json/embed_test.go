package json_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

type Inner struct {
	Key   string
	Value string
}

type Outer struct {
	Inner
}

func TestJson_Embed(t *testing.T) {
	var v = &Outer{}
	v.Key = "hello"
	v.Value = "world"

	b, _ := json.MarshalIndent(v, "", "  ")

	fmt.Printf("====> content: %v\n", string(b))

	var vv = Outer{}
	_ = json.Unmarshal(b, &vv)
	fmt.Printf("====> new content: %+v\n", vv)
}

func TestMarshal_Embed(t *testing.T) {
	type inner struct {
		Name string
		Val  string // index [0 1]
	}

	type Inner2 struct {
		Key string
		Val string // index [1 1]
	}

	type Outer struct {
		inner
		Inner2
		Key string
	}

	var a Outer
	// a.Val = "hello"
	a.inner = inner{"n1", "1"}
	a.Inner2 = Inner2{"v2", "v2"}
	a.Key = "out-key"

	data, err := json.Marshal(a)

	fmt.Printf("json: [%v], err: %v\n", string(data), err)
	// Output:
	// {"Name":"n1","Key":"out-key"}
}

func TestMarshal_EmbedPointer(t *testing.T) {
	type Inner struct {
		Key string
		Val string
	}

	type Outer struct {
		*Inner
		// NOTE: compile failed
		*context.Context
		Val string
	}

	var a Outer

	var inner = &Inner{Key: "k2", Val: "v2"}
	a.Inner = inner
	a.Val = "v1"
	data, err := json.Marshal(a)

	fmt.Printf("json: %v, err: %v\n", string(data), err)
	// Output:
	// {"inner":{"Val":"v2"},"ctx":0,"Val":"v1"}
}

func TestMarshal_EmbedInterface(t *testing.T) {
	type Inner struct {
		Val string
	}

	var a struct {
		Inner           `json:"inner"`
		context.Context `json:"ctx"`
		Item            struct {
			Count int
		}
		Val string
	}

	a.Context = context.Background()
	a.Inner.Val = "v2"
	a.Val = "v1"
	a.Item.Count = 1

	data, err := json.Marshal(a)

	fmt.Printf("json: %v, err: %v\n", string(data), err)
	// Output:
	// {"inner":{"Val":"v2"},"ctx":0,"Val":"v1"}
}

// https://medium.com/@xcoulon/nested-structs-in-golang-2c750403a007
func TestMarshal_Indent(t *testing.T) {
	type Config struct {
		Server struct {
			Host string `json:"host"`
			Port string `json:"port"`
		} `json:"server"`
		Postgres struct {
			Host     string `json:"host"`
			User     string `json:"user"`
			Password string `json:"password"`
			DB       string `json:"db"`
		} `json:"database"`
	}

	config := Config{
		Server: struct {
			Host string `json:"host"`
			Port string `json:"port"`
		}{
			Host: "localhost",
			Port: "8080",
		},
		Postgres: struct {
			Host     string `json:"host"`
			User     string `json:"user"`
			Password string `json:"password"`
			DB       string `json:"db"`
		}{
			Host:     "localhost",
			User:     "db_user",
			Password: "supersecret",
			DB:       "my_db",
		},
	}

	data, _ := json.MarshalIndent(config, "$", "a")
	fmt.Printf("json:\n%v\n", string(data))

	// Output:
	//
	// {
	// $a"server": {
	// $aa"host": "localhost",
	// $aa"port": "8080"
	// $a},
	// $a"database": {
	// $aa"host": "localhost",
	// $aa"user": "db_user",
	// $aa"password": "supersecret",
	// $aa"db": "my_db"
	// $a}
	// $}

}
