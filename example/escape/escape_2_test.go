package escape_test

import (
	"fmt"
	"testing"
)

type user struct {
	name  string
	email string
}

func TestEscape_Frame_1(t *testing.T) {
	var str = "hello"
	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", &u2)
	fmt.Printf("%v-%v-%v\n", &u2, str, 8)
}

//go:noinline
func createUserV1() user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V1", &u)
	return u
}

//go:noinline
func createUserV2() *user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V2", &u)
	return &u
}
