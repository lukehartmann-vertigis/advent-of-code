package collections

// Pair represents a generic key/value tuple.
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// Map is a generic wrapper around the builtin `map` type that additionally provides helpers.
type Map[K comparable, V any] map[K]V

// NewMap creates a new map from key/value pairs.
func NewMap[K comparable, V any](pairs ...Pair[K, V]) Map[K, V] {
	m := make(Map[K, V], len(pairs))
	for _, pair := range pairs {
		m[pair.Key] = pair.Value
	}
	return m
}

// Has shows wether an entry exists for the given Key.
func (m Map[K, V]) Has(key K) bool {
	_, exists := m[key]
	return exists
}

// Get fetches an element for the given key if it exists.
func (m Map[K, V]) Get(key K) (V, bool) {
	v, exists := m[key]
	return v, exists
}

// Set sets an element for a given key.
func (m Map[K, V]) Set(key K, value V) {
	m[key] = value
}

// Delete deletes the entry for the given key.
func (m Map[K, V]) Delete(key K) bool {
	_, existed := m[key]
	delete(m, key)
	return existed
}

// Keys returns all Map[K, V] keys as a slice.
func (m Map[K, V]) Keys() []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values returns all Map[K, V] values as a slice.
func (m Map[K, V]) Values() []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
