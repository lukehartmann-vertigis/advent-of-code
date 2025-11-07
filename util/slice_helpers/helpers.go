package slice_helpers

func Filter[T any](source []T, condition func(item T) bool) []T {
	filtered := []T{}
	for _, elem := range source {
		if allowed := condition(elem); allowed {
			filtered = append(filtered, elem)
		}
	}
	return filtered
}

func Map[T any, U any](source []T, mapFn func(item T) U) []U {
	mapped := make([]U, 0, len(source))
	for _, elem := range source {
		mapped = append(mapped, mapFn(elem))
	}
	return mapped
}
