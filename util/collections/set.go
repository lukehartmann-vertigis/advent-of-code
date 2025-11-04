package collections

// Set is a generic set implemented as a map[T]struct{}.
type Set[T comparable] map[T]struct{}

// NewSet returns a new empty Set[T].
func NewSet[T comparable]() Set[T] {
	s := Set[T]{}
	return s
}

// Add inserts one or more elements into the set.
func (s Set[T]) Add(elements ...T) {
	for _, element := range elements {
		s[element] = struct{}{}
	}
}

// Remove deletes an element from the set. It is a no-op if the element is absent.
func (s Set[T]) Remove(element T) {
	delete(s, element)
}

// Has reports whether the element is a member of the set.
func (s Set[T]) Has(element T) bool {
	_, exists := s[element]
	return exists
}

// Len returns the number of elements in the set.
func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Clear() {
	for v := range s {
		delete(s, v)
	}
}

// ToSlice returns a slice containing the set's elements. The order is unspecified.
func (s Set[T]) ToSlice() []T {
	elements := make([]T, 0, len(s))
	for e := range s {
		elements = append(elements, e)
	}
	return elements
}

// Union returns a new set containing all elements that are in s or other.
func (s Set[T]) Union(other Set[T]) Set[T] {
	union := make(Set[T], len(s)+len(other))
	for v := range s {
		union[v] = struct{}{}
	}
	for v := range other {
		union[v] = struct{}{}
	}
	return union
}

// Intersect returns a new set containing elements present in both s and other.
func (s Set[T]) Intersect(other Set[T]) Set[T] {
	intersect := NewSet[T]()
	for v := range s {
		if _, ok := other[v]; ok {
			intersect[v] = struct{}{}
		}
	}
	return intersect
}

// Difference returns a new set containing elements present in s but not in other.
func (s Set[T]) Difference(other Set[T]) Set[T] {
	difference := make(Set[T], len(s))
	for v := range s {
		if _, ok := other[v]; !ok {
			difference[v] = struct{}{}
		}
	}
	return difference
}
