package corecli

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"reflect"
	"strings"
	"time"
)

// GenerateFlags godoc
//
// Define urfave/cli flag slice from some golang struct. Support struct tags
// - `cli.flag.name` name of flag, default lowerCamelCase(structField.Name)
// - `cli.flag.usage` usage for flag
// - `cli.flag.value` default value for flag
// - `cli.flag.envVars` coma separated list of environment vars, default structField.Name
// - `cli.flag.required` set flag is required
func GenerateFlags(i interface{}) ([]cli.Flag, error) {
	var err error
	flags := make([]cli.Flag, 0)

	valueOf := reflect.Indirect(reflect.ValueOf(i))
	typeOf := valueOf.Type()
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
		required, err := generateFlagRequired(field.Tag)
		if err != nil {
			return nil, fmt.Errorf("define field %s error: %v", field.Name, err)
		}

		flagValueOf := reflect.ValueOf(flag).Elem()
		flagValueOf.FieldByName("Name").SetString(name)
		flagValueOf.FieldByName("Required").SetBool(required)
		if len(envVars) > 0 {
			flagValueOf.FieldByName("EnvVars").Set(reflect.ValueOf(envVars))
		}
		if len(usage) > 0 {
			flagValueOf.FieldByName("Usage").SetString(usage)
		}
		if value != nil {
			flagValueOf.FieldByName("Value").Set(reflect.ValueOf(value))
		}

		flags = append(flags, flag)
	}

	return flags, err
}

// LoadFlags godoc
//
// Extract and load values from urfave/cli context
func LoadFlags(i reflect.Value, c *cli.Context) error {

	valueOf := reflect.Indirect(i)
	typeOf := valueOf.Type()
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		fieldValueOf := valueOf.Field(i)

		// Recursive process embedded field
		if field.Anonymous {
			if err := LoadFlags(fieldValueOf, c); err != nil {
				return err
			}
			continue
		}

		// Bypass user struct tag `cli.flag:"-"`
		if field.Tag.Get("cli.flag") == "-" {
			continue
		}

		name := generateFlagName(field)
		switch fieldValueOf.Interface().(type) {
		case bool:
			fieldValueOf.SetBool(c.Bool(name))
		case int:
			fieldValueOf.SetInt(int64(c.Int(name)))
		case int64:
			fieldValueOf.SetInt(c.Int64(name))
		case uint:
			fieldValueOf.SetUint(uint64(c.Uint(name)))
		case uint64:
			fieldValueOf.SetUint(c.Uint64(name))
		case float64:
			fieldValueOf.SetFloat(c.Float64(name))
		case string:
			fieldValueOf.SetString(c.String(name))
		case time.Duration:
			fieldValueOf.Set(reflect.ValueOf(c.Duration(name)))
		default:
			return fmt.Errorf("unsupported flag type %T", fieldValueOf.Interface())
		}
	}

	return nil
}
