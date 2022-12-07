package main

import (
	"errors"
	"fmt"
)

var (
	ErrValidation = errors.New("validation failed")
	ErrSignal     = errors.New("received signal")
)

//stepErr represent a class of errors associated with CI steps,
//step to record the step name in an error; a message msg that
//describe the condition; and a cause to store the underlying error
//that caused this step error
type stepErr struct {
	step  string
	msg   string
	cause error
}

func (s *stepErr) Error() string {
	return fmt.Sprintf("Step: %q: %s: Cause: %v", s.step, s.msg, s.cause)
}

func (s *stepErr) Is(target error) bool {
	t, ok := target.(*stepErr)
	if !ok {
		return false
	}

	return t.step == s.step
}

//Unwrap try to unwrap the error to see
//if an underlying error matches the target,
//it returns the error stored in the cause field
func (s *stepErr) Unwrap() error {
	return s.cause
}
