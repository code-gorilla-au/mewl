package mewl

// Pipe - left to right function composition
func Pipe[T any](fns ...ComposeFunc[T]) ComposeFunc[T] {
	return func(input T) T {
		result := input
		for _, fn := range fns {
			result = fn(result)
		}
		return result
	}
}

func Once[T any](fn ComposeFunc[T]) ComposeFunc[T] {
	invoked := false

	var result T

	return func(t T) T {
		if !invoked {
			invoked = true
			result = fn(t)
			return result
		}
		return result
	}
}
