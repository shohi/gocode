package test_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssert_Regex(t *testing.T) {
	assert := assert.New(t)

	assert.Regexp("start", "start here")
}
