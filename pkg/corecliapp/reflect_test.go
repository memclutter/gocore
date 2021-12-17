package corecliapp

import (
	"reflect"
	"testing"
)

func Test_typeOfPtr(t *testing.T)  {
	type TestStruct struct {
		Prop string
	}

	tables := []struct{
		i reflect.Value
		result reflect.Type
	}{
		{
			i: reflect.ValueOf(TestStruct{Prop: "value"}),
			result: reflect.TypeOf(TestStruct{Prop: "value"}),
		},
		{
			i: reflect.ValueOf(&TestStruct{Prop: "value"}),
			result: reflect.TypeOf(TestStruct{Prop: "value"}),
		},
	}

	for _, table := range tables {
		result := typeOfPtr(table.i)

		if result.String() != table.result.String() {
			t.Fatalf("assert result type failed, excepted '%s', actual '%s'", table.result, result)
		}
	}
}

func Test_valueOfPtr(t *testing.T)  {
	testString := "string"

	tables := []struct{
		v reflect.Value
		result reflect.Value
	}{
		{
			v: reflect.ValueOf("string"),
			result: reflect.ValueOf("string"),
		},
		{
			v: reflect.ValueOf(&testString),
			result: reflect.ValueOf(testString),
		},
	}

	for _, table := range tables {
		result := valueOfPtr(table.v)
		if result.Type().Kind() != table.result.Type().Kind() {
			t.Errorf("assert kind of type failed, excepted '%s', actual '%s'", table.result.Type().Kind(), result.Type().Kind())
		}
	}
}