package mstruct_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

type strResp struct {
	v interface{}
}

func TestStructInitWithEmpty(t *testing.T) {
	assert := assert.New(t)
	res := &strResp{}

	log.Printf("res ==> %v", res)
	assert.NotNil(res)
}
