package basic

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	s := "ABCDE"

	log.Printf("%v ==> %v\n", ([]byte)(s), s)
}

func TestStringCompare(t *testing.T) {
	a := "bb"
	b := "bb"

	log.Println(a == b)
}

func TestStringAndBytes(t *testing.T) {
	str := "12"
	b := []byte(str)

	log.Println(str, b)

	b = []byte{0x01, 0x02}
	str = string(b)

	log.Println(str, b)
}

func TestStringsFold(t *testing.T) {
	want := true
	got := strings.EqualFold("Get", "GET")

	if want != got {
		t.Errorf("strings.EqualFold(%q, %q) = %v, want %v", "Get", "GET", got, want)
	}
}

func TestStringFromNIL(t *testing.T) {
	var a []byte
	a = nil
	b := string(a)
	log.Println(b == "")
}

func TestStringTrim(t *testing.T) {
	str := "     hello   "
	log.Println(strings.TrimSpace(str))
}

func TestStringConvert(t *testing.T) {
	aa := 10
	log.Println(strconv.Itoa(aa))
}

func TestStringTrimSuffix(t *testing.T) {
	aa := "aaa/bbb"
	bb := strings.TrimPrefix(aa, "aaa/")

	log.Println(aa, bb)
}

func TestStringAffix(t *testing.T) {
	log.Println(strings.HasPrefix("/aaa", "/"))
	log.Println(strings.HasSuffix("bbbb/", "/"))
}

func TestStringPointerConvert(t *testing.T) {
	var strptr *string
	var str string
	str = "hello world"

	// Must be initialized before using
	strptr = &str
	*strptr = "world"

	log.Println(*strptr)
	log.Println(str)
}

func TestStringType(t *testing.T) {
	c := '/'
	s := "/"
	log.Println(fmt.Sprintf("%T", c))
	log.Println(fmt.Sprintf("%T", s))

	log.Println(string(c))
}

func TestStringSplit(t *testing.T) {
	str := ""
	strSlice := strings.Split(str, ",")
	log.Println(strSlice)
}
