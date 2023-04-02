package set_test

import (
	"testing"

	"github.com/playgroundgo/genlib/container/set"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func TestBasicSetOperations(t *testing.T) {
	s := set.New[int]()
	s.Add(1)
	s.Add(2)

	if !s.Contains(1) {
		t.Fatal("expected 1 to be present in the set")
	}
	if !s.ContainsAll(1, 2) {
		t.Fatal("expected 1 and 2 to be present in the set")
	}
	if s.Contains(3) {
		t.Fatal("didn't expected 3 to be present in the set")
	}
	if s.ContainsAll(1, 3) {
		t.Fatal("didn't expected 3 to be present in the set")
	}
	if !s.ContainsAny(1, 3) {
		t.Fatal("expected 1 to be present in the set")
	}

	if s.ContainsAny(4, 5) {
		t.Fatal("didn't expected 4 or 5 to be present in the set")
	}

	if s.Size() != 2 {
		t.Fatal("expected to have 2 elements in the set")
	}

	s.AddAll(1, 2, 3, 4, 5)

	if s.Size() != 5 {
		t.Fatal("expected to have 5 elements in the set")
	}

	if !s.Contains(4) {
		t.Fatal("expected 4 to be present in the set")
	}

	s.Remove(4)
	if s.Contains(4) {
		t.Fatal("didn't expected 4 to be present in the set")
	}

	s.RemoveAll(3, 5)
	if s.Size() != 2 {
		t.Fatal("expected to have 2 elements in the set")
	}

	clone := s.Clone()
	s.Clear()

	if clone.Size() != 2 {
		t.Fatal("expected to have 2 elements in the cloned set")
	}

	if !s.IsEmpty() {
		t.Fatal("expected the original set to be empty")
	}

	s.AddAll(1, 2)
	if !s.Equal(clone) {
		t.Fatal("expected the set to be equal with the clone")
	}

	s.Add(3)
	if s.Equal(clone) {
		t.Fatal("didn't expected the set to be equal with the clone")
	}

	s.Clear()
	s.AddAll(1, 3)
	if s.Equal(clone) {
		t.Fatal("didn't expected the set to be equal with the clone")
	}

	items := s.ToSlice()
	slices.Sort(items)
	expectedItems := []int{1, 3}
	if !slices.Equal(items, expectedItems) {
		t.Fatalf("expected set converted to slice to be %v, got %v", expectedItems, items)
	}
}

func TestSetCallbackIteration(t *testing.T) {
	s := set.NewWithInitialSpace[int](4)

	elems := make(map[int]struct{})
	for i := 1; i <= 4; i++ {
		elems[i] = struct{}{}
	}

	s.AddAll(maps.Keys(elems)...)
	collected := make(map[int]struct{})

	s.ForEach(func(elem int) bool {
		collected[elem] = struct{}{}
		return true
	})

	if !maps.Equal(elems, collected) {
		t.Fatalf("invalid elements added, expected %v, got %v", elems, collected)
	}

	maps.Clear(collected)

	s.ForEach(func(elem int) bool {
		if elem == 3 {
			return false
		}
		collected[elem] = struct{}{}
		return true
	})

	if _, found := collected[3]; found {
		t.Fatal("didn't expected to find 3 in the collected map")
	}
}

func TestSetIterators(t *testing.T) {
	s := set.NewWithInitialSpace[int](4)

	elems := make(map[int]struct{})
	for i := 1; i <= 100; i++ {
		elems[i] = struct{}{}
	}

	s.AddAll(maps.Keys(elems)...)
	collected := make(map[int]struct{})

	for elem := range s.Iterator().C {
		collected[elem] = struct{}{}
	}

	if !maps.Equal(elems, collected) {
		t.Fatalf("invalid elements added, expected %v, got %v", elems, collected)
	}

	maps.Clear(collected)

	for elem := range s.BufferedIterator().C {
		collected[elem] = struct{}{}
	}

	if !maps.Equal(elems, collected) {
		t.Fatalf("invalid elements added, expected %v, got %v", elems, collected)
	}

	maps.Clear(collected)
	iter := s.Iterator()
	iter.Stop()

	for elem := range iter.C {
		collected[elem] = struct{}{}
	}

	iter.Stop()

	if len(collected) > 0 {
		t.Fatalf("didn't expected to have collected elements, got %v", collected)
	}

	if !iter.Stopped() {
		t.Fatal("expected the iterator to be stopped")
	}
}

func TestIntersect(t *testing.T) {
	s1 := set.NewWithInitialSpace[int](4)
	s2 := set.NewWithInitialSpace[int](4)

	s1.AddAll(1, 2, 3, 4)
	s2.AddAll(5, 6, 7, 8)

	s := s1.Intersect(s2)
	if !s.IsEmpty() {
		t.Fatalf("expected empty interesection set, got %v", s)
	}

	s.Clear()
	s2.AddAll(2, 3)
	s = s1.Intersect(s2)
	if s.Size() != 2 {
		t.Fatalf("expected 2 elements in the intersection set, got %d", s.Size())
	}

	if !s.ContainsAll(2, 3) {
		t.Fatalf("expected 2 and 3 to be present in the interesection set %v", s)
	}
}
