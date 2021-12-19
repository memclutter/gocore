package corecliapp

import (
	"fmt"
	"github.com/memclutter/gocore/pkg/corereflect"
	"github.com/memclutter/gocore/pkg/coreslices"
	"github.com/memclutter/gocore/pkg/corestrings"
	"github.com/urfave/cli/v2"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// createFlag godoc
//
// Create urfave/cli Flag from reflect struct field, preset default value from reflect value or struct tag
//
// Struct tags:
// - cli.flag.name - flag name, if empty use struct field name in lowerCamelCase
// - cli.flag.value - flag value (default for usage and help), if empty use reflect value
// - cli.flag.envVars - flag environment variables, if empty use struct field name in SNAKE_CASE
// - @TODO usage, help and etc urfave/cli Flag struct field
func createFlag(structField reflect.StructField, value reflect.Value) (cli.Flag, error) {

	// cli.flag.name
	name := strings.TrimSpace(structField.Tag.Get("cli.flag.name"))
	if len(name) == 0 {
		name = corestrings.ToLowerFirst(structField.Name)
	}

	// cli.flag.value
	envVars := strings.Split(structField.Tag.Get("cli.flag.envVars"), ",")
	envVars = coreslices.StringApply(envVars, func(i int, s string) string { return strings.TrimSpace(s) })
	envVars = coreslices.StringFilter(envVars, func(i int, s string) bool { return len(s) > 0 })
	if len(envVars) == 0 {
		envVars = []string{
			strings.ToUpper(corestrings.CamelToSnake(name)),
		}
	} else if len(envVars) == 1 && envVars[0] == "-" {
		envVars = []string{}
	}

	// cli.flag.value
	var flagValue reflect.Value
	tagValue := strings.TrimSpace(structField.Tag.Get("cli.flag.value"))
	if len(tagValue) == 0 {
		flagValue = value
	} else {
		switch value.Interface().(type) {
		case string:
			flagValue = reflect.ValueOf(tagValue)
		case time.Duration:
			v, err := time.ParseDuration(tagValue)
			if err != nil {
				return nil, fmt.Errorf("error parse duration '%s': %v", tagValue, err)
			}
			flagValue = reflect.ValueOf(v)
		case bool:
			v, err := strconv.ParseBool(tagValue)
			if err != nil {
				return nil, fmt.Errorf("error parse bool '%s': %v", tagValue, err)
			}
			flagValue = reflect.ValueOf(v)
		case float64:
			v, err := strconv.ParseFloat(tagValue, 64)
			if err != nil {
				return nil, fmt.Errorf("error parse float '%s': %v", tagValue, err)
			}
			flagValue = reflect.ValueOf(v)
		case time.Time:
			layout := time.RFC3339
			v, err := time.Parse(layout, tagValue)
			if err != nil {
				return nil, fmt.Errorf("error parse time.Time '%s' (layout '%s'): %v", tagValue, layout, err)
			}
			flagValue = reflect.ValueOf(v)
		case int:
			v, err := strconv.ParseInt(tagValue, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("error parse int '%s': %v", tagValue, err)
			}
			flagValue = reflect.ValueOf(int(v))
		case int64:
			v, err := strconv.ParseInt(tagValue, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("error parse int64 '%s': %v", tagValue, err)
			}
			flagValue = reflect.ValueOf(v)
		case uint:
			v, err := strconv.ParseUint(tagValue, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("error parse uint '%s': %v", tagValue, err)
			}
			flagValue = reflect.ValueOf(uint(v))
		case uint64:
			v, err := strconv.ParseUint(tagValue, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("error parse uint64 '%s': %v", tagValue, err)
			}
			flagValue = reflect.ValueOf(v)
		}
	}

	// create cli.Flag
	switch value.Interface().(type) {
	case string:
		return &cli.StringFlag{
			Name:    name,
			EnvVars: envVars,
			Value:   flagValue.String(),
		}, nil
	case time.Duration:
		return &cli.DurationFlag{
			Name:    name,
			EnvVars: envVars,
			Value:   time.Duration(flagValue.Int()),
		}, nil
	case bool:
		return &cli.BoolFlag{
			Name:    name,
			EnvVars: envVars,
			Value:   flagValue.Bool(),
		}, nil
	case float64:
		return &cli.Float64Flag{
			Name:    name,
			EnvVars: envVars,
			Value:   flagValue.Float(),
		}, nil
	case time.Time:
		return &cli.TimestampFlag{
			Name:    name,
			EnvVars: envVars,
			Value:   cli.NewTimestamp(flagValue.Interface().(time.Time)),
		}, nil
	case int:
		return &cli.IntFlag{
			Name:    name,
			EnvVars: envVars,
			Value:   int(flagValue.Int()),
		}, nil
	case int64:
		return &cli.Int64Flag{
			Name:    name,
			EnvVars: envVars,
			Value:   flagValue.Int(),
		}, nil
	case uint:
		return &cli.UintFlag{
			Name:    name,
			EnvVars: envVars,
			Value:   uint(flagValue.Uint()),
		}, nil
	case uint64:
		return &cli.Uint64Flag{
			Name:    name,
			EnvVars: envVars,
			Value:   flagValue.Uint(),
		}, nil
	default:
		return nil, fmt.Errorf("unsupport flag type '%T' for field '%s'", value.Interface(), structField.Name)
	}
}

// createFlags godoc
//
// Process struct tags and create urfave/cli Flag slice
func createFlags(i interface{}, flags []cli.Flag) ([]cli.Flag, error) {
	value := corereflect.PtrValueOf(reflect.ValueOf(i))
	valueType := corereflect.PtrTypeOf(value)

	// Bypass non struct type kinds
	if valueType.Kind() != reflect.Struct {
		return flags, nil
	}

	for i := 0; i < valueType.NumField(); i++ {
		structField := valueType.Field(i)

		// Bypass embedded fields
		if structField.Anonymous {
			continue
		}

		flag, err := createFlag(structField, value.Field(i))
		if err != nil {
			return flags, err
		}

		flags = append(flags, flag)
	}

	return flags, nil
}

// setFlags godoc
//
// Parse cli context and set all flags
func setFlags(c *cli.Context, rFlags reflect.Value) error {
	rFlags = corereflect.PtrValueOf(rFlags)
	rtFlags := corereflect.PtrTypeOf(rFlags)
	for j := 0; j < rtFlags.NumField(); j++ {
		rfField := rtFlags.Field(j)
		rField := rFlags.Field(j)

		if rfField.Anonymous {
			if err := setFlags(c, rField); err != nil {
				return err
			}
			continue
		}

		name := strings.TrimSpace(rfField.Tag.Get("cli.flag.name"))
		if len(name) == 0 {
			name = corestrings.ToLowerFirst(rfField.Name)
		}

		switch rField.Interface().(type) {
		case string:
			rField.SetString(c.String(name))
		case time.Duration:
			rField.Set(reflect.ValueOf(c.Duration(name)))
		case bool:
			rField.SetBool(c.Bool(name))
		case float64:
			rField.SetFloat(c.Float64(name))
		case time.Time:
			rField.Set(corereflect.PtrValueOf(reflect.ValueOf(c.Timestamp(name))))
		case int:
			rField.SetInt(int64(c.Int(name)))
		case int64:
			rField.SetInt(c.Int64(name))
		case uint:
			rField.SetUint(uint64(c.Uint(name)))
		case uint64:
			rField.SetUint(c.Uint64(name))
		default:
			return fmt.Errorf("unsupported type %T for set flag '%s'", rField.Interface(), rfField.Name)
		}
	}

	return nil
}
