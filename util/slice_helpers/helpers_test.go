package slice_helpers

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	odds := Filter(values, func(item int) bool {
		return item%2 != 0
	})

	if len(odds) != 3 {
		t.Errorf("expected filtered slice to have length of 3, instead got %d", len(odds))
	}

	oddsCompare := []int{1, 3, 5}
	if !reflect.DeepEqual(odds, oddsCompare) {
		t.Errorf("expected %v, instead got %v", oddsCompare, odds)
	}

	valuesCompare := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(values, valuesCompare) {
		t.Errorf("expected %v, instead got %v", values, valuesCompare)
	}
}

func TestMap(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	doubled := Map(values, func(item int) int {
		return 2 * item
	})

	if len(values) != len(doubled) {
		t.Errorf("expected slices to have the same length of %d, instead got %d", len(values), len(doubled))
	}

	doubledCompare := []int{2, 4, 6, 8, 10}
	if !reflect.DeepEqual(doubled, doubledCompare) {
		t.Errorf("expected slice to match compare value %v, instead got %v", doubledCompare, doubled)
	}
}

func TestReduce(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	prod := Reduce(values, func(acc int, elem int) int {
		return acc * elem
	}, 1)

	if prod != 120 {
		t.Errorf("expected product to be %d, instead got %d", 120, prod)
	}
}
