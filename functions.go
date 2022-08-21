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
