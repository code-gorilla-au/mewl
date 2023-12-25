package mewl

// ComposeFunc - function that receives an input and returns an input of the same type.
type ComposeFunc[T any] func(T) T

// PredicateFunc - function that receives an input and returns a boolean value.
type PredicateFunc[T any] func(item T) bool

// PredicateSliceFunc - function that receives an index of the current item, current item being iterated on and the array Every was called upon.
type PredicateSliceFunc[T any] func(item T, index int, slice []T) bool

// MapperFunc - transform function that receives an input and returns a new type.
type MapperFunc[T any, K any] func(item T) K

// AnyFunc - function that receives n arguments
type AnyFunc[T any] func(args ...T) T
