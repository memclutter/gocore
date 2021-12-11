package corecliapp

import (
	"fmt"
	"github.com/memclutter/gocore/pkg/coreslices"
	"github.com/memclutter/gocore/pkg/corestrings"
	"github.com/urfave/cli/v2"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// ConfigToFlags godoc
//
// Parse config structure and create cli flags for app
func ConfigToFlags(config interface{}) []cli.Flag {
	flags := make([]cli.Flag, 0)

	refConfig := reflect.ValueOf(config)
	refConfigType := refConfig.Type()

	for i := 0; i < refConfigType.NumField(); i++ {
		field := refConfigType.Field(i)
		name := strings.TrimSpace(field.Tag.Get("name"))
		value := field.Tag.Get("value")
		envVars := strings.Split(field.Tag.Get("envVars"), ",")
		envVars = coreslices.StringApply(envVars, func(i int, s string) string { return strings.ToLower(s) })
		envVars = coreslices.StringFilter(envVars, func(i int, s string) bool { return len(s) > 0 })

		// Set name as struct name lowerCamelCase
		if len(name) == 0 {
			name = corestrings.ToLowerFirst(field.Name)
		}

		// Set env vars as struct field name converted from CamelCase -> snake_case -> UPPER_SNAKE_CASE
		if len(envVars) == 0 {
			envVars = []string{
				strings.ToUpper(corestrings.CamelToSnake(field.Name)),
			}
		}

		switch v := refConfig.Field(i).Interface().(type) {
		case bool:
			defaultValue, err := strconv.ParseBool(value)
			if err != nil {
				log.Fatalf("invalid boolean value '%s' for config param '%s'", value, field.Name)
			}
			flags = append(flags, &cli.BoolFlag{
				Name:    name,
				Value:   defaultValue,
				EnvVars: envVars,
			})
		case time.Duration:
			defaultValue, err := time.ParseDuration(value)
			if err != nil {
				log.Fatalf("invalid duration value '%s' for config param '%s'", value, field.Name)
			}
			flags = append(flags, &cli.DurationFlag{
				Name: name,
				Value: defaultValue,
				EnvVars: envVars,
			})
		case int:
			defaultValue, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				log.Fatalf("invalid int value '%s' for config param '%s'", value, field.Name)
			}
			flags = append(flags, &cli.IntFlag{
				Name:    name,
				Value:   int(defaultValue),
				EnvVars: envVars,
			})
		case string:
			flags = append(flags, &cli.StringFlag{
				Name:    name,
				Value:   value,
				EnvVars: envVars,
			})
		default:
			log.Fatalf("unsuported config type %T for config param '%s'", v, field.Name)
		}
	}

	return flags
}

// ContextToConfig godoc
//
// Parse all cli flags and set config structure
func ContextToConfig(c *cli.Context, config interface{}) error {

	refConfig := reflect.ValueOf(config).Elem()
	refConfigType := refConfig.Type()

	for i := 0; i < refConfigType.NumField(); i++ {
		field := refConfigType.Field(i)
		fieldValue := refConfig.Field(i)
		name := strings.TrimSpace(field.Tag.Get("name"))
		value := field.Tag.Get("value")

		// Set name as struct name lowerCamelCase
		if len(name) == 0 {
			name = corestrings.ToLowerFirst(field.Name)
		}

		switch v := fieldValue.Interface().(type) {
		case bool:
			defaultValue, err := strconv.ParseBool(value)
			if err != nil {
				return fmt.Errorf("invalid boolean value '%s' for config param '%s': %v", value, field.Name, err)
			}
			if c.IsSet(name) {
				fieldValue.SetBool(c.Bool(name))
			} else {
				fieldValue.SetBool(defaultValue)
			}
		case time.Duration:
			defaultValue, err := time.ParseDuration(value)
			if err != nil {
				return fmt.Errorf("invalid duration value '%s' for config params '%s': %v", value, field.Name, err)
			}
			if c.IsSet(name) {
				fieldValue.Set(reflect.ValueOf(c.Duration(name)))
			} else {
				fieldValue.Set(reflect.ValueOf(defaultValue))
			}
		case int:
			defaultValue, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid int value '%s' for config param '%s': %v", value, field.Name, err)
			}
			if c.IsSet(name) {
				fieldValue.SetInt(int64(c.Int(name)))
			} else {
				fieldValue.SetInt(defaultValue)
			}
		case string:
			if c.IsSet(name) {
				fieldValue.SetString(c.String(name))
			} else {
				fieldValue.SetString(value)
			}
		default:
			return fmt.Errorf("unknown type %T", v)
		}
	}

	return nil
}
