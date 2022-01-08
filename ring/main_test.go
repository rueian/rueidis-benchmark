package ring

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

const size = 4096

type ring struct {
	_     [8]uint64
	write uint64
	_     [7]uint64
	read1 uint64
	_     [7]uint64
	mask  uint64
	_     [7]uint64
	store [size]node
}

type node struct {
	mark uint32
	val  []byte
}

type nop struct{}

func (n nop) Lock() {}

func (n nop) Unlock() {}

func (r *ring) Put(val []byte) {
	n := &r.store[atomic.AddUint64(&r.write, 1)&r.mask]
	for !atomic.CompareAndSwapUint32(&n.mark, 0, 1) {
		runtime.Gosched()
	}
	n.val = val
	atomic.StoreUint32(&n.mark, 2)
}

func (r *ring) Read() []byte {
	r.read1++
	p := r.read1 & r.mask
	n := &r.store[p]
	if atomic.CompareAndSwapUint32(&n.mark, 2, 0) {
		return n.val
	}
	r.read1--
	return nil
}

func newRing() *ring {
	r := &ring{}
	r.mask = uint64(len(r.store) - 1)
	return r
}

func BenchmarkRing(b *testing.B) {
	val := make([]byte, 0)
	b.Run("ring", func(b *testing.B) {
		s := int32(0)
		c := sync.NewCond(nop{})
		r := newRing()
		go func() {
			for atomic.LoadInt32(&s) == 0 {
				if r.Read() == nil {
					c.Wait()
				}
			}
		}()
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			r.Put(val)
			c.Broadcast()
		}
		b.StartTimer()
		atomic.StoreInt32(&s, 1)
	})
	b.Run("chan", func(b *testing.B) {
		c := make(chan []byte, size)
		go func() {
			for range c {
			}
		}()
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			c <- val
		}
		b.StopTimer()
		close(c)
	})
}
