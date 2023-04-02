package set

import (
	"github.com/playgroundgo/genlib/generic"
)

// Set implements a set container.
type Set[T comparable] map[T]struct{}

// New creates a new set.
func New[T comparable]() Set[T] {
	return make(Set[T])
}

// NewWithInitialSpace returns a new set having enough space to hold 'n' elements.
func NewWithInitialSpace[T comparable](n int) Set[T] {
	return make(Set[T], n)
}

// Size returns the number of elements in the set.
func (s Set[T]) Size() int {
	return len(s)
}

// IsEmpty returns 'true' if the set is empty.
func (s Set[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Equal checks if this set is equal with another one.
func (s Set[T]) Equal(other Set[T]) bool {
	if s.Size() != other.Size() {
		return false
	}
	for elem := range s {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// Clear clears the content of the set.
func (s Set[T]) Clear() {
	for elem := range s {
		delete(s, elem)
	}
}

// Clone returns a copy of the given set.
func (s Set[T]) Clone() Set[T] {
	cloned := NewWithInitialSpace[T](s.Size())
	for elem := range s {
		cloned.Add(elem)
	}
	return cloned
}

// Contains verifies if an element belongs to the set.
func (s Set[T]) Contains(elem T) bool {
	_, found := s[elem]
	return found
}

// ContainsAll verifies if all the specified elements are found in the set.
func (s Set[T]) ContainsAll(elems ...T) bool {
	for _, elem := range elems {
		if !s.Contains(elem) {
			return false
		}
	}
	return true
}

// ContainsAny verifies if any element is found in the set.
func (s Set[T]) ContainsAny(elems ...T) bool {
	for _, elem := range elems {
		if s.Contains(elem) {
			return true
		}
	}
	return false
}

// ForEach calls the 'f' function for each element in the set while the function returns true.
func (s Set[T]) ForEach(f func(elem T) bool) {
	for elem := range s {
		if !f(elem) {
			break
		}
	}
}

// Iterator returns an iterator over the set.
func (s Set[T]) Iterator() *generic.Iterator[T] {
	return s.iterate(generic.NewIterator[T]())
}

// Iterator returns a buffered iterator over the set.
func (s Set[T]) BufferedIterator() *generic.Iterator[T] {
	return s.iterate(generic.NewBufferedIterator[T](len(s)))
}

// Add adds an element to the set.
func (s Set[T]) Add(elem T) {
	s[elem] = struct{}{}
}

// AddAll adds all the specified elements in the set.
func (s Set[T]) AddAll(elems ...T) {
	for _, elem := range elems {
		s.Add(elem)
	}
}

// Remove removes an element from the set.
func (s Set[T]) Remove(elem T) {
	delete(s, elem)
}

// RemoveAll removes all the specified elements from the set.
func (s Set[T]) RemoveAll(elems ...T) {
	for _, elem := range elems {
		s.Remove(elem)
	}
}

// Intersect returns the intersection between this set and another one.
func (s Set[T]) Intersect(other Set[T]) Set[T] {
	if s.Size() < other.Size() {
		return intersect(s, other)
	}
	return intersect(other, s)
}

// Intersect returns the union between this set and another one.
func (s Set[T]) Union(other Set[T]) Set[T] {
	result := NewWithInitialSpace[T](generic.Max(s.Size(), other.Size()))
	for elem := range s {
		result.Add(elem)
	}
	for elem := range other {
		result.Add(elem)
	}
	return result
}

// ToSlice returns a slice containing all the elements from the set.
func (s Set[T]) ToSlice() []T {
	elems := make([]T, 0, len(s))
	for elem := range s {
		elems = append(elems, elem)
	}
	return elems
}

func (s Set[T]) iterate(it *generic.Iterator[T], itCh chan<- T) *generic.Iterator[T] {
	go func() {
	L:
		for elem := range s {
			select {
			case <-it.StopCh:
				break L
			case itCh <- elem:
			}
		}
		close(itCh)
	}()
	return it
}

func intersect[T comparable](s1 Set[T], s2 Set[T]) Set[T] {
	result := New[T]()
	for elem := range s1 {
		if s2.Contains(elem) {
			result.Add(elem)
		}
	}
	return result
}
