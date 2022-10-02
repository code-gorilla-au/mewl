package mewl

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type CmpUUID struct {
	id   uuid.UUID
	date time.Time
}

type KeyVal struct {
	Key   string
	Value string
}

type Val struct {
	Value string
}

func TestFilter_struct(t *testing.T) {
	list := []KeyVal{
		{
			Key:   "foo",
			Value: "bar",
		},
		{
			Key:   "bin",
			Value: "baz",
		},
	}

	got := Filter(list, func(item KeyVal) bool {
		return item.Key == "foo"
	})
	if len(got) != 1 {
		t.Fatal("expected len of 1, got: ", len(got))
	}
	if got[0].Key != "foo" {
		t.Fatalf("struct not matching")
	}
}

func ExampleFilter() {
	list := []KeyVal{
		{
			Key:   "foo",
			Value: "bar",
		},
		{
			Key:   "bin",
			Value: "baz",
		},
	}

	got := Filter(list, func(item KeyVal) bool {
		return item.Key == "foo"
	})
	fmt.Println(got)
	// Output: [{foo bar}]
}

func TestMap_struct(t *testing.T) {
	list := []KeyVal{
		{
			Key:   "foo",
			Value: "bar",
		},
		{
			Key:   "bin",
			Value: "baz",
		},
	}

	got := Map(list, func(item KeyVal) Val {
		return Val{
			Value: item.Value,
		}
	})
	for index, item := range got {
		if item.Value != list[index].Value {
			t.Fatal("non matching value")
		}
	}
}

func ExampleMap() {
	list := []KeyVal{
		{
			Key:   "foo",
			Value: "bar",
		},
		{
			Key:   "bin",
			Value: "baz",
		},
	}

	got := Map(list, func(item KeyVal) Val {
		return Val{
			Value: item.Value,
		}
	})

	fmt.Println(got)
	// Output: [{bar} {baz}]
}

func TestForEach_struct(t *testing.T) {
	list := []int{1, 1, 2}

	total := 0

	ForEach(list, func(item int) {
		total += item
	})
	if total != 4 {
		t.Fatal("expected 4, got ", total)
	}
}

func ExampleForEach() {
	list := []int{1, 1, 2}

	total := 0

	ForEach(list, func(item int) {
		total += item
	})

	fmt.Println(total)
	// Output: 4
}

func TestUnique_struct(t *testing.T) {
	list := []KeyVal{
		{
			Key:   "foo",
			Value: "bar",
		},
		{
			Key:   "foo",
			Value: "bar",
		},
		{
			Key:   "bin",
			Value: "baz",
		},
	}

	got := Unique(list)
	if len(got) != 2 {
		t.Fatal("expected 2 got ", len(got))
	}
	if got[0].Value != list[0].Value {
		t.Fatal("non matching value")
	}
	if got[1].Value != list[2].Value {
		t.Fatal("non matching value")
	}
}

func ExampleUnique() {
	list := []KeyVal{
		{
			Key:   "foo",
			Value: "bar",
		},
		{
			Key:   "foo",
			Value: "bar",
		},
		{
			Key:   "bin",
			Value: "baz",
		},
	}

	got := Unique(list)

	fmt.Println(got)
	// Output: [{bin baz}]
}

func TestFind_struct(t *testing.T) {
	list := []KeyVal{
		{
			Key:   "foo",
			Value: "bar",
		},
		{
			Key:   "bin",
			Value: "baz",
		},
	}

	got, ok := Find(list, func(item KeyVal) bool {
		return item.Value == "bar"
	})
	if !ok {
		t.Fatal("expected true, got false")
	}
	if got != list[0] {
		t.Fatal("expected first element, got ", got)
	}
}

func ExampleFind() {
	list := []KeyVal{
		{
			Key:   "foo",
			Value: "bar",
		},
		{
			Key:   "bin",
			Value: "baz",
		},
	}

	got, ok := Find(list, func(item KeyVal) bool {
		return item.Value == "bar"
	})
	fmt.Println(got, ok)
	// Output: {foo bar} true
}

func TestFind_struct_not_found(t *testing.T) {
	list := []KeyVal{
		{
			Key:   "foo",
			Value: "bar",
		},
		{
			Key:   "bin",
			Value: "baz",
		},
	}

	got, ok := Find(list, func(item KeyVal) bool {
		return item.Value == "silly"
	})
	if ok {
		t.Fatal("expected false, got true")
	}
	if got.Value != "" {
		t.Fatal("expected zero value , got ", got)
	}
}

func TestEvery_struct(t *testing.T) {
	list := []KeyVal{
		{
			Key:   "foo",
			Value: "bar",
		},
		{
			Key:   "bin",
			Value: "baz",
		},
	}
	if ok := Every(list, func(item KeyVal) bool {
		return item.Value != ""
	}); !ok {
		t.Fatal("expected true, got false")
	}
}

func ExampleEvery() {
	list := []KeyVal{
		{
			Key:   "foo",
			Value: "bar",
		},
		{
			Key:   "bin",
			Value: "baz",
		},
	}
	ok := Every(list, func(item KeyVal) bool {
		return item.Value != ""
	})

	fmt.Println(ok)
	// Output: true
}

func TestEvery_struct_no_match(t *testing.T) {
	list := []KeyVal{
		{
			Key:   "foo",
			Value: "bar",
		},
		{
			Key:   "bin",
			Value: "baz",
		},
	}
	if ok := Every(list, func(item KeyVal) bool {
		return item.Value == "frisky"
	}); ok {
		t.Fatal("expected false, got true")
	}
}

func TestReduce_int(t *testing.T) {
	list := []int{1, 2, 3}

	add := Reduce(list, func(prev, current int) int {
		return prev + current
	})

	got := add(1)

	assert.Equal(t, 7, got)
}

func ExampleReduce() {
	list := []int{1, 2, 3}

	add := Reduce(list, func(prev, current int) int {
		return prev + current
	})

	got := add(1)

	fmt.Println(got)
	// Output: 7
}

func TestUnion(t *testing.T) {
	id := uuid.New()
	now := time.Now()
	lists := [][]CmpUUID{
		{
			{
				id:   id,
				date: now,
			},
		},
		{
			{
				id:   id,
				date: now,
			},
		},
	}

	got := Union(lists...)

	assert.Equal(t, 1, len(got))
	assert.Equal(t, []CmpUUID{{id: id, date: now}}, got)
}

func ExampleUnion() {
	id := "1"
	value := "some value"
	lists := [][]KeyVal{
		{
			{
				Key:   id,
				Value: value,
			},
		},
		{
			{
				Key:   id,
				Value: value,
			},
		},
	}

	got := Union(lists...)

	fmt.Println(got)
	// Output: [{1 some value}]
}

func TestReverse(t *testing.T) {
	list := []int{1, 2, 3}

	got := Reverse(list)

	assert.Equal(t, []int{3, 2, 1}, got)
}

func ExampleReverse() {
	list := []int{1, 2, 3}

	got := Reverse(list)

	fmt.Println(got)
	// Output: [3 2 1]
}

func TestChunk_even_allocation(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}

	got := Chunk(list, 2)

	assert.Equal(t, 3, len(got))
	assert.Equal(t, []int{1, 2}, got[0])
	assert.Equal(t, []int{3, 4}, got[1])
	assert.Equal(t, []int{5, 6}, got[2])
}

func ExampleChunk() {
	list := []int{1, 2, 3, 4, 5, 6}

	got := Chunk(list, 2)

	fmt.Println(got)
	// Output: [[1 2] [3 4] [5 6]]
}

func TestChunk_odd_allocation(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7}

	got := Chunk(list, 2)

	assert.Equal(t, 4, len(got))
	assert.Equal(t, []int{1, 2}, got[0])
	assert.Equal(t, []int{3, 4}, got[1])
	assert.Equal(t, []int{5, 6}, got[2])
	assert.Equal(t, []int{7}, got[3])
}

func TestChunk_less_than_chunk_size(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7}

	got := Chunk(list, 22)

	assert.Equal(t, 1, len(got))
	assert.Equal(t, [][]int{list}, got)
}

func TestDifference(t *testing.T) {
	list1 := []int{1, 2, 3}
	list2 := []int{2, 3, 4}
	got := Difference(list1, list2)

	assert.Equal(t, []int{1, 4}, got)
}

func ExampleDifference() {
	list1 := []int{1, 2, 3}
	list2 := []int{2, 3, 4}
	got := Difference(list1, list2)

	fmt.Println(got)
	// Output: [1 4]
}

func TestDifference_no_difference(t *testing.T) {
	list1 := []int{1, 2, 3}
	list2 := []int{1, 2, 3}
	got := Difference(list1, list2)

	assert.Nil(t, got)
}

func TestDifference_key_val(t *testing.T) {
	list1 := []KeyVal{
		{
			Key:   "foo",
			Value: "bar",
		},
		{
			Key:   "bin",
			Value: "baz",
		},
	}
	list2 := []KeyVal{
		{
			Key:   "foo",
			Value: "bar",
		},
	}
	got := Difference(list1, list2)

	assert.Equal(t, []KeyVal{{
		Key:   "bin",
		Value: "baz",
	}}, got)
}

func TestWithout(t *testing.T) {
	list := []KeyVal{
		{Key: "foo", Value: "bar"},
		{Key: "bin", Value: "baz"},
	}

	got := Without(list, KeyVal{Key: "bin", Value: "baz"})

	assert.Equal(t, []KeyVal{{Key: "foo", Value: "bar"}}, got)
}

func ExampleWithout() {
	list := []KeyVal{
		{Key: "foo", Value: "bar"},
		{Key: "bin", Value: "baz"},
	}

	got := Without(list, KeyVal{Key: "bin", Value: "baz"})
}

func TestSome(t *testing.T) {
	list := []int{1, 2, 3}

	got := Some(list, func(i int) bool {
		return i == 2
	})
	assert.Equal(t, true, got)
}
