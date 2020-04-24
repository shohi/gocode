package http_test

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAuthBasic(t *testing.T) {
	auth := "cm9vdDpibHU="
	payload, _ := base64.StdEncoding.DecodeString(auth)
	payload2, _ := base64.URLEncoding.DecodeString(auth)

	fmt.Printf("=====> auth: [%v], url decode: [%v]", string(payload), string(payload2))
}
