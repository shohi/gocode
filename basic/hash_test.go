package basic

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"testing"
)

func TestMd5(t *testing.T) {
	str := "hello world"

	md5sum := md5.Sum([]byte(str))
	md5sumStr := hex.EncodeToString(md5sum[:])

	log.Println(md5sumStr)

}
