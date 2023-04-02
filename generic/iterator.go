package generic

import "sync/atomic"

// Iterator models an iterator over a collection.
type Iterator[T any] struct {
	C       <-chan T
	StopCh  chan struct{}
	stopped atomic.Bool
}

// NewBufferedIterator creates a new iterator backed by a buffered channel with the given capacity.
// The client code is responsible for inserting elements in the iterator channel.
func NewBufferedIterator[T any](bufSize int) (*Iterator[T], chan<- T) {
	iterCh := make(chan T, bufSize)
	return &Iterator[T]{
		C:      iterCh,
		StopCh: make(chan struct{}),
	}, iterCh
}

// NewIterator creates a new iterator backed by an unbuffered channel.
func NewIterator[T any]() (*Iterator[T], chan<- T) {
	return NewBufferedIterator[T](0)
}

// Stop will signal the client code to stop inserting elements in the iterator channel and will
// consume all the elements remaining in the channel.
// Trying to stop the iterators multiple times is safe (all subsequent operations are no-ops).
func (it *Iterator[T]) Stop() {
	stopped := it.stopped.Swap(true)
	if stopped {
		return
	}

	close(it.StopCh)

	for range it.C {
	}
}

// Stopped returns true if the iterator was stopped.
func (it *Iterator[T]) Stopped() bool {
	return it.stopped.Load()
}
