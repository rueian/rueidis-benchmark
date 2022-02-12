package ring

import (
	"runtime"
	"sync/atomic"
	"testing"
)

const size = 1024
const mask = size - 1

type request struct {
	ch chan response
}
type response struct{}

type PipeliningQueue interface {
	EnqueueRequest(req request) chan response
	NextRequestToSend() (req request, ok bool)
	ReplyToNextRequest(response)
}

type ring struct {
	_     [8]uint64
	write uint64
	_     [7]uint64
	read1 uint64
	_     [7]uint64
	read2 uint64
	_     [7]uint64
	slots [size]slot
}

type slot struct {
	mark uint32
	req  request
	ch   chan response
}

func (r *ring) EnqueueRequest(req request) chan response {
	s := &r.slots[atomic.AddUint64(&r.write, 1)&mask]
	for atomic.CompareAndSwapUint32(&s.mark, 0, 1) {
		runtime.Gosched()
	}
	s.req = req
	atomic.StoreUint32(&s.mark, 2)
	return s.ch
}

func (r *ring) NextRequestToSend() (req request, ok bool) {
	s := &r.slots[(r.read1+1)&mask]
	if ok = atomic.LoadUint32(&s.mark) == 2; ok {
		req = s.req
		r.read1++
		atomic.StoreUint32(&s.mark, 3)
	}
	return
}

func (r *ring) ReplyToNextRequest(resp response) {
	r.read2++
	s := &r.slots[r.read2&mask]
	if atomic.LoadUint32(&s.mark) != 3 {
		r.read2--
	} else {
		atomic.StoreUint32(&s.mark, 0)
	}
}

func newRing() *ring {
	r := &ring{}
	for i := range r.slots {
		r.slots[i] = slot{ch: make(chan response, 0)}
	}
	return r
}

type double struct {
	writing chan request
	waiting chan request
}

func (d *double) EnqueueRequest(req request) chan response {
	req.ch = make(chan response, 1)
	d.writing <- req
	return req.ch
}

func (d *double) NextRequestToSend() (req request, ok bool) {
	select {
	case req, ok = <-d.writing:
	default:
		return request{}, false
	}
	if ok {
		d.waiting <- req
	}
	return
}

func (d *double) ReplyToNextRequest(r response) {
	<-d.waiting
}

func newDouble() *double {
	return &double{
		writing: make(chan request, size/2),
		waiting: make(chan request, size/2),
	}
}

func BenchmarkPipeliningQueue(b *testing.B) {
	bench := func(queue PipeliningQueue) (func(), func()) {
		stop := int32(0)
		go func() {
			for atomic.LoadInt32(&stop) == 0 {
				if _, ok := queue.NextRequestToSend(); ok {
					queue.ReplyToNextRequest(response{})
				} else {
					runtime.Gosched()
				}
			}
		}()
		return func() {
				_ = queue.EnqueueRequest(request{})
			}, func() {
				if d, ok := queue.(*double); ok {
					close(d.writing)
				}
				atomic.StoreInt32(&stop, 1)
			}
	}
	b.Run("Lockless", func(b *testing.B) {
		fn, cancel := bench(newRing())
		b.SetParallelism(64)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				fn()
			}
		})
		b.StopTimer()
		cancel()
	})
	b.Run("Channel", func(b *testing.B) {
		fn, cancel := bench(newDouble())
		b.SetParallelism(64)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				fn()
			}
		})
		b.StopTimer()
		cancel()
	})
}
