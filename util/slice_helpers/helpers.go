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

func Map[T, U any](source []T, mapFn func(item T) U) []U {
	mapped := make([]U, 0, len(source))
	for _, elem := range source {
		mapped = append(mapped, mapFn(elem))
	}
	return mapped
}

func Reduce[T, U any](source []T, reduceFn func(acc U, elem T) U, initial U) U {
	result := initial
	for _, v := range source {
		result = reduceFn(result, v)
	}

	return result
}

func Some[T any](target []T, someFunc func(item T) bool) bool {
	for _, item := range target {
		if someFunc(item) {
			return true
		}
	}
	return false
}

func Count[T comparable](source []T, candidate T) (count int) {
	for _, item := range source {
		if item == candidate {
			count++
		}
	}
	return count
}
