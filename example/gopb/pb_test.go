package gopb

import (
	"encoding/hex"
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestProtobuf(t *testing.T) {
	req := &SearchRequest{
		Query:         "hello",
		PageNumber:    1,
		ResultPerPage: 50,
	}

	out, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("serialize: [%v]\n", hex.Dump(out))

	nreq := &SearchRequest{}
	err = proto.Unmarshal(out, nreq)
	if err != nil {
		panic(err)
	}

	fmt.Printf("deserialize: [%v]\n", nreq)
}
