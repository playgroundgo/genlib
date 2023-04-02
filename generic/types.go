package generic

// EqualsFn is a function type that checks if two arguments are equal.
type EqualsFn[T any] func(a, b T) bool

// LessFn is a function type that checks if the first argument is less than the second one.
type LessFn[T any] func(a, b T) bool

// Pair is a generic struct holding two generic elements.
type Pair[T, U any] struct {
	First  T
	Second U
}
