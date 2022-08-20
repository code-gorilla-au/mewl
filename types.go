package mewl

// ComposeFunc - function that receives an input and returns an input of the same type.
type ComposeFunc[T any] func(T) T

// PredicateFunc - function that receives an input and returns a boolean value.
type PredicateFunc[T any] func(item T) bool

// MapperFunc - transform function that receives an input and returns a new type.
type MapperFunc[T any, K any] func(item T) K
