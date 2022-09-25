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

func TestMapOmitKey(t *testing.T) {
	got := MapOmitKeys(map[string]int{"1": 1, "2": 2}, "2")

	assert.Equal(t, map[string]int{"1": 1}, got)
}

func TestMapOmitKey_missing_key_returns_whole_object(t *testing.T) {
	got := MapOmitKeys(map[string]int{"1": 1, "2": 2}, "9")

	assert.Equal(t, map[string]int{"1": 1, "2": 2}, got)
}

func TestMapPickKey(t *testing.T) {
	want := map[int]int{1: 1, 2: 2}

	got := MapPickKeys(map[int]int{1: 1, 2: 2, 3: 3}, 1, 2)

	assert.Equal(t, want, got)
}

func TestMapPickKey_missing_keys(t *testing.T) {
	want := map[int]int{2: 2}

	got := MapPickKeys(map[int]int{1: 1, 2: 2, 3: 3}, 99, 2)

	assert.Equal(t, want, got)
}

func TestMapPickKey_no_keys(t *testing.T) {
	want := map[int]int{}

	got := MapPickKeys(map[int]int{1: 1, 2: 2, 3: 3})

	assert.Equal(t, want, got)
}
