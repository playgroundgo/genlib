package xsync

import (
	"sync"
)

// Pool is a generic and safer version of the standard library sync.Pool.
type Pool[T any] struct {
	syncPool sync.Pool
}

// NewPool creates a new pool.
func NewPool[T any](newFn func() T) *Pool[T] {
	return &Pool[T]{
		syncPool: sync.Pool{
			New: func() any { return newFn() },
		},
	}
}

// Get returns an item from the pool.
func (p *Pool[T]) Get() T {
	return p.syncPool.Get().(T)
}

// Put adds an item to the pool.
func (p *Pool[T]) Put(value T) {
	p.syncPool.Put(value)
}
