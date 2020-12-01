package serde_test

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	s := "aHR0cDovL2Nkbi1lYy1uYXAtMDAzLTA0LmRhc2hjaWYtd2VzdC1mbC54Y3IuY29tY2FzdC5uZXQvV0VEVURfSERfV0VTVGRGTF81NjQ1XzBfNDk5NjEzNzM3Mjc4MDAzMTE2My9yb290X2F1ZGlvX3ZpZGVvNC80ODM5NTc4OTI0LnRz"
	s = "aHR0cDovL2Nkbk1lYy1sYWstMzAxLmNsYy1jaWYtZGNmLnhjci5jb21jYXN0Lm5ldC9TVFpIRF9IRF9OQVRfMTE3MTZfMF83NTA5NDI0NDkyNDE2MDg5MTYzL3Jvb3RfYXVkaW8xMDMvMTYzMjM2NzgwNC50cw=="

	data, err := base64.StdEncoding.DecodeString(s)
	fmt.Printf("====> data: %v, err: %v\n", string(data), err)

	// data2, err := base64.RawStdEncoding.DecodeString(s)
	// fmt.Printf("====> data2: %v, err: %v\n", string(data2), err)

	data3, err := base64.URLEncoding.DecodeString(s)
	fmt.Printf("====> data3: %v, err: %v\n", string(data3), err)
}
