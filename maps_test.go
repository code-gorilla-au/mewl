package mewl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapKeys(t *testing.T) {
	want := []string{"1", "2"}

	got := MapKeys(map[string]int{"1": 1, "2": 2})

	ForEach(want, func(item string) {
		assert.Contains(t, got, item)
	})
	assert.Equal(t, want, got)
}

func TestMapValues(t *testing.T) {
	want := []int{1, 2, 3}

	got := MapValues(map[string]int{"1": 1, "2": 2, "3": 3})

	ForEach(want, func(item int) {
		assert.Contains(t, got, item)
	})
}
