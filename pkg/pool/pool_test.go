package pool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPool_Fixed(t *testing.T) {
	assert := assert.New(t)

	type item struct {
		value int
	}
	p := NewFixedPool(1, func() interface{} {
		return &item{value: 10}
	})

	v1 := p.Get().(*item)

	v2 := p.Get().(*item)
	v2.value = 20

	p.Put(v1)
	p.Put(v2)

	y := p.Get().(*item)

	assert.Equal(y.value, 20)
}

func TestPool_Growable(t *testing.T) {
	assert := assert.New(t)

	type item struct {
		value int
	}
	p := NewPool(1, func() interface{} {
		return &item{value: 10}
	})

	v1 := p.Get().(*item)
	v2 := p.Get().(*item)
	v2.value = 20

	p.Put(v1)
	p.Put(v2)

	y := p.Get().(*item)
	assert.Equal(y.value, 10)
}

// copy from `sync.Pool`
func TestPoolStress(t *testing.T) {
	const P = 10
	N := int(100)

	p := NewFixedPool(P, func() interface{} {
		return 0
	})

	done := make(chan bool)
	for i := 0; i < P; i++ {
		go func() {
			var v interface{} = 0
			for j := 0; j < N; j++ {
				if v == nil {
					v = 0
				}
				p.Put(v)
				v = p.Get()
				if v != nil && v.(int) != 0 {
					t.Errorf("expect 0, got %v", v)
					break
				}
			}
			done <- true
		}()
	}
	for i := 0; i < P; i++ {
		<-done
	}
}
