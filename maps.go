package mewl

import "slices"

// MapClone clones provided map
func MapClone[T comparable, K any](obj map[T]K) map[T]K {
	result := make(map[T]K, len(obj))
	for key, value := range obj {
		result[key] = value
	}
	return result
}

// Keys - return map's keys
func MapKeys[T comparable, K any](obj map[T]K) []T {
	var result []T
	for key := range obj {
		result = append(result, key)
	}
	return result
}

// MapValues - return map's values
func MapValues[T comparable, K any](obj map[T]K) []K {
	var result []K
	for _, value := range obj {
		result = append(result, value)
	}
	return result
}

// MapOmitKeys - returns a partial copy of an object omitting the keys specified.
// If the key does not exist, the property is ignored.
func MapOmitKeys[T comparable, K any](obj map[T]K, omits ...T) map[T]K {
	result := map[T]K{}

	for key, value := range obj {
		if slices.Contains(omits, key) {
			continue
		}

		result[key] = value
	}
	return result
}

// MapOmitBy - returns a partial copy of an object omitting values based on a predicate func.
func MapOmitBy[T comparable, K any](obj map[T]K, fn PredicateFunc[K]) map[T]K {
	result := make(map[T]K, 0)

	for key, value := range obj {
		if !fn(value) {
			result[key] = value
		}
	}
	return result
}

// MapPickKeys- Returns a partial copy of an object containing only the keys specified.
// If the key does not exist, the property is ignored.
func MapPickKeys[T comparable, K any](obj map[T]K, picks ...T) map[T]K {
	result := make(map[T]K, 0)

	for _, pick := range picks {
		if value, ok := obj[pick]; ok {
			result[pick] = value
		}
	}

	return result
}

// MapPickKeys- Returns a partial copy of an object containing only the keys specified by a predicate func.
func MapPickBy[T comparable, K any](obj map[T]K, fn PredicateFunc[K]) map[T]K {
	result := make(map[T]K, 0)

	for key, value := range obj {
		if fn(value) {
			result[key] = value
		}
	}
	return result
}
