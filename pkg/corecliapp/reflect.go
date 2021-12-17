package corecliapp

import (
	"reflect"
)

// typeOfPtr godoc
//
// Process interface object and return reflect.TypeOf() for ptr or struct
func typeOfPtr(v reflect.Value) reflect.Type {
	typeOf := v.Type()
	kind := typeOf.Kind()

	switch kind {
	case reflect.Ptr:
		return typeOf.Elem()
	default:
		return typeOf
	}
}

// valueOfPtr godoc
//
// Process interface object and return reflect.Value for ptr or non ptr
func valueOfPtr(v reflect.Value) reflect.Value {
	typeOf := v.Type()
	kind := typeOf.Kind()

	switch kind {
	case reflect.Ptr:
		return v.Elem()
	default:
		return v
	}
}
