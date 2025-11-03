package util

type Map[K comparable, V any] map[K]V

func NewMap[K comparable, V any](pairs ...struct {
	Key   K
	Value V
}) Map[K, V] {
	m := Map[K, V]{}
	for _, pair := range pairs {
		m[pair.Key] = pair.Value
	}
	return m
}
