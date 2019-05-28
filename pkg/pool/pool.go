package pool

import "sync"

// A Pool is the interface to reuse objects.
// After `Put` is called, object MUST not be used.
type Pool interface {
	Put(interface{})
	Get() interface{}
}

type pool struct {
	lock sync.Mutex
	r    *Ring
	fn   func() interface{} // factory
}

// NewPool creates a new growable pool
func NewPool(cap int, fn func() interface{}) Pool {
	return &pool{
		r:  NewRing(cap, true),
		fn: fn,
	}
}

// NewFixedPool creates a new pool with fixed size.
func NewFixedPool(cap int, fn func() interface{}) Pool {
	return &pool{
		r:  NewRing(cap, false),
		fn: fn,
	}
}

// Get get an object from pool, create a new one if non exist.
func (p *pool) Get() interface{} {
	p.lock.Lock()
	defer p.lock.Unlock()

	if v, ok := p.r.Read(); ok {
		return v
	}

	return p.fn()
}

// Put returns object to the pool for reuse.
func (p *pool) Put(v interface{}) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.r.Write(v)
}
