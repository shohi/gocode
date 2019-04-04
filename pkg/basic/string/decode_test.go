package string_test

import (
	"bytes"
	"log"
	"strconv"
	"testing"
)

//
// s = "3[a]2[bc]", return "aaabcbc".
// s = "3[a2[c]]", return "accaccacc".
// s = "2[abc]3[cd]ef", return "abcabccdcdcdef".

func decodeString(s string) string {
	type Item struct {
		Num int
		Val bytes.Buffer
	}

	stack := []*Item{&Item{}}
	counting := false
	var numStart int

	for i, ch := range s {
		switch {
		case ch >= '0' && ch <= '9':
			if !counting {
				var newBuf bytes.Buffer
				item := &Item{Val: newBuf}
				stack = append(stack, item)
				numStart = i
				counting = true
			}
		case ch == '[':
			counting = false
			numStr := s[numStart:i]
			numInt64, _ := strconv.ParseInt(numStr, 10, 32)
			stack[len(stack)-1].Num = int(numInt64)
		case ch == ']':
			tmpNum, tmpVal := stack[len(stack)-1].Num, stack[len(stack)-1].Val.Bytes()
			stack = stack[:len(stack)-1]
			for j := 0; j < tmpNum; j++ {
				stack[len(stack)-1].Val.Write(tmpVal)
			}
		default:
			stack[len(stack)-1].Val.WriteRune(ch)
		}
	}
	return stack[0].Val.String()
}

func TestDecode(t *testing.T) {
	ptn := []string{"3[a]2[bc]", "3[a2[c]]", "2[abc]3[cd]ef"}

	for _, p := range ptn {
		log.Printf("format: %v, output: %v", p, decodeString(p))
	}
}

func TestNil(t *testing.T) {
	type myStruct struct {
		Val string
	}

	ms := &myStruct{}

	if ms != nil {
		log.Printf("hello world")
	}
}
