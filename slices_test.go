package mewl

import (
	"fmt"
	"testing"
	"time"

	"github.com/code-gorilla-au/odize"
	"github.com/google/uuid"
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

	group := odize.NewGroup(t, nil)

	group.AfterEach(func() {
		total = 0
		list = []int{1, 1, 2}
	})

	err := group.
		Test("should add total of 4", func(t *testing.T) {
			ForEach(list, func(item int, _ int, _ []int) {
				total += item
			})
			odize.AssertEqual(t, 4, total)
		}).
		Test("should add total of -4", func(t *testing.T) {
			ForEach(list, func(item int, _ int, _ []int) {
				total -= item
			})
			odize.AssertEqual(t, -4, total)
		}).
		Run()

	odize.AssertNoError(t, err)
}

func ExampleForEach() {
	list := []int{1, 1, 2}

	total := 0

	ForEach(list, func(item int, _ int, _ []int) {
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
	// Output: [{foo bar} {bin baz}]
}

func TestFind_struct(t *testing.T) {
	var list []KeyVal
	group := odize.NewGroup(t, nil)
	group.BeforeEach(func() {
		list = []KeyVal{
			{
				Key:   "foo",
				Value: "bar",
			},
			{
				Key:   "bin",
				Value: "baz",
			},
		}
	})
	group.AfterEach(func() {
		list = []KeyVal{
			{
				Key:   "foo",
				Value: "bar",
			},
			{
				Key:   "bin",
				Value: "baz",
			},
		}
	})
	err := group.
		Test("should return the first item that matches the predicate", func(t *testing.T) {
			got, ok := Find(list, func(item KeyVal, _ int, _ []KeyVal) bool {
				return item.Value == "bar"
			})
			odize.AssertTrue(t, ok)
			odize.AssertEqual(t, got, list[0])
		}).
		Test("should false if not able to find element within slice", func(t *testing.T) {
			got, ok := Find(list, func(item KeyVal, _ int, _ []KeyVal) bool {
				return item.Value == "silly"
			})
			odize.AssertFalse(t, ok)
			odize.AssertNil(t, got)
		}).
		Run()
	odize.AssertNoError(t, err)
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

	got, ok := Find(list, func(item KeyVal, _ int, _ []KeyVal) bool {
		return item.Value == "bar"
	})
	fmt.Println(got, ok)
	// Output: {foo bar} true
}

func TestEvery_struct(t *testing.T) {
	var list []KeyVal
	group := odize.NewGroup(t, nil)
	group.BeforeEach(func() {
		list = []KeyVal{
			{
				Key:   "foo",
				Value: "bar",
			},
			{
				Key:   "bin",
				Value: "baz",
			},
		}
	})
	group.AfterEach(func() {
		list = []KeyVal{
			{
				Key:   "foo",
				Value: "bar",
			},
			{
				Key:   "bin",
				Value: "baz",
			},
		}
	})

	err := group.
		Test("should return true if every item is not empty", func(t *testing.T) {
			ok := Every(list, func(item KeyVal, _ int, _ []KeyVal) bool {
				return item.Value != ""
			})
			odize.AssertTrue(t, ok)
		}).
		Test("should return false if an item does not match", func(t *testing.T) {
			ok := Every(list, func(item KeyVal, _ int, _ []KeyVal) bool {
				return item.Value == "frisky"
			})
			odize.AssertFalse(t, ok)
		}).
		Test("should mutate original list", func(t *testing.T) {
			ok := Every(list, func(item KeyVal, idx int, originalList []KeyVal) bool {
				originalList[idx] = KeyVal{
					Key: "override",
				}
				return true
			})
			odize.AssertTrue(t, ok)
			odize.AssertEqual(t, list[0].Key, "override")
			odize.AssertEqual(t, list[1].Key, "override")
		}).
		Run()

	odize.AssertNoError(t, err)

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
	// Note: the original slice is passed into the predicate function
	ok := Every(list, func(item KeyVal, i int, originalSlice []KeyVal) bool {
		return item.Value != ""
	})

	fmt.Println(ok)
	// Output: true
}

func TestReduce_int(t *testing.T) {
	list := []int{1, 2, 3}

	add := Reduce(list, func(prev, current int) int {
		return prev + current
	})

	got := add(1)

	odize.AssertEqual(t, 7, got)
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

	odize.AssertEqual(t, 1, len(got))
	odize.AssertEqual(t, []CmpUUID{{id: id, date: now}}, got)
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

	odize.AssertEqual(t, []int{3, 2, 1}, got)
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

	odize.AssertEqual(t, 3, len(got))
	odize.AssertEqual(t, []int{1, 2}, got[0])
	odize.AssertEqual(t, []int{3, 4}, got[1])
	odize.AssertEqual(t, []int{5, 6}, got[2])
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

	odize.AssertEqual(t, 4, len(got))
	odize.AssertEqual(t, []int{1, 2}, got[0])
	odize.AssertEqual(t, []int{3, 4}, got[1])
	odize.AssertEqual(t, []int{5, 6}, got[2])
	odize.AssertEqual(t, []int{7}, got[3])
}

func TestChunk_less_than_chunk_size(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7}

	got := Chunk(list, 22)

	odize.AssertEqual(t, 1, len(got))
	odize.AssertEqual(t, [][]int{list}, got)
}

func TestDifference(t *testing.T) {
	list1 := []int{1, 2, 3}
	list2 := []int{2, 3, 4}
	got := Difference(list1, list2)

	odize.AssertEqual(t, []int{1, 4}, got)
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
	var expected []int

	odize.AssertEqual(t, expected, got)

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

	odize.AssertEqual(t, []KeyVal{{
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

	odize.AssertEqual(t, []KeyVal{{Key: "foo", Value: "bar"}}, got)
}

func ExampleWithout() {
	list := []KeyVal{
		{Key: "foo", Value: "bar"},
		{Key: "bin", Value: "baz"},
	}

	got := Without(list, KeyVal{Key: "bin", Value: "baz"})

	fmt.Println(got)
	// Output: [{foo bar}]
}

func TestSome(t *testing.T) {
	list := []int{1, 2, 3}

	group := odize.NewGroup(t, nil)

	group.AfterEach(func() {
		list = []int{1, 2, 3}
	})

	err := group.
		Test("should return true if any item matches predicate", func(t *testing.T) {
			got := Some(list, func(item, index int, slice []int) bool {
				return item == 2
			})
			odize.AssertTrue(t, got)
		}).
		Test("should return true if all items match", func(t *testing.T) {
			got := Some(list, func(item, index int, slice []int) bool {
				return item <= 3
			})
			odize.AssertTrue(t, got)
		}).
		Test("should return false if no matches", func(t *testing.T) {
			got := Some(list, func(item, index int, slice []int) bool {
				return item > 10
			})
			odize.AssertFalse(t, got)
		}).
		Run()

	odize.AssertNoError(t, err)
}

func ExampleSome() {
	list := []int{1, 2, 3}

	got := Some(list, func(item, index int, slice []int) bool {
		return item == 2
	})

	fmt.Println(got)
	// Output: true
}
