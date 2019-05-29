package pool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrowableRing_Read(t *testing.T) {
	assert := assert.New(t)
	r := NewRing(2, true)

	type item struct {
		value int
	}

	_, ok := r.Read()
	assert.False(ok)

	v1 := &item{value: 1}
	v2 := &item{value: 2}

	r.Write(v1)
	r.Write(v2)

	// FIFO
	y, ok := r.Read()
	assert.True(ok)
	assert.Equal(y, v1)
}

func TestGrowableRing_Write(t *testing.T) {
	assert := assert.New(t)
	r := NewRing(1, true)

	type item struct {
		value int
	}

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

func TestFixedRing_Read(t *testing.T) {
	assert := assert.New(t)
	r := NewRing(2, false)

	type item struct {
		value int
	}

	_, ok := r.Read()
	assert.False(ok)

	v1 := &item{value: 1}
	v2 := &item{value: 2}

	r.Write(v1)
	r.Write(v2)

	// FIFO
	y, ok := r.Read()
	assert.True(ok)
	assert.Equal(y, v1)
}

func TestFixedRing_Write(t *testing.T) {
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
