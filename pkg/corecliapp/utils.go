package corecliapp

import (
	"reflect"
)

func callRun(rCommand reflect.Value) error {
	rResult := rCommand.MethodByName("Run").Call([]reflect.Value{})
	if len(rResult) == 0 || rResult[0].IsNil() {
		return nil
	}
	return rResult[0].Interface().(error)
}
