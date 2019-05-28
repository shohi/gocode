package pool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRing_Growable(t *testing.T) {
	assert := assert.New(t)
	r := NewRing(1, true)

	type item struct {
		value int
	}

	_, ok := r.Read()
	assert.False(ok)

	v1 := &item{value: 1}
	v2 := &item{value: 2}

	r.Write(v1)
	r.Write(v2)

	y1, ok := r.Read()
	assert.True(ok)
	assert.Equal(y1, v1)

	y2, ok := r.Read()
	assert.True(ok)
	assert.Equal(y2, v2)

	_, ok = r.Read()
	assert.False(ok)
}

func TestRing_NonGrowable(t *testing.T) {
	assert := assert.New(t)
	r := NewRing(1, false)

	type item struct {
		value int
	}

	_, ok := r.Read()
	assert.False(ok)

	v1 := &item{value: 1}
	v2 := &item{value: 2}

	r.Write(v1)
	r.Write(v2)

	y, ok := r.Read()
	assert.True(ok)
	assert.Equal(y, v2)

	_, ok = r.Read()
	assert.False(ok)
}
