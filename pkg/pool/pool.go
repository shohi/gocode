package pool

import "sync"

type Pool interface {
	Put(interface{})
	Get() interface{}
}

type ringPool struct {
	lock sync.Mutex
	r    *Ring
	fn   func() interface{} // factory
}

func NewRingPool(fn func() interface{}) Pool {
	return &ringPool{
		r:  NewRing(1024),
		fn: fn,
	}
}

func (p *ringPool) Get() interface{} {
	p.lock.Lock()
	defer p.lock.Unlock()

	if v, ok := p.r.Read(); ok {
		return v
	}

	return p.fn()
}

func (p *ringPool) Put(v interface{}) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.r.Write(v)
}
