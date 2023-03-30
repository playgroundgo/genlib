package slices

import (
	"github.com/playgroundgo/genlib/errors"
	"golang.org/x/exp/constraints"
)

// Any returns true if the function 'f' evaluates to true for any element of the slice 'items'
func Any[S ~[]T, T any](items S, f func(item T) bool) bool {
	for _, item := range items {
		if f(item) {
			return true
		}
	}
	return false
}

// All returns true if the function 'f' evaluates to true for all elements in the slice 'items'.
func All[S ~[]T, T any](items S, f func(item T) bool) bool {
	for _, item := range items {
		if !f(item) {
			return false
		}
	}
	return true
}

// Count returns the number of occurrences for an element in a slice.
func Count[S ~[]T, T comparable](items S, item T) int {
	count := 0
	for _, val := range items {
		if val == item {
			count++
		}
	}
	return count
}

// CountFunc returns the number of elements in slice 'items' for which the function 'f' evaluates
// to true.
func CountFunc[S ~[]T, T any](items S, f func(item T) bool) int {
	count := 0
	for _, item := range items {
		if f(item) {
			count++
		}
	}
	return count
}

// Filter returns a new slice with the elements from the 'items' slice for which the function 'f'
// returned true.
func Filter[S ~[]T, T any](items S, f func(item T) bool) S {
	result := make([]T, 0, len(items))
	for _, item := range items {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}

// ForEach calls the function 'f' for each element in the slice 'items'.
func ForEach[S ~[]T, T any](items S, f func(item T)) {
	for _, item := range items {
		f(item)
	}
}

// Map calls the function 'f' for each element of the 'items' slice and returns a new slice with
// the results.
func Map[S ~[]T, T any, U any](items S, f func(item T) U) []U {
	result := make([]U, len(items))
	for i, item := range items {
		result[i] = f(item)
	}
	return result
}

// Max returns the maximum element from the given slice.
func Max[S ~[]T, T constraints.Ordered](items S) (T, error) {
	if len(items) == 0 {
		var tmp T
		return tmp, errors.ErrEmpty
	}
	max := items[0]
	for _, item := range items[1:] {
		if item > max {
			max = item
		}
	}
	return max, nil
}

// Min returns the minimum element from the given slice.
func Min[S ~[]T, T constraints.Ordered](items S) (T, error) {
	if len(items) == 0 {
		var tmp T
		return tmp, errors.ErrEmpty
	}
	min := items[0]
	for _, item := range items[1:] {
		if item < min {
			min = item
		}
	}
	return min, nil
}

// Reduce applies the function 'f' to every element in the slice 'items' and to accumulator 'acc'
// and returns the accumulator.
func Reduce[S ~[]T, T any, U any](items S, acc U, f func(item T, acc U) U) U {
	for _, item := range items {
		acc = f(item, acc)
	}
	return acc
}
