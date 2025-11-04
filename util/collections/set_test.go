package collections

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	s := NewSet[int]()
	if s == nil {
		t.Fatal("expected NewSet() to return a non-nil set")
	}
	if s.Len() != 0 {
		t.Fatalf("expected empty set, got Len=%d", s.Len())
	}
}

func TestAddHasAndRemove(t *testing.T) {
	items := []int{1, 2, 3}
	s := NewSet[int]()
	s.Add(items...)

	for _, value := range items {
		if !s.Has(value) {
			t.Errorf("expected Set to have %d", value)
		}

		s.Remove(value)

		if s.Has(value) {
			t.Errorf("expected %d to be removed from the set", value)
		}
	}

}

func TestLen(t *testing.T) {
	items := []int{1, 2, 3}
	s := NewSet[int]()
	s.Add(items...)
	compareLen := len(items)

	if s.Len() != compareLen {
		t.Errorf("expected Set to have Len=%d, got Len=%d", compareLen, s.Len())
	}
}

func TestClear(t *testing.T) {
	items := []int{1, 2, 3}
	s := NewSet[int]()
	s.Add(items...)
	s.Clear()

	if s.Len() > 0 {
		t.Errorf("expected Set to be empty, got Len=%d", s.Len())
	}
}

func TestToSlice(t *testing.T) {
	items := []int{1, 2, 3}
	s := NewSet[int]()
	s.Add(items...)
	sl := s.ToSlice()

	if s.Len() != len(sl) {
		t.Errorf("expected the Set and Slice to be of the same size, Slice: Len=%d, Set: Len=%d", len(sl), s.Len())
	}

	for _, val := range sl {
		if !s.Has(val) {
			t.Errorf("expected %d to be in the Set", val)
		}
	}
}
