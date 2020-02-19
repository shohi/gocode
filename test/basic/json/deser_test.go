package json_test

import (
	"encoding/json"
	"log"
	"testing"
)

type myStruct struct {
	Key   string `json:"key"`
	Value []byte `json:"value"`
}

func TestDeserial_Bytes(t *testing.T) {
	m := myStruct{
		Key:   "tea",
		Value: []byte("dragonwell"),
	}

	data, err := json.MarshalIndent(&m, "", "  ")
	log.Printf("err: %v, data: [%v]", err, string(data))

	var m2 = myStruct{}
	err = json.Unmarshal(data, &m2)
	log.Printf("unmarshall error: %v\n", err)

	log.Printf("struct: %+v, value: %v\n", m2, string(m2.Value))

	// null bytes
	m3 := myStruct{Key: "key"}
	data, err = json.MarshalIndent(&m3, "", "  ")
	log.Printf("null ==> err: %v, data: [%v]", err, string(data))
}
