package mewl

// Keys - return map's keys
func MapKeys[T comparable, K any](obj map[T]K) []T {
	var result []T
	for key, _ := range obj {
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
