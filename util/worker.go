package util

import (
	"fmt"
	"sync"
	"time"
)

type Pool struct {
	workers []*worker
	jobs    chan Job
	wg      sync.WaitGroup
}

type Job struct {
	id   int
	data interface{}
}

type worker struct {
	id     int
	jobs   chan Job
	wg     *sync.WaitGroup
	closed chan struct{}
}

func NewPool(numWorkers int) *Pool {
	p := &Pool{
		workers: make([]*worker, numWorkers),
		jobs:    make(chan Job),
	}

	for i := 0; i < numWorkers; i++ {
		w := &worker{
			id:     i,
			jobs:   p.jobs,
			wg:     &p.wg,
			closed: make(chan struct{}),
		}
		p.workers[i] = w
		go w.run()
	}
	return p
}

func (p *Pool) Submit(id int, data interface{}) {
	p.jobs <- Job{id, data}
}

func (p *Pool) Close() {
	close(p.jobs)
	p.wg.Wait()
}

func (w *worker) run() {
	for {
		select {
		case job, ok := <-w.jobs:
			if !ok {
				close(w.closed)
				return
			}
			fmt.Printf("worker %d processing job %d with data %v\n", w.id, job.id, job.data)
			w.wg.Add(1)
			time.Sleep(1 * time.Second)
			w.wg.Done()
		case <-w.closed:
			return
		}
	}
}
