package util

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	s := Set[T]{}
	return s
}

func (s Set[T]) Add(elements ...T) {
	for _, element := range elements {
		s[element] = struct{}{}
	}
}

func (s Set[T]) Remove(element T) {
	delete(s, element)
}

func (s Set[T]) Has(element T) bool {
	_, exists := s[element]
	return exists
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) ToSlice() []T {
	elements := make([]T, 0, len(s))
	for e := range s {
		elements = append(elements, e)
	}
	return elements
}
