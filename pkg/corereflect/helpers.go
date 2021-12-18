package corereflect

import (
	"reflect"
)

// PtrTypeOf godoc
//
// Process interface object and return reflect.TypeOf() for ptr or struct
func PtrTypeOf(v reflect.Value) reflect.Type {
	typeOf := v.Type()
	kind := typeOf.Kind()

	switch kind {
	case reflect.Ptr:
		return typeOf.Elem()
	default:
		return typeOf
	}
}

// PtrValueOf godoc
//
// Process interface object and return reflect.Value for ptr or non ptr
func PtrValueOf(v reflect.Value) reflect.Value {
	typeOf := v.Type()
	kind := typeOf.Kind()

	switch kind {
	case reflect.Ptr:
		return v.Elem()
	default:
		return v
	}
}
