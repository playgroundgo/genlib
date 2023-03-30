package slices_test

import (
	"errors"
	"reflect"
	"testing"

	gerrors "github.com/playgroundgo/genlib/errors"
	"github.com/playgroundgo/genlib/generic/slices"
)

func TestAny(t *testing.T) {
	positives := []int{1, 2, 3, 4, 2, 3, 2, 5, 6, 7}
	posFunc := func(item int) bool {
		return item >= 0
	}
	negFunc := func(item int) bool {
		return item < 0
	}
	pos := slices.Any(positives, posFunc)
	if !pos {
		t.Fatal("expected positives numbers")
	}
	neg := slices.Any(positives, negFunc)
	if neg {
		t.Fatal("didn't expected negative numbers")
	}
}

func TestAll(t *testing.T) {
	items := []int{1, 2, 3, 4, 2, 3, 2, 5, 6, 7}
	posFunc := func(item int) bool {
		return item >= 0
	}
	evenFunc := func(item int) bool {
		return item%2 == 0
	}
	pos := slices.All(items, posFunc)
	if !pos {
		t.Fatal("expected positives numbers")
	}
	even := slices.All(items, evenFunc)
	if even {
		t.Fatal("didn't expected all numbers to be even")
	}
}

func TestCount(t *testing.T) {
	items := []int{1, 2, 3, 4, 2, 3, 2, 5, 6, 7}
	count := slices.Count(items, 2)
	if count != 3 {
		t.Fatalf("expected number 2 to appear 3 times, not %d times", count)
	}
}

func TestCountFunc(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 11}
	count := slices.CountFunc(items, func(item int) bool {
		return item%2 == 0
	})
	if count != 5 {
		t.Fatalf("expected 5 even numbers, got %d", count)
	}
}

func TestFilter(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7}
	expectedItems := []int{2, 4, 6}
	evenItems := slices.Filter(items, func(item int) bool {
		return item%2 == 0
	})
	if !reflect.DeepEqual(expectedItems, evenItems) {
		t.Fatalf("expected slice %v, got %v", expectedItems, evenItems)
	}
}

func TestForEach(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7}
	result := make([]int, 0, 3)
	expectedItems := []int{2, 4, 6}
	slices.ForEach(items, func(item int) {
		if item%2 == 0 {
			result = append(result, item)
		}
	})
	if !reflect.DeepEqual(expectedItems, result) {
		t.Fatalf("expected slice %v, got %v", expectedItems, result)
	}
}

func TestMap(t *testing.T) {
	items := []int{1, 3, 5, 7}
	expectedItems := []int{2, 6, 10, 14}
	result := slices.Map(items, func(item int) int {
		return item * 2
	})
	if !reflect.DeepEqual(expectedItems, result) {
		t.Fatalf("expected slice %v, got %v", expectedItems, result)
	}
}

func TestMax(t *testing.T) {
	items := []int{10, 22, 14, 3, 99, 8, 9, 6}
	max, _ := slices.Max(items)
	if max != 99 {
		t.Fatalf("expected max to be 99, got %d", max)
	}
	_, err := slices.Max([]int{})
	if err == nil || !errors.Is(err, gerrors.ErrEmpty) {
		t.Fatal("expected empty container error")
	}
}

func TestMin(t *testing.T) {
	items := []int{10, 22, 14, 3, 99, 8, 9, 6}
	min, _ := slices.Min(items)
	if min != 3 {
		t.Fatalf("expected min to be 3, got %d", min)
	}
	_, err := slices.Min([]int{})
	if err == nil || !errors.Is(err, gerrors.ErrEmpty) {
		t.Fatal("expected empty container error")
	}
}

func TestReduce(t *testing.T) {
	items := []int{1, 2, 3, 4}
	result := slices.Reduce(items, 0, func(item, acc int) int {
		return item + acc
	})
	if result != 10 {
		t.Fatalf("expected result to be 10, got %d", result)
	}
}
