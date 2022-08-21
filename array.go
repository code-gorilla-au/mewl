package mewl

// Filter - return a new list of elements that return true on the predicate func.
func Filter[T any](list []T, fn PredicateFunc[T]) []T {
	var result []T
	for _, item := range list {
		if fn(item) {
			result = append(result, item)
		}
	}

	return result
}

// Map - creates a new array populated with the results of calling a provided function on every element in the calling array.
func Map[T comparable, K any](list []T, fn MapperFunc[T, K]) []K {
	var result []K

	for _, item := range list {
		result = append(result, fn(item))
	}
	return result
}

// ForEach - iterates over the list and invokes the function on the element.
func ForEach[T comparable](list []T, fn func(input T)) {
	for _, item := range list {
		fn(item)
	}
}

// Unique - return unique items from a provided list
func Unique[T comparable](list []T) []T {
	cache := make(map[T]struct{}, 0)
	var result []T

	for _, item := range list {
		if _, ok := cache[item]; ok {
			continue
		}

		result = append(result, item)
		cache[item] = struct{}{}
	}
	return result
}

// Union - merges two lists into a slice with no duplicates composed of the elements of each list.
func Union[T comparable](lists ...[]T) []T {
	cache := make(map[T]struct{}, 0)
	var result []T

	for _, list := range lists {
		for _, item := range list {
			if _, ok := cache[item]; ok {
				continue
			}

			result = append(result, item)
			cache[item] = struct{}{}
		}
	}

	return result
}

// Find - returns the first element in the provided array that satisfies the provided testing function.
// If item is not found return nil value.
func Find[T any](list []T, fn PredicateFunc[T]) (T, bool) {
	for _, item := range list {
		if fn(item) {
			return item, true
		}
	}

	var nilValue T
	return nilValue, false
}

// Every - tests whether all elements in the array pass the test implemented by the provided function.
func Every[T any](list []T, fn PredicateFunc[T]) bool {
	for _, item := range list {
		if !fn(item) {
			return false
		}
	}

	return true
}

// Pipe - executes a user-supplied "reducer" callback function on each element of the array,
// in order, passing in the return value from the calculation on the preceding element.
// The final result of running the reducer across all elements of the array is a single value
func Reduce[T any](list []T, fn func(prev T, next T) T) ComposeFunc[T] {
	return func(initVal T) T {
		result := initVal
		for _, item := range list {
			result = fn(result, item)
		}
		return result
	}
}

func Reverse[T comparable](list []T) []T {

	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return list
}
