package mewl

import (
	"fmt"
	"slices"
	"sort"
	"testing"

	"github.com/code-gorilla-au/odize"
)

func TestMapKeys(t *testing.T) {
	want := []string{"1", "2"}

	got := MapKeys(map[string]int{"1": 1, "2": 2})

	ForEach(want, func(item string, _ int, _ []string) {
		odize.AssertTrue(t, slices.Contains(got, item))
	})
}

func ExampleMapKeys() {
	obj := map[int]string{
		1: "1",
		2: "flash",
	}

	got := MapKeys(obj)
	sort.Ints(got)

	fmt.Println(got)
	// Output: [1 2]
}

func TestMapValues(t *testing.T) {
	want := []int{1, 2, 3}

	got := MapValues(map[string]int{"1": 1, "2": 2, "3": 3})

	ForEach(want, func(item int, _ int, _ []int) {
		odize.AssertTrue(t, slices.Contains(got, item))
	})
}

func ExampleMapValues() {
	obj := map[int]string{
		1: "hello",
		2: "world",
	}

	got := MapValues(obj)
	sort.Strings(got)

	fmt.Println(got)
	// Output: [hello world]
}

func TestMapOmitKey(t *testing.T) {
	got := MapOmitKeys(map[string]int{"1": 1, "2": 2}, "2")

	odize.AssertEqual(t, map[string]int{"1": 1}, got)
}

func ExampleMapOmitKeys() {
	obj := map[string]int{
		"hello": 1,
		"world": 2,
		"bin":   3,
	}

	got := MapOmitKeys(obj, "hello", "world")

	fmt.Println(got)
	// Output: map[bin:3]

}

func TestMapOmitKey_missing_key_returns_whole_object(t *testing.T) {
	got := MapOmitKeys(map[string]int{"1": 1, "2": 2}, "9")

	odize.AssertEqual(t, map[string]int{"1": 1, "2": 2}, got)
}

func TestMapPickKey(t *testing.T) {
	want := map[int]int{1: 1, 2: 2}

	got := MapPickKeys(map[int]int{1: 1, 2: 2, 3: 3}, 1, 2)

	odize.AssertEqual(t, want, got)
}

func ExampleMapPickKeys() {
	obj := map[int]string{
		1: "hello",
		2: "world",
		3: "bin",
	}

	got := MapPickKeys(obj, 1)

	fmt.Println(got)
	// Output: map[1:hello]
}

func TestMapPickKey_missing_keys(t *testing.T) {
	want := map[int]int{2: 2}

	got := MapPickKeys(map[int]int{1: 1, 2: 2, 3: 3}, 99, 2)

	odize.AssertEqual(t, want, got)
}

func TestMapPickKey_no_keys(t *testing.T) {
	want := map[int]int{}

	got := MapPickKeys(map[int]int{1: 1, 2: 2, 3: 3})

	odize.AssertEqual(t, want, got)
}

func TestMapOmitBy(t *testing.T) {
	want := map[int]int{1: 1}

	got := MapOmitBy(map[int]int{1: 1, 2: 2}, func(item int) bool {
		return item == 2
	})

	odize.AssertEqual(t, want, got)
}

func ExampleMapOmitBy() {
	obj := map[string]int{
		"hello": 1,
		"world": 2,
	}
	got := MapOmitBy(obj, func(item int) bool {
		return item == 1
	})

	fmt.Println(got)
	// Output: map[world:2]
}

func TestMapPickBy(t *testing.T) {
	want := map[int]int{1: 1}

	got := MapPickBy(map[int]int{1: 1, 2: 2}, func(item int) bool {
		return item == 1
	})

	odize.AssertEqual(t, want, got)
}

func ExampleMapPickBy() {
	obj := map[string]int{
		"hello": 1,
		"world": 2,
	}
	got := MapPickBy(obj, func(item int) bool {
		return item == 1
	})

	fmt.Println(got)
	// Output: map[hello:1]
}

func TestMapClone(t *testing.T) {
	group := odize.NewGroup(t, nil)

	err := group.
		Test("should clone map of strings", func(t *testing.T) {
			testVal := map[string]string{
				"hello": "world",
			}

			result := MapClone(testVal)
			odize.AssertEqual(t, testVal, result)
		}).
		Test("should clone empty map", func(t *testing.T) {
			testVal := map[string]string{}

			result := MapClone(testVal)
			odize.AssertEqual(t, testVal, result)
		}).
		Run()

	odize.AssertNoError(t, err)

}

func ExampleMapClone() {
	obj := map[string]int{
		"hello": 1,
		"world": 2,
	}
	got := MapClone(obj)

	fmt.Println(got)
	// Output: map[hello:1 world:2]
}
