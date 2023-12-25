package mewl

import (
	"testing"

	"github.com/code-gorilla-au/odize"
)

func TestPipe_int(t *testing.T) {
	add := Pipe(
		func(i int) int {
			return i + 1
		},
	)
	got := add(1)
	odize.AssertEqual(t, 2, got)

}
func TestPipe_int_product(t *testing.T) {
	add := Pipe(
		func(i int) int {
			return i + 1
		},
		func(i int) int {
			return i * 2
		},
	)
	got := add(1)
	odize.AssertEqual(t, 4, got)

}

func TestOnce(t *testing.T) {
	invoker := Once(func(i int) int {
		return i + 1
	})

	got := invoker(1)
	odize.AssertEqual(t, 2, got)
	got = invoker(2)
	odize.AssertEqual(t, 2, got)
}

func TestBefore(t *testing.T) {

	invoker := Before(2, func(args ...int) int {
		return args[0] + 1
	})

	result := invoker(1)
	odize.AssertEqual(t, 2, result)
	result = invoker(3)
	odize.AssertEqual(t, 2, result)
	result = invoker(3)
	odize.AssertEqual(t, 2, result)
}
