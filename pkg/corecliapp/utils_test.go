package corecliapp

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCallRunCommand struct {
}

func (cmd TestCallRunCommand) Run() error {
	return nil
}

func Test_callRun(t *testing.T) {
	err := callRun(reflect.ValueOf(TestCallRunCommand{}))
	if err != nil {
		t.Fatalf("error call run, excepted no errors, actual have error: %v", err)
	}
}

type TestCallRunCommandWithError struct {
}

func (cmd TestCallRunCommandWithError) Run() error {
	return fmt.Errorf("test")
}

func Test_callRunWithError(t *testing.T) {
	exceptedErr := fmt.Errorf("test")
	err := callRun(reflect.ValueOf(TestCallRunCommandWithError{}))
	if err.Error() != exceptedErr.Error() {
		t.Fatalf("error call run, excepted have error '%s', actual '%s'", exceptedErr, err)
	}
}
