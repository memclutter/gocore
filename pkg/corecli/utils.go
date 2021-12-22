package corecli

import (
	"fmt"
	"github.com/memclutter/gocore/pkg/coreslices"
	"github.com/memclutter/gocore/pkg/corestrings"
	"github.com/urfave/cli/v2"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// typeToFlag godoc
//
// Create cli.Flag for specified type.
func typeToFlag(i interface{}) (cli.Flag, error) {
	switch i.(type) {
	case bool:
		return &cli.BoolFlag{}, nil
	case int:
		return &cli.IntFlag{}, nil
	case int64:
		return &cli.Int64Flag{}, nil
	case float64:
		return &cli.Float64Flag{}, nil
	case uint:
		return &cli.UintFlag{}, nil
	case uint64:
		return &cli.Uint64Flag{}, nil
	case string:
		return &cli.StringFlag{}, nil
	case time.Duration:
		return &cli.DurationFlag{}, nil
	case time.Time:
		return &cli.TimestampFlag{}, nil
	case []int:
		return &cli.IntSliceFlag{}, nil
	case []int64:
		return &cli.Int64SliceFlag{}, nil
	case []float64:
		return &cli.Float64SliceFlag{}, nil
	case []string:
		return &cli.StringSliceFlag{}, nil
	default:
		return nil, fmt.Errorf("unsupported flag type %T", i)
	}
}

func generateFlagName(field reflect.StructField) string {
	name := strings.TrimSpace(field.Tag.Get("cli.flag.name"))
	if len(name) == 0 {
		name = corestrings.ToLowerFirst(field.Name)
	}
	return name
}

func generateFlagEnvVars(tag reflect.StructTag, name string) []string {
	envVars := strings.Split(tag.Get("cli.flag.envVars"), ",")
	envVars = coreslices.StringApply(envVars, func(i int, s string) string { return strings.TrimSpace(s) })
	envVars = coreslices.StringFilter(envVars, func(i int, s string) bool { return len(s) > 0 })
	if len(envVars) == 0 {
		envVars = []string{strings.ToUpper(corestrings.CamelToSnake(name))}
	}
	if envVars[0] == "-" {
		envVars = make([]string, 0)
	}
	return envVars
}

func generateFlagValue(tag reflect.StructTag, i interface{}) (interface{}, error) {
	tagValue := strings.TrimSpace(tag.Get("cli.flag.value"))
	switch i.(type) {
	case bool:
		if len(tagValue) == 0 {
			return false, nil
		}
		return strconv.ParseBool(tagValue)
	case int:
		if len(tagValue) == 0 {
			return int(0), nil
		}
		return strconv.Atoi(tagValue)
	case int64:
		if len(tagValue) == 0 {
			return int64(0), nil
		}
		return strconv.ParseInt(tagValue, 10, 64)
	case float64:
		if len(tagValue) == 0 {
			return float64(0), nil
		}
		return strconv.ParseFloat(tagValue, 64)
	case uint:
		if len(tagValue) == 0 {
			return uint(0), nil
		}
		v, err := strconv.ParseUint(tagValue, 10, 64)
		return uint(v), err
	case uint64:
		if len(tagValue) == 0 {
			return uint64(0), nil
		}
		return strconv.ParseUint(tagValue, 10, 64)
	case string:
		return tagValue, nil
	case time.Duration:
		if len(tagValue) == 0 {
			return time.Duration(0), nil
		}
		return time.ParseDuration(tagValue)
	case time.Time:
		if len(tagValue) == 0 {
			return nil, nil
		}
		v, err := time.Parse(time.RFC3339, tagValue)
		if err != nil {
			return nil, err
		}
		return cli.NewTimestamp(v), nil
	case []int:
		if len(tagValue) == 0 {
			return nil, nil
		}
		v := make([]int, 0)
		for i, s := range stringToStringSlice(tagValue) {
			e, err := strconv.Atoi(s)
			if err != nil {
				return nil, fmt.Errorf("[%d]int: %v", i, err)
			}
			v = append(v, e)
		}
		return cli.NewIntSlice(v...), nil
	case []int64:
		if len(tagValue) == 0 {
			return nil, nil
		}
		v := make([]int64, 0)
		for i, s := range stringToStringSlice(tagValue) {
			e, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("[%d]int64: %v", i, err)
			}
			v = append(v, e)
		}
		return cli.NewInt64Slice(v...), nil
	case []float64:
		if len(tagValue) == 0 {
			return nil, nil
		}
		v := make([]float64, 0)
		for i, s := range stringToStringSlice(tagValue) {
			e, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return nil, fmt.Errorf("[%d]float64: %v", i, err)
			}
			v = append(v, e)
		}
		return cli.NewFloat64Slice(v...), nil
	case []string:
		if len(tagValue) == 0 {
			return nil, nil
		}
		return cli.NewStringSlice(stringToStringSlice(tagValue)...), nil
	default:
		return nil, fmt.Errorf("unsupported flag type %T", i)
	}
}

func generateFlagRequired(tag reflect.StructTag) (bool, error) {
	s := strings.TrimSpace(tag.Get("cli.flag.required"))
	if len(s) == 0 {
		return false, nil
	}
	return strconv.ParseBool(s)
}

func stringToStringSlice(s string) []string {
	ss := strings.Split(s, ",")
	ss = coreslices.StringApply(ss, func(i int, s string) string { return strings.TrimSpace(s)})
	ss = coreslices.StringFilter(ss, func(i int, s string) bool { return len(s) > 0 })
	return 	ss
}