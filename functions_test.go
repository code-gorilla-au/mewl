package mewl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipe_int(t *testing.T) {
	add := Pipe(
		func(i int) int {
			return i + 1
		},
	)
	got := add(1)
	assert.Equal(t, 2, got)

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
	assert.Equal(t, 4, got)

}

func TestOnce(t *testing.T) {
	invoker := Once(func(i int) int {
		return i + 1
	})

	got := invoker(1)
	assert.Equal(t, 2, got)
	got = invoker(2)
	assert.Equal(t, 2, got)
}
