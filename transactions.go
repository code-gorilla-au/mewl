package mewl

import (
	"errors"
	"fmt"
)

type Txn[T any] struct {
	txnState TxnState[T]
	errors   []error
	steps    []TxnStep[T]

	// failFast - if set to true, the transaction will stop at the first error.
	failFast bool
	// verbose - if set to true, the transaction will log out the steps as they are run.
	verbose bool
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

type TxnOpts[T any] func(*Txn[T])

// NewTxn - creates a new transaction.
func NewTxn[T any](state T, opts ...TxnOpts[T]) *Txn[T] {

	t := &Txn[T]{
		txnState: TxnState[T]{
			state:       &state,
			currentStep: 0,
		},
	}

	for _, opt := range opts {
		opt(t)
	}

	return t
}

// Step - adds a step to the transaction workflow.
// All steps must have a handler and a rollback func.
func (t *Txn[T]) Step(handler TxnFunc[T], rollback TxnFunc[T]) *Txn[T] {
	t.steps = append(t.steps, TxnStep[T]{handler: handler, rollback: rollback})
	return t
}

// Run - runs the transaction.
// If an error occurs within one of the steps, it will rollback the transaction.

// Errors caught within the steps and rollback funcs will be able to be unwrapped and inspected using Unwrap() []error.
func (t *Txn[T]) Run() (T, error) {
	var err error

	t.log(fmt.Sprintf("starting transaction with %d steps", len(t.steps)))
	for index, step := range t.steps {

		t.txnState.currentStep = index
		logStep := index + 1
		t.log(fmt.Sprintf("step %d: executing", logStep))

		*t.txnState.state, err = step.handler(*t.txnState.state)
		if err != nil {
			t.log(fmt.Sprintf("step %d execution failed: %s, rolling back", logStep, err))

			errWithCtx := fmt.Errorf("step failed at step %d: %w", logStep, err)
			t.errors = append(t.errors, errWithCtx)

			if err := t.rollback(); err != nil {
				// fail fast stops the rollback early and returns first error
				t.errors = append(t.errors, err)
			}

			return *t.txnState.state, errors.Join(t.errors...)
		}

		t.log(fmt.Sprintf("step %d: complete", logStep))
	}

	return *t.txnState.state, nil
}

// rollback - rolls back the transaction.
// If failFast is set to true, it will stop at the first error on a rollback handler, otherwise it will continue.
func (t *Txn[T]) rollback() error {
	var err error
	for i := t.txnState.currentStep; i >= 0; i-- {
		logStep := i + 1
		step := t.steps[i]

		t.log(fmt.Sprintf("rollback step %d: executing", logStep))

		*t.txnState.state, err = step.rollback(*t.txnState.state)
		if err != nil {
			t.log(fmt.Sprintf("rollback step %d: failed: %s", logStep, err))

			errWithCtx := fmt.Errorf("rollback failed at step %d: %w", logStep, err)
			if t.failFast {
				return errWithCtx
			}

			// add it to the list, but continue with rollback
			t.errors = append(t.errors, errWithCtx)
		}

		t.log(fmt.Sprintf("rollback step %d: complete", logStep))
	}

	return nil
}

func (t *Txn[T]) log(msg string) {
	if !t.verbose {
		return
	}

	fmt.Println(msg)
}

// TxnOptFailFast - if set to true, the transaction will stop at the first error.
func TxnOptFailFast[T any]() TxnOpts[T] {
	return func(t *Txn[T]) {
		t.failFast = true
	}
}

// TxnOptVerbose - if set to true, the transaction will log out the steps as they are run.
func TxnOptVerbose[T any]() TxnOpts[T] {
	return func(t *Txn[T]) {
		t.verbose = true
	}
}
