package mewl

type ComposeFunc[T any] func(T) T

type PredicateFunc[T any] func(item T) bool

type MapperFunc[T any, K any] func(item T) K
