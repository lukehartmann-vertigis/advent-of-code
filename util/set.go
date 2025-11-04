package util

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() *Set[T] {
	s := Set[T]{}
	return &s
}

func (s *Set[T]) Add(elements ...T) {
	for _, element := range elements {
		(*s)[element] = struct{}{}
	}
}

func (s *Set[T]) Remove(element T) {
	delete((*s), element)
}

func (s *Set[T]) Has(element T) bool {
	_, exists := (*s)[element]
	return exists
}

func (s *Set[T]) Len() int {
	return len(*s)
}

func (s *Set[T]) ToSlice() []T {
	elements := make([]T, 0, len(*s))
	for e := range *s {
		elements = append(elements, e)
	}
	return elements
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	union := make(Set[T], len(*s)+len(*other))

	for v := range *s {
		union[v] = struct{}{}
	}
	for v := range *other {
		union[v] = struct{}{}
	}

	return &union
}

func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	intersect := NewSet[T]()

	for v := range *s {
		if _, ok := (*other)[v]; ok {
			(*intersect)[v] = struct{}{}
		}
	}

	return intersect
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	difference := make(Set[T], len(*s))

	for v := range *s {
		if _, ok := (*other)[v]; !ok {
			difference[v] = struct{}{}
		}
	}

	return &difference
}
