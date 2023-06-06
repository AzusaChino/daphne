package util

import (
	"fmt"
	"testing"
)

func TestPool(t *testing.T) {
	p := NewPool(5)
	for i := 0; i < 10; i++ {
		p.Submit(i, fmt.Sprintf("data %d", i))
	}
	p.Close()
}
