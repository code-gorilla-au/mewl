package mewl

import (
	"fmt"
	"testing"

	"github.com/code-gorilla-au/odize"
)

func TestTxn(t *testing.T) {
	type testState struct {
		Name string
	}

	state := testState{Name: "hello"}

	group := odize.NewGroup(t, nil)
	group.AfterEach(func() {
		state = testState{Name: "hello"}
	})

	err := group.
		Test("should run all steps and return result", func(t *testing.T) {
			txn := NewTxn(state)
			result, err := txn.Step(
				func(ts testState) (testState, error) {
					ts.Name = "world"
					return ts, nil
				},
				func(ts testState) (testState, error) {
					ts.Name = "failed"
					return ts, nil
				},
			).Run()
			odize.AssertNoError(t, err)

			odize.AssertEqual(t, "world", result.Name)

		}).
		Test("should maintain state between states", func(t *testing.T) {
			txn := NewTxn(state)
			result, err := txn.
				Step(
					func(ts testState) (testState, error) {
						ts.Name = "bin"
						return ts, nil
					},
					func(ts testState) (testState, error) {
						ts.Name = "failed"
						return ts, nil
					},
				).
				Step(
					func(ts testState) (testState, error) {
						odize.AssertEqual(t, "bin", ts.Name)

						ts.Name = "world"
						return ts, nil
					},
					func(ts testState) (testState, error) {
						ts.Name = "failed"
						return ts, nil
					},
				).
				Run()
			odize.AssertNoError(t, err)

			odize.AssertEqual(t, "world", result.Name)

		}).
		Test("should fail and return error", func(t *testing.T) {
			txn := NewTxn(state)
			result, err := txn.Step(
				func(ts testState) (testState, error) {
					ts.Name = "world"
					return ts, fmt.Errorf("expected failure")
				},
				func(ts testState) (testState, error) {
					ts.Name = "failed"
					return ts, nil
				},
			).Run()
			odize.AssertEqual(t, err.Error(), fmt.Errorf("step failed at step 1: expected failure").Error())

			odize.AssertEqual(t, "failed", result.Name)

		}).
		Test("should not execute subsequent test on first failure", func(t *testing.T) {
			nextStepCall := 0
			txn := NewTxn(state)
			result, err := txn.
				Step(
					func(ts testState) (testState, error) {
						ts.Name = "world"
						return ts, fmt.Errorf("expected failure")
					},
					func(ts testState) (testState, error) {
						ts.Name = "failed"
						return ts, nil
					},
				).
				Step(
					func(ts testState) (testState, error) {
						nextStepCall++

						ts.Name = "bin"
						return ts, nil
					},
					func(ts testState) (testState, error) {
						ts.Name = "failed2"
						return ts, nil
					},
				).
				Run()
			odize.AssertEqual(t, fmt.Errorf("step failed at step 1: expected failure").Error(), err.Error())

			odize.AssertEqual(t, "failed", result.Name)
			odize.AssertEqual(t, 0, nextStepCall)

		}).
		Test("should execute both rollbacks", func(t *testing.T) {
			rollbackCall := 0
			txn := NewTxn(state)
			_, err := txn.
				Step(
					func(ts testState) (testState, error) {
						ts.Name = "world"
						return ts, nil
					},
					func(ts testState) (testState, error) {
						rollbackCall++

						return ts, nil
					},
				).
				Step(
					func(ts testState) (testState, error) {
						ts.Name = "bin"
						return ts, fmt.Errorf("expected failure")
					},
					func(ts testState) (testState, error) {
						rollbackCall++
						return ts, nil
					},
				).
				Run()
			odize.AssertEqual(t, fmt.Errorf("step failed at step 2: expected failure").Error(), err.Error())

			odize.AssertEqual(t, 2, rollbackCall)

		}).
		Run()

	odize.AssertNoError(t, err)
}
