package util

import (
	"fmt"
	"testing"
)

func TestPool(t *testing.T) {
	p := NewPool(5, func(job Job) {
		fmt.Printf("worker %d: %v\n", job.id, job.data)
	})
	for i := 0; i < 10; i++ {
		p.Submit(i, fmt.Sprintf("data %d", i))
	}
	p.Close()
}
