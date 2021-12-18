package corecliapp

import (
	"github.com/memclutter/gocore/pkg/corereflect"
	"github.com/urfave/cli/v2"
	"reflect"
)

func setFlags(c *cli.Context, rFlags reflect.Value) {
	rFlags = corereflect.PtrValueOf(rFlags)
	rtFlags := corereflect.PtrTypeOf(rFlags)
	for j := 0; j < rtFlags.NumField(); j++ {
		rfField := rtFlags.Field(j)
		rField := rFlags.Field(j)

		if rfField.Anonymous {
			setFlags(c, rField)
			continue
		}

		name := rfField.Tag.Get("cli.flag.name")
		switch rfField.Type.Kind() {
		case reflect.String:
			rField.SetString(c.String(name))
		case reflect.Bool:
			rField.SetBool(c.Bool(name))
		case reflect.Int:
			rField.SetInt(int64(c.Int(name)))
		case reflect.Int64:
			rField.SetInt(int64(c.Int64(name)))
		}
	}
}

func callRun(rCommand reflect.Value) error {
	rResult := rCommand.MethodByName("Run").Call([]reflect.Value{})
	if len(rResult) == 0 || rResult[0].IsNil() {
		return nil
	}
	return rResult[0].Interface().(error)
}
