package mewl

import (
	"errors"
	"fmt"
)

type Txn[T any] struct {
	txnState TxnState[T]
	errors   []error
	Steps    []TxnStep[T]
	failFast bool
}

type TxnState[T any] struct {
	state       *T
	currentStep int
}

type TxnFunc[T any] func(T) (T, error)

type TxnStep[T any] struct {
	handler  TxnFunc[T]
	rollback TxnFunc[T]
}

// NewTxn - creates a new transaction.
//
// Example:
//
// txn := NewTxn(state)
//
// result, err := txn.
//
//	Step(
//		func(s state) (state, error) {
//		state + 1
//		},
//		func(s state) (state, error) {
//			state - 1
//		}).Run()
func NewTxn[T any](state T) *Txn[T] {
	return &Txn[T]{
		txnState: TxnState[T]{
			state:       &state,
			currentStep: 0,
		},
	}
}

// Step - adds a step to the transaction workflow.
// All steps must have a handler and a rollback func.
func (t *Txn[T]) Step(handler TxnFunc[T], rollback TxnFunc[T]) *Txn[T] {
	t.Steps = append(t.Steps, TxnStep[T]{handler: handler, rollback: rollback})
	return t
}

// Run - runs the transaction.
// If an error occurs within one of the steps, it will rollback the transaction.

// Errors caught within the steps and rollback funcs will be able to be unwrapped and inspected using Unwrap() []error.
func (t *Txn[T]) Run() (T, error) {
	var err error

	for index, step := range t.Steps {
		fmt.Println("count", index)
		t.txnState.currentStep = index

		*t.txnState.state, err = step.handler(*t.txnState.state)

		if err != nil {
			errWithCtx := fmt.Errorf("step failed at step %d: %w", index+1, err)
			t.errors = append(t.errors, errWithCtx)

			if err := t.rollback(); err != nil {
				// fail fast stops the rollback early and returns first error
				t.errors = append(t.errors, err)
			}

			return *t.txnState.state, errors.Join(t.errors...)
		}
	}

	return *t.txnState.state, nil
}

// rollback - rolls back the transaction.
// If failFast is set to true, it will stop at the first error on a rollback handler, otherwise it will continue.
func (t *Txn[T]) rollback() error {
	var err error
	for i := t.txnState.currentStep; i >= 0; i-- {
		step := t.Steps[i]

		*t.txnState.state, err = step.rollback(*t.txnState.state)
		if err != nil {
			errWithCtx := fmt.Errorf("rollback failed at step %d: %w", i+1, err)
			if t.failFast {
				return errWithCtx
			}

			// add it to the list, but continue with rollback
			t.errors = append(t.errors, errWithCtx)
		}
	}

	return nil
}
