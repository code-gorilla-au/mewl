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

// Creates a function that invokes func while it's called less than n times.
// Subsequent calls to the created function return the result of the last func invocation.
func Before[T any](n int, fn AnyFunc[T]) AnyFunc[T] {
	invoked := 0
	var result T

	return func(args ...T) T {
		invoked++
		if invoked >= n {
			return result

		}
		result = fn(args...)
		return result
	}
}
