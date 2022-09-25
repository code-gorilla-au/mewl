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

// MapOmitKeys - returns a partial copy of an object omitting the keys specified.
// If the key does not exist, the property is ignored.
func MapOmitKeys[T comparable, K any](obj map[T]K, omits ...T) map[T]K {
	for _, omit := range omits {
		delete(obj, omit)
	}
	return obj
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