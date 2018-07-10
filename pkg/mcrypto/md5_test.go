package mcrypto

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"testing"
)

func TestMd5Sum(t *testing.T) {
	var data []byte
	md5Sum := md5.Sum(data)
	md5str := hex.EncodeToString(md5Sum[:])

	log.Printf("md5sum ==> %v", md5str)
	data = []byte{}

	md5Sum = md5.Sum(data)
	md5str = hex.EncodeToString(md5Sum[:])

	log.Printf("md5sum ==> %v", md5str)
}
