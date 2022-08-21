package mewl

import (
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

func TestUnion(t *testing.T) {
	id := uuid.New()
	now := time.Now()
	lists := [][]CmpUUID{
		{
			{
				id:   id,
				date: now,
			},
			{
				id:   id,
				date: now,
			},
		},
	}

	got := Union(lists...)

	assert.Equal(t, 1, len(got))
}

func TestReverse(t *testing.T) {
	list := []int{1, 2, 3}

	got := Reverse(list)

	assert.Equal(t, []int{3, 2, 1}, got)
}
