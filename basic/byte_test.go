package basic

import (
	"log"
	"strconv"
	"testing"
)

//
func TestByteBinaryPrint(t *testing.T) {
	var b = byte(0x01)
	log.Println(b)
	log.Printf("%b", b)
	log.Println(strconv.FormatInt(int64(b), 2))
}

func TestByteLiteral(t *testing.T) {
	b := byte(0x21)
	log.Println(b)

	bb := []byte{0x00, 0x01}
	log.Println(bb)
}
