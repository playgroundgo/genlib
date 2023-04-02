package generic

import "golang.org/x/exp/constraints"

// Equals verifies if the arguments are equal using the equality operator.
func Equals[T comparable](a, b T) bool {
	return a == b
}

// Less verifies if the first argument is less then the other using the "less than" comparison
// operator.
func Less[T constraints.Ordered](a, b T) bool {
	return a < b
}

// // Compare compares the two arguments.
func Compare[T constraints.Ordered](a, b T) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// CompareFunc compares the two arguments using the provided comparator function.
func CompareFunc[T any](a, b T, less LessFn[T]) int {
	if less(a, b) {
		return -1
	}
	if less(b, a) {
		return 1
	}
	return 0
}

// Min compares the given arguments and returns the smaller one.
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// MinFunc compares the given arguments using the provided comparator function and returns the
// smaller one.
func MinFunc[T any](a, b T, less LessFn[T]) T {
	if less(a, b) {
		return a
	}
	return b
}

// Max compares the given arguments and returns the largest one.
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// MaxFunc compares the given arguments using the provided comparator function and returns the
// largest one.
func MaxFunc[T any](a, b T, less LessFn[T]) T {
	if less(a, b) {
		return b
	}
	return a
}
