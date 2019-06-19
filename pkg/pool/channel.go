package pool

// use channel to implement pool
type chPool struct {
	items chan interface{}
	fn    func() interface{}
}

// NewFixedPool2 creates new fixed pool implemented by channel
func NewFixedPool2(cap int, fn func() interface{}) Pool {
	items := make(chan interface{}, cap)
	for k := 0; k < cap; k++ {
		items <- fn()
	}

	return &chPool{
		items: items,
		fn:    fn,
	}
}

func (p *chPool) Get() interface{} {
	select {
	case v := <-p.items:
		return v
	default:
		// If no item available, create a new one
		return p.fn()
	}
}

// Put puts v into pool without overwrite.
func (p *chPool) Put(v interface{}) {
	select {
	case p.items <- v:
		return
	default:
		// discard if channel is full
		return
	}
}
