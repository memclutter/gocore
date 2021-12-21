package corecli

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"reflect"
	"strings"
)

// defineFlags godoc
//
// Define urfave/cli flag slice from some golang struct. Support struct tags
// - `cli.flag.name` name of flag, default lowerCamelCase(structField.Name)
// - `cli.flag.usage` usage for flag
// - `cli.flag.value` default value for flag
// - `cli.flag.envVars` coma separated list of environment vars, default structField.Name
// - `cli.flag.required` set flag is required
func defineFlags(i interface{}) ([]cli.Flag, error) {
	var err error
	flags := make([]cli.Flag, 0)

	typeOf := reflect.TypeOf(i)
	valueOf := reflect.ValueOf(i)
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)

		// Bypass embedded struct fields
		if field.Anonymous {
			continue
		}

		// Bypass user struct tag `cli.flag:"-"`
		if field.Tag.Get("cli.flag") == "-" {
			continue
		}

		fieldValueOf := valueOf.Field(i)
		flag, err := typeToFlag(fieldValueOf.Interface())
		if err != nil {
			return nil, fmt.Errorf("define field %s error: %v", field.Name, err)
		}

		name := generateFlagName(field)
		envVars := generateFlagEnvVars(field.Tag, name)
		usage := strings.TrimSpace(field.Tag.Get("cli.flag.usage"))
		value, err := generateFlagValue(field.Tag, fieldValueOf.Interface())
		if err != nil {
			return nil, fmt.Errorf("define field %s error: %v", field.Name, err)
		}

		flagValueOf := reflect.ValueOf(flag).Elem()
		flagValueOf.FieldByName("Name").SetString(name)
		flagValueOf.FieldByName("EnvVars").Set(reflect.ValueOf(envVars))
		flagValueOf.FieldByName("Usage").SetString(usage)
		if value != nil {
			flagValueOf.FieldByName("Value").Set(reflect.ValueOf(value))
		}

		flags = append(flags, flag)
	}

	return flags, err
}
