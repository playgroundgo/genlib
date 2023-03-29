package maps_test

import (
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/playgroundgo/genlib/functional/maps"
)

func TestHasKey(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	if !maps.ContainsKey(m, "b") {
		t.Fatal("key b should exists")
	}
	if maps.ContainsKey(m, "c") {
		t.Fatal("key c should not exists")
	}
}

func TestMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	expected := map[string]string{"A": "1", "B": "2"}
	result := maps.Map(m, func(k string, v int) (string, string) {
		return strings.ToUpper(k), strconv.Itoa(v)
	})
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected map %v, got %v", expected, result)
	}
}

func TestMapKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	expected := map[string]int{"A": 1, "B": 2}
	result := maps.MapKeys(m, func(k string) string {
		return strings.ToUpper(k)
	})
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected map %v, got %v", expected, result)
	}
}

func TestMapValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	expected := map[string]int{"a": 2, "b": 4}
	result := maps.MapValues(m, func(v int) int {
		return v * 2
	})
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected map %v, got %v", expected, result)
	}
}

func TestMerge(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"a": 11, "c": 3}
	expected := map[string]int{"a": 11, "b": 2, "c": 3}
	result := maps.Merge(m1, m2)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected map %v, got %v", expected, result)
	}
}

func TestMergeFunc(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"a": 11, "c": 3}
	expected := map[string]int{"a": 12, "b": 2, "c": 3}
	result := maps.MergeFunc(m1, m2, func(_ string, v1, v2 int) int {
		return v1 + v2
	})
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected map %v, got %v", expected, result)
	}
}

func TestWithKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	expected := map[string]int{"b": 2, "c": 3}
	result := maps.WithKeys(m, "b", "c")
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected map %v, got %v", expected, result)
	}
}

func TestWithoutKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	expected := map[string]int{"a": 1, "d": 4}
	result := maps.WithoutKeys(m, "b", "c")
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected map %v, got %v", expected, result)
	}
}
