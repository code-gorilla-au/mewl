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

// Once - Creates a function that is restricted to invoking func once.
// Repeat calls to the function return the value of the first invocation
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

func Before[T any](count int, fn ComposeFunc[T]) ComposeFunc[T] {
	invoked := 0
	var result T

	return func(t T) T {
		invoked++
		if invoked >= count {
			return result

		}
		result = fn(t)
		return result
	}
}
