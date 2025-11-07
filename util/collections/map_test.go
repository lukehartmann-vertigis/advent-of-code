package collections

import (
	"testing"
)

func TestNewMap(t *testing.T) {
	pairs := make([]Pair[string, []int], 3)

	m := NewMap(pairs...)

	if m == nil {
		t.Fatal("expected Map to be not nil.")
	}
}

func TestSetHasAndDelete(t *testing.T) {
	pairs := make([]Pair[string, []int], 1)
	pairs[0] = Pair[string, []int]{
		Key:   "foo",
		Value: []int{1, 2, 3},
	}
	m := NewMap(pairs...)
	m.Set("test", []int{1, 2, 3})

	if !m.Has("test") {
		t.Error("expected Map to have an element for the Key 'test'.")
	}

	m.Delete("test")
	if m.Has("test") {
		t.Error("expected Map not to have an element for the Key 'test'.")
	}
}

func TestKeys(t *testing.T) {
	keysDict := map[string]struct{}{
		"foo1": {},
		"foo2": {},
		"foo3": {},
	}
	pairs := make([]Pair[string, int], len(keysDict))

	idx := 0
	for key := range keysDict {
		pairs[idx] = Pair[string, int]{
			Key:   key,
			Value: idx,
		}
		idx++
	}

	m := NewMap(pairs...)
	mapKeys := m.Keys()

	if len(mapKeys) != len(keysDict) {
		t.Errorf("expected map keys to have the same amount of elements as the source keyDict: %d", len(keysDict))
	}

	for _, key := range mapKeys {
		_, exists := keysDict[key]
		if !exists {
			t.Errorf("Expected Key %s to be in the source keysDict", key)
		}
	}
}

func TestValues(t *testing.T) {
	valuesDict := map[int]struct{}{
		1: {},
		2: {},
		3: {},
	}

	pairs := make([]Pair[string, int], len(valuesDict))

	idx := 0
	for val := range valuesDict {
		pairs[idx] = Pair[string, int]{
			Key:   "key" + string(rune(idx)),
			Value: val,
		}
		idx++
	}

	m := NewMap(pairs...)
	mapValues := m.Values()

	if len(mapValues) != len(valuesDict) {
		t.Errorf("expected values slice to have the same amount of elements as the source valuesDict: %d", len(valuesDict))
	}

	for _, val := range mapValues {
		if _, exists := valuesDict[val]; !exists {
			t.Errorf("Expected Value %d to be in the source valuesDict", val)
		}
	}
}
