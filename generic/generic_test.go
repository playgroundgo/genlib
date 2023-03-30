package generic_test

import (
	"testing"

	"github.com/playgroundgo/genlib/generic"
)

func TestLess(t *testing.T) {
	a, b := 1, 2
	if !generic.Less(a, b) {
		t.Fatalf("expected that %d is less than %d", a, b)
	}
	if generic.Less(b, a) {
		t.Fatalf("expected that %d is less than %d", a, b)
	}
}

func TestCompare(t *testing.T) {
	a, b := 1, 2
	if generic.Compare(a, b) != -1 {
		t.Fatalf("expected that %d is less than %d", a, b)
	}
	if generic.Compare(b, a) != 1 {
		t.Fatalf("expected that %d is greater than %d", b, a)
	}
	b = 1
	if generic.Compare(a, b) != 0 {
		t.Fatalf("expected that %d is equal to %d", a, b)
	}
}
