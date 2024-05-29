package worker

import (
	"sync"
)

// Pool is a generic pool of workers that can be used to execute tasks concurrently.
type Pool[T any] struct {
	sync.Mutex
	workers []*workerContainer[T]
}

type workerContainer[T any] struct {
	real      *T
	available bool
}

func NewPool[T any]() *Pool[T] {
	return &Pool[T]{}
}

func (p *Pool[T]) AddWorker(worker *T) {
	p.Lock()
	p.workers = append(p.workers, &workerContainer[T]{real: worker, available: true})
	p.Unlock()
}

func (p *Pool[T]) Size() int {
	p.Lock()
	defer p.Unlock()
	return len(p.workers)
}

func (p *Pool[T]) Checkout() (worker *Worker[T], ok bool) {
	p.Lock()
	defer p.Unlock()
	for i, w := range p.workers {
		if w.available {
			p.workers[i].available = false
			return &Worker[T]{container: w, expired: false}, true
		}
	}
	return nil, false
}

func (c *workerContainer[T]) Real() *T {
	return c.real
}

type Worker[T any] struct {
	container *workerContainer[T]
	expired   bool
}

func (w *Worker[T]) Return() {
	if w.expired {
		panic("worker is already returned")
	}
	w.expired = true
	w.container.available = true
}

func (w *Worker[T]) Upgrade() (real *T) {
	if w.expired {
		panic("worker is already returned")
	}
	return w.container.Real()
}
