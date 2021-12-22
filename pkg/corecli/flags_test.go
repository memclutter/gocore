package corecli

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"reflect"
	"testing"
	"time"
)

func TestGenerateFlags(t *testing.T) {
	type Embedded struct {
		EmbeddedField string
	}

	tables := []struct {
		title string
		i     interface{}
		flags []cli.Flag
		err   error
	}{
		{
			title: "Can work with types without struct tags",
			i: struct {
				Bool         bool
				Int          int
				Int64        int64
				Float64      float64
				Uint         uint
				Uint64       uint64
				String       string
				Duration     time.Duration
				Timestamp    time.Time
				IntSlice     []int
				Int64Slice   []int64
				Float64Slice []float64
				StringSlice  []string
			}{},
			flags: []cli.Flag{
				&cli.BoolFlag{Name: "bool", EnvVars: []string{"BOOL"}},
				&cli.IntFlag{Name: "int", EnvVars: []string{"INT"}},
				&cli.Int64Flag{Name: "int64", EnvVars: []string{"INT64"}},
				&cli.Float64Flag{Name: "float64", EnvVars: []string{"FLOAT64"}},
				&cli.UintFlag{Name: "uint", EnvVars: []string{"UINT"}},
				&cli.Uint64Flag{Name: "uint64", EnvVars: []string{"UINT64"}},
				&cli.StringFlag{Name: "string", EnvVars: []string{"STRING"}},
				&cli.DurationFlag{Name: "duration", EnvVars: []string{"DURATION"}},
				&cli.TimestampFlag{Name: "timestamp", EnvVars: []string{"TIMESTAMP"}},
				&cli.IntSliceFlag{Name: "intSlice", EnvVars: []string{"INT_SLICE"}},
				&cli.Int64SliceFlag{Name: "int64Slice", EnvVars: []string{"INT64_SLICE"}},
				&cli.Float64SliceFlag{Name: "float64Slice", EnvVars: []string{"FLOAT64_SLICE"}},
				&cli.StringSliceFlag{Name: "stringSlice", EnvVars: []string{"STRING_SLICE"}},
			},
			err: nil,
		},
		{
			title: "Can't define flags from embedded struct",
			i: struct {
				Embedded
				Field string
			}{},
			flags: []cli.Flag{
				&cli.StringFlag{Name: "field", EnvVars: []string{"FIELD"}},
			},
			err: nil,
		},
		{
			title: "Can parse cli.flag.name and cli.flag.envVars",
			i: struct {
				Bool         bool          `cli.flag.name:"flagBool" cli.flag.envVars:"FLAG_BOOL,BOOL"`
				Int          int           `cli.flag.name:"flagInt" cli.flag.envVars:"FLAG_INT,INT"`
				Int64        int64         `cli.flag.name:"flagInt64" cli.flag.envVars:"FLAG_INT64,INT64"`
				Float64      float64       `cli.flag.name:"flagFloat64" cli.flag.envVars:"FLAG_FLOAT64,FLOAT64"`
				Uint         uint          `cli.flag.name:"flagUint" cli.flag.envVars:"FLAG_UINT,UINT"`
				Uint64       uint64        `cli.flag.name:"flagUint64" cli.flag.envVars:"FLAG_UINT64,UINT64"`
				String       string        `cli.flag.name:"flagString" cli.flag.envVars:"FLAG_STRING,STRING"`
				Duration     time.Duration `cli.flag.name:"flagDuration" cli.flag.envVars:"FLAG_DURATION,DURATION"`
				Timestamp    time.Time     `cli.flag.name:"flagTimestamp" cli.flag.envVars:"FLAG_TIMESTAMP,TIMESTAMP"`
				IntSlice     []int         `cli.flag.name:"flagIntSlice" cli.flag.envVars:"FLAG_INT_SLICE,INT_SLICE"`
				Int64Slice   []int64       `cli.flag.name:"flagInt64Slice" cli.flag.envVars:"FLAG_INT64_SLICE,INT64_SLICE"`
				Float64Slice []float64     `cli.flag.name:"flagFloat64Slice" cli.flag.envVars:"FLAG_FLOAT64_SLICE,FLOAT64_SLICE"`
				StringSlice  []string      `cli.flag.name:"flagStringSlice" cli.flag.envVars:"FLAG_STRING_SLICE,STRING_SLICE"`
			}{},
			flags: []cli.Flag{
				&cli.BoolFlag{Name: "flagBool", EnvVars: []string{"FLAG_BOOL", "BOOL"}},
				&cli.IntFlag{Name: "flagInt", EnvVars: []string{"FLAG_INT", "INT"}},
				&cli.Int64Flag{Name: "flagInt64", EnvVars: []string{"FLAG_INT64", "INT64"}},
				&cli.Float64Flag{Name: "flagFloat64", EnvVars: []string{"FLAG_FLOAT64", "FLOAT64"}},
				&cli.UintFlag{Name: "flagUint", EnvVars: []string{"FLAG_UINT", "UINT"}},
				&cli.Uint64Flag{Name: "flagUint64", EnvVars: []string{"FLAG_UINT64", "UINT64"}},
				&cli.StringFlag{Name: "flagString", EnvVars: []string{"FLAG_STRING", "STRING"}},
				&cli.DurationFlag{Name: "flagDuration", EnvVars: []string{"FLAG_DURATION", "DURATION"}},
				&cli.TimestampFlag{Name: "flagTimestamp", EnvVars: []string{"FLAG_TIMESTAMP", "TIMESTAMP"}},
				&cli.IntSliceFlag{Name: "flagIntSlice", EnvVars: []string{"FLAG_INT_SLICE", "INT_SLICE"}},
				&cli.Int64SliceFlag{Name: "flagInt64Slice", EnvVars: []string{"FLAG_INT64_SLICE", "INT64_SLICE"}},
				&cli.Float64SliceFlag{Name: "flagFloat64Slice", EnvVars: []string{"FLAG_FLOAT64_SLICE", "FLOAT64_SLICE"}},
				&cli.StringSliceFlag{Name: "flagStringSlice", EnvVars: []string{"FLAG_STRING_SLICE", "STRING_SLICE"}},
			},
			err: nil,
		},
		{
			title: "Can bypass envVars",
			i: struct {
				Bool         bool          `cli.flag.name:"flagBool" cli.flag.envVars:"-"`
				Int          int           `cli.flag.name:"flagInt" cli.flag.envVars:"-"`
				Int64        int64         `cli.flag.name:"flagInt64" cli.flag.envVars:"-"`
				Float64      float64       `cli.flag.name:"flagFloat64" cli.flag.envVars:"-"`
				Uint         uint          `cli.flag.name:"flagUint" cli.flag.envVars:"-"`
				Uint64       uint64        `cli.flag.name:"flagUint64" cli.flag.envVars:"-"`
				String       string        `cli.flag.name:"flagString" cli.flag.envVars:"-"`
				Duration     time.Duration `cli.flag.name:"flagDuration" cli.flag.envVars:"-"`
				Timestamp    time.Time     `cli.flag.name:"flagTimestamp" cli.flag.envVars:"-"`
				IntSlice     []int         `cli.flag.name:"flagIntSlice" cli.flag.envVars:"-"`
				Int64Slice   []int64       `cli.flag.name:"flagInt64Slice" cli.flag.envVars:"-"`
				Float64Slice []float64     `cli.flag.name:"flagFloat64Slice" cli.flag.envVars:"-"`
				StringSlice  []string      `cli.flag.name:"flagStringSlice" cli.flag.envVars:"-"`
			}{},
			flags: []cli.Flag{
				&cli.BoolFlag{Name: "flagBool"},
				&cli.IntFlag{Name: "flagInt"},
				&cli.Int64Flag{Name: "flagInt64"},
				&cli.Float64Flag{Name: "flagFloat64"},
				&cli.UintFlag{Name: "flagUint"},
				&cli.Uint64Flag{Name: "flagUint64"},
				&cli.StringFlag{Name: "flagString"},
				&cli.DurationFlag{Name: "flagDuration"},
				&cli.TimestampFlag{Name: "flagTimestamp"},
				&cli.IntSliceFlag{Name: "flagIntSlice"},
				&cli.Int64SliceFlag{Name: "flagInt64Slice"},
				&cli.Float64SliceFlag{Name: "flagFloat64Slice"},
				&cli.StringSliceFlag{Name: "flagStringSlice"},
			},
			err: nil,
		},
		{
			title: "Can bypass flag",
			i: struct {
				Flag       string
				BypassFlag string `cli.flag:"-"`
			}{},
			flags: []cli.Flag{
				&cli.StringFlag{Name: "flag", EnvVars: []string{"FLAG"}},
			},
		},
		{
			title: "Can set required flag",
			i: struct {
				Bool         bool          `cli.flag.required:"true"`
				Int          int           `cli.flag.required:"true"`
				Int64        int64         `cli.flag.required:"true"`
				Float64      float64       `cli.flag.required:"true"`
				Uint         uint          `cli.flag.required:"true"`
				Uint64       uint64        `cli.flag.required:"true"`
				String       string        `cli.flag.required:"true"`
				Duration     time.Duration `cli.flag.required:"true"`
				Timestamp    time.Time     `cli.flag.required:"true"`
				IntSlice     []int         `cli.flag.required:"true"`
				Int64Slice   []int64       `cli.flag.required:"true"`
				Float64Slice []float64     `cli.flag.required:"true"`
				StringSlice  []string      `cli.flag.required:"true"`
			}{},
			flags: []cli.Flag{
				&cli.BoolFlag{Name: "bool", EnvVars: []string{"BOOL"}, Required: true},
				&cli.IntFlag{Name: "int", EnvVars: []string{"INT"}, Required: true},
				&cli.Int64Flag{Name: "int64", EnvVars: []string{"INT64"}, Required: true},
				&cli.Float64Flag{Name: "float64", EnvVars: []string{"FLOAT64"}, Required: true},
				&cli.UintFlag{Name: "uint", EnvVars: []string{"UINT"}, Required: true},
				&cli.Uint64Flag{Name: "uint64", EnvVars: []string{"UINT64"}, Required: true},
				&cli.StringFlag{Name: "string", EnvVars: []string{"STRING"}, Required: true},
				&cli.DurationFlag{Name: "duration", EnvVars: []string{"DURATION"}, Required: true},
				&cli.TimestampFlag{Name: "timestamp", EnvVars: []string{"TIMESTAMP"}, Required: true},
				&cli.IntSliceFlag{Name: "intSlice", EnvVars: []string{"INT_SLICE"}, Required: true},
				&cli.Int64SliceFlag{Name: "int64Slice", EnvVars: []string{"INT64_SLICE"}, Required: true},
				&cli.Float64SliceFlag{Name: "float64Slice", EnvVars: []string{"FLOAT64_SLICE"}, Required: true},
				&cli.StringSliceFlag{Name: "stringSlice", EnvVars: []string{"STRING_SLICE"}, Required: true},
			},
			err: nil,
		},
		{
			title: "Can set usage of flag",
			i: struct {
				Bool         bool          `cli.flag.name:"flagBool" cli.flag.usage:"Usage of flag \"flagBool\""`
				Int          int           `cli.flag.name:"flagInt" cli.flag.usage:"Usage of flag \"flagInt\""`
				Int64        int64         `cli.flag.name:"flagInt64" cli.flag.usage:"Usage of flag \"flagInt64\""`
				Float64      float64       `cli.flag.name:"flagFloat64" cli.flag.usage:"Usage of flag \"flagFloat64\""`
				Uint         uint          `cli.flag.name:"flagUint" cli.flag.usage:"Usage of flag \"flagUint\""`
				Uint64       uint64        `cli.flag.name:"flagUint64" cli.flag.usage:"Usage of flag \"flagUint64\""`
				String       string        `cli.flag.name:"flagString" cli.flag.usage:"Usage of flag \"flagString\""`
				Duration     time.Duration `cli.flag.name:"flagDuration" cli.flag.usage:"Usage of flag \"flagDuration\""`
				Timestamp    time.Time     `cli.flag.name:"flagTimestamp" cli.flag.usage:"Usage of flag \"flagTimestamp\""`
				IntSlice     []int         `cli.flag.name:"flagIntSlice" cli.flag.usage:"Usage of flag \"flagIntSlice\""`
				Int64Slice   []int64       `cli.flag.name:"flagInt64Slice" cli.flag.usage:"Usage of flag \"flagInt64Slice\""`
				Float64Slice []float64     `cli.flag.name:"flagFloat64Slice" cli.flag.usage:"Usage of flag \"flagFlat64Slice\""`
				StringSlice  []string      `cli.flag.name:"flagStringSlice" cli.flag.usage:"Usage of flag \"flagStringSlice\""`
			}{},
			flags: []cli.Flag{
				&cli.BoolFlag{Name: "flagBool", EnvVars: []string{"FLAG_BOOL"}, Usage: "Usage of flag \"flagBool\""},
				&cli.IntFlag{Name: "flagInt", EnvVars: []string{"FLAG_INT"}, Usage: "Usage of flag \"flagInt\""},
				&cli.Int64Flag{Name: "flagInt64", EnvVars: []string{"FLAG_INT64"}, Usage: "Usage of flag \"flagInt64\""},
				&cli.Float64Flag{Name: "flagFloat64", EnvVars: []string{"FLAG_FLOAT64"}, Usage: "Usage of flag \"flagFloat64\""},
				&cli.UintFlag{Name: "flagUint", EnvVars: []string{"FLAG_UINT"}, Usage: "Usage of flag \"flagUint\""},
				&cli.Uint64Flag{Name: "flagUint64", EnvVars: []string{"FLAG_UINT64"}, Usage: "Usage of flag \"flagUint64\""},
				&cli.StringFlag{Name: "flagString", EnvVars: []string{"FLAG_STRING"}, Usage: "Usage of flag \"flagString\""},
				&cli.DurationFlag{Name: "flagDuration", EnvVars: []string{"FLAG_DURATION"}, Usage: "Usage of flag \"flagDuration\""},
				&cli.TimestampFlag{Name: "flagTimestamp", EnvVars: []string{"FLAG_TIMESTAMP"}, Usage: "Usage of flag \"flagTimestamp\""},
				&cli.IntSliceFlag{Name: "flagIntSlice", EnvVars: []string{"FLAG_INT_SLICE"}, Usage: "Usage of flag \"flagIntSlice\""},
				&cli.Int64SliceFlag{Name: "flagInt64Slice", EnvVars: []string{"FLAG_INT64_SLICE"}, Usage: "Usage of flag \"flagInt64Slice\""},
				&cli.Float64SliceFlag{Name: "flagFloat64Slice", EnvVars: []string{"FLAG_FLOAT64_SLICE"}, Usage: "Usage of flag \"flagFlat64Slice\""},
				&cli.StringSliceFlag{Name: "flagStringSlice", EnvVars: []string{"FLAG_STRING_SLICE"}, Usage: "Usage of flag \"flagStringSlice\""},
			},
			err: nil,
		},
		{
			title: "Can't generate unsupported type",
			i:    struct{
				Unsupported byte
			}{},
			flags: nil,
			err:   fmt.Errorf(`define field Unsupported error: unsupported flag type uint8`),
		},
		{
			title: "Can't generate default invalid value",
			i: struct{
				Default int `cli.flag.value:"invalid"`
			}{},
			flags: nil,
			err: fmt.Errorf(`define field Default error: strconv.Atoi: parsing "invalid": invalid syntax`),
		},
		{
			title: "Can't generate invalid required",
			i: struct{
				Default int `cli.flag.required:"invalid"`
			}{},
			flags: nil,
			err: fmt.Errorf(`define field Default error: strconv.ParseBool: parsing "invalid": invalid syntax`),
		},
	}

	for _, table := range tables {
		t.Run(table.title, func(t *testing.T) {
			flags, err := GenerateFlags(table.i)
			if !reflect.DeepEqual(table.err, err) {
				t.Fatalf("assert err failed, excepted '%s', actual '%s'", table.err, err)
			}

			if !reflect.DeepEqual(table.flags, flags) {
				t.Fatalf("assert flags failed\n\texcepted: %+v\n\tactual:   %+v", table.flags, flags)
			}
		})
	}
}
