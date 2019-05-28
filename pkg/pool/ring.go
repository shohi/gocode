package pool

// Ring is growing container, when item exceeds the
// size, new backend array will be doubled to hold new data.
type Ring struct {
	buf     []interface{}
	readAt  int
	writeAt int

	readable int // how many items are avaliable in the Ring
	cap      int // capacity
	growable bool
}

func NewRing(cap int, growable bool) *Ring {
	return &Ring{
		buf:      make([]interface{}, cap),
		readAt:   0,
		writeAt:  0,
		readable: 0,
		cap:      cap,
		growable: growable,
	}
}

// Read will take an item from ring if available
// and mark the space unoccupied.
func (r *Ring) Read() (v interface{}, ok bool) {
	if r.readable == 0 {
		return nil, false
	}
	v = r.buf[r.readAt]
	r.buf[r.readAt] = nil // mark the space unoccupied
	r.readable--
	r.readAt = (r.readAt + 1) % r.cap

	return v, true
}

// Read will take an item from ring if available
// and mark the space unoccupied.
func (r *Ring) Write(v interface{}) {
	if r.readable == r.cap && r.growable {
		r.grow(r.cap)
	}

	r.buf[r.writeAt] = v
	r.writeAt = (r.writeAt + 1) % r.cap
	r.incrReadable()
}

// incrReadable increase `readable` count if ring is not full
func (r *Ring) incrReadable() {
	if r.readable == r.cap && !r.growable {
		return
	}

	r.readable++
}

// grow grows the ring to guarantee space for n more items.
// It will update `readAt/writeAt` indices.
func (r *Ring) grow(n int) {
	newCap := r.cap + n
	newItems := make([]interface{}, newCap)

	r.copyTo(newItems)
	r.cap = newCap

	r.buf = newItems
	r.readAt = 0
	r.writeAt = r.readable
}

func (r *Ring) copyTo(data []interface{}) {
	if r.readable == 0 {
		return
	}

	if r.readAt < r.writeAt {
		copy(data, r.buf[r.readAt:r.writeAt])
	} else {
		copy(data, r.buf[r.readAt:])
		copy(data[(len(r.buf)-r.readAt):], r.buf[:r.writeAt])
	}
}
