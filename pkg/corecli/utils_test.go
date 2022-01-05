package corecli

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"reflect"
	"runtime"
	"testing"
	"time"
)

func Test_typeToFlag(t *testing.T) {
	tables := []struct {
		title string
		i     interface{}
		flag  cli.Flag
		err   error
	}{
		{
			title: "Can type to bool flag",
			i:     true,
			flag:  &cli.BoolFlag{},
			err:   nil,
		},
		{
			title: "Can type to int flag",
			i:     int(1),
			flag:  &cli.IntFlag{},
			err:   nil,
		},
		{
			title: "Can type to int64 flag",
			i:     int64(1),
			flag:  &cli.Int64Flag{},
			err:   nil,
		},
		{
			title: "Can type to float64 flag",
			i:     float64(1),
			flag:  &cli.Float64Flag{},
			err:   nil,
		},
		{
			title: "Can type to uint flag",
			i:     uint(1),
			flag:  &cli.UintFlag{},
			err:   nil,
		},
		{
			title: "Can type to uint64 flag",
			i:     uint64(1),
			flag:  &cli.Uint64Flag{},
			err:   nil,
		},
		{
			title: "Can type to string flag",
			i:     "lorem ipsum",
			flag:  &cli.StringFlag{},
			err:   nil,
		},
		{
			title: "Can type to duration flag",
			i:     1 * time.Second,
			flag:  &cli.DurationFlag{},
			err:   nil,
		},
		{
			title: "Can type to timestamp flag",
			i:     time.Now().UTC(),
			flag:  &cli.TimestampFlag{},
			err:   nil,
		},
		{
			title: "Can type to []int flag",
			i:     []int{1, 2},
			flag:  &cli.IntSliceFlag{},
			err:   nil,
		},
		{
			title: "Can type to []int64 flag",
			i:     []int64{1, 2},
			flag:  &cli.Int64SliceFlag{},
			err:   nil,
		},
		{
			title: "Can type to []float64 flag",
			i:     []float64{1, 2},
			flag:  &cli.Float64SliceFlag{},
			err:   nil,
		},
		{
			title: "Can type to []string flag",
			i:     []string{"lorem", "ipsum"},
			flag:  &cli.StringSliceFlag{},
			err:   nil,
		},
		{
			title: "Can't type for unsupported type",
			i:     byte('a'),
			flag:  nil,
			err:   fmt.Errorf(`unsupported flag type uint8`),
		},
	}

	for _, table := range tables {
		t.Run(table.title, func(t *testing.T) {
			flag, err := typeToFlag(table.i)
			if fmt.Sprintf("%s", table.err) != fmt.Sprintf("%s", err) {
				t.Fatalf("assert err failed, excepted '%s', actual '%s'", table.err, err)
			}

			if table.flag == nil && flag == nil {
				return
			}

			if reflect.TypeOf(table.flag).Name() != reflect.TypeOf(flag).Name() {
				t.Fatalf("assert flags failed, excepted: %T actual: %T", table.flag, flag)
			}
		})
	}
}

func Test_generateFlagName(t *testing.T) {
	tables := []struct {
		title string
		field reflect.StructField
		name  string
	}{
		{
			title: "Can generate name from specified `cli.flag.name` tag",
			field: reflect.StructField{Name: "String", Tag: reflect.StructTag(`cli.flag.name:"fieldString"`)},
			name:  "fieldString",
		},
		{
			title: "Can generate name from field name",
			field: reflect.StructField{Name: "Int64"},
			name:  "int64",
		},
	}

	for _, table := range tables {
		t.Run(table.title, func(t *testing.T) {
			name := generateFlagName(table.field)
			if table.name != name {
				t.Errorf("assert name failed, excepted %s, actual %s", table.name, name)
			}
		})
	}
}

func Test_generateFlagEnvVars(t *testing.T) {
	tables := []struct {
		title   string
		tag     reflect.StructTag
		name    string
		envVars []string
	}{
		{
			title:   "Can generate envVars from name",
			tag:     reflect.StructTag(``),
			name:    "Int64",
			envVars: []string{"INT64"},
		},
		{
			title:   "Can generate envVars from cli.flag.envVars tag",
			tag:     reflect.StructTag(`cli.flag.envVars:"FLAG_INT64,INT64"`),
			name:    "Int64",
			envVars: []string{"FLAG_INT64", "INT64"},
		},
		{
			title:   "Can't generate envVars when cli.flag.envVars == '-'",
			tag:     reflect.StructTag(`cli.flag.envVars:"-"`),
			name:    "Int64",
			envVars: []string{},
		},
	}

	for _, table := range tables {
		t.Run(table.title, func(t *testing.T) {
			envVars := generateFlagEnvVars(table.tag, table.name)

			if !reflect.DeepEqual(table.envVars, envVars) {
				t.Errorf("assert envVars failed, excepted %#v, actual %#v", table.envVars, envVars)
			}
		})
	}
}

func Test_generateFlagValue(t *testing.T) {
	tables := []struct {
		title  string
		tag    reflect.StructTag
		i      interface{}
		value  interface{}
		errMap map[string]error
	}{
		{
			title: "Can generate bool value",
			tag:   reflect.StructTag(`cli.flag.value:"true"`),
			i:     false,
			value: true,
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can generate bool value from empty cli.flag.value",
			tag:   reflect.StructTag(`cli.flag.value:""`),
			i:     false,
			value: false,
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can't generate invalid bool value",
			tag:   reflect.StructTag(`cli.flag.value:"invalid"`),
			i:     false,
			value: false,
			errMap: map[string]error{
				"": fmt.Errorf(`strconv.ParseBool: parsing "invalid": invalid syntax`),
			},
		},

		{
			title: "Can generate int value",
			tag:   reflect.StructTag(`cli.flag.value:"10"`),
			i:     int(0),
			value: int(10),
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can generate int value from empty cli.flag.value",
			tag:   reflect.StructTag(`cli.flag.value:""`),
			i:     int(0),
			value: int(0),
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can't generate invalid int value",
			tag:   reflect.StructTag(`cli.flag.value:"invalid"`),
			i:     int(0),
			value: int(0),
			errMap: map[string]error{
				"": fmt.Errorf(`strconv.Atoi: parsing "invalid": invalid syntax`),
			},
		},

		{
			title: "Can generate int64 value",
			tag:   reflect.StructTag(`cli.flag.value:"10"`),
			i:     int64(0),
			value: int64(10),
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can generate int64 value from empty cli.flag.value",
			tag:   reflect.StructTag(`cli.flag.value:""`),
			i:     int64(0),
			value: int64(0),
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can't generate invalid int64 value",
			tag:   reflect.StructTag(`cli.flag.value:"invalid"`),
			i:     int64(0),
			value: int64(0),
			errMap: map[string]error{
				"": fmt.Errorf(`strconv.ParseInt: parsing "invalid": invalid syntax`),
			},
		},

		{
			title: "Can generate float64 value",
			tag:   reflect.StructTag(`cli.flag.value:"100.50"`),
			i:     float64(0),
			value: 100.5,
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can generate float64 value from empty cli.flag.value",
			tag:   reflect.StructTag(`cli.flag.value:""`),
			i:     float64(0),
			value: float64(0),
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can't generate invalid float64 value",
			tag:   reflect.StructTag(`cli.flag.value:"invalid"`),
			i:     float64(0),
			value: float64(0),
			errMap: map[string]error{
				"": fmt.Errorf(`strconv.ParseFloat: parsing "invalid": invalid syntax`),
			},
		},

		{
			title: "Can generate uint value",
			tag:   reflect.StructTag(`cli.flag.value:"10"`),
			i:     uint(0),
			value: uint(10),
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can generate uint value from empty cli.flag.value",
			tag:   reflect.StructTag(`cli.flag.value:""`),
			i:     uint(0),
			value: uint(0),
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can't generate invalid uint value",
			tag:   reflect.StructTag(`cli.flag.value:"invalid"`),
			i:     uint(0),
			value: uint(0),
			errMap: map[string]error{
				"": fmt.Errorf(`strconv.ParseUint: parsing "invalid": invalid syntax`),
			},
		},

		{
			title: "Can generate uint64 value",
			tag:   reflect.StructTag(`cli.flag.value:"10"`),
			i:     uint64(0),
			value: uint64(10),
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can generate uint64 value from empty cli.flag.value",
			tag:   reflect.StructTag(`cli.flag.value:""`),
			i:     uint64(0),
			value: uint64(0),
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can't generate invalid uint64 value",
			tag:   reflect.StructTag(`cli.flag.value:"invalid"`),
			i:     uint64(0),
			value: uint64(0),
			errMap: map[string]error{
				"": fmt.Errorf(`strconv.ParseUint: parsing "invalid": invalid syntax`),
			},
		},

		{
			title: "Can generate string value",
			tag:   reflect.StructTag(`cli.flag.value:"string"`),
			i:     "",
			value: "string",
			errMap: map[string]error{
				"": nil,
			},
		},

		{
			title: "Can generate time.Duration value",
			tag:   reflect.StructTag(`cli.flag.value:"10s"`),
			i:     time.Duration(0),
			value: 10 * time.Second,
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can generate time.Duration value from empty cli.flag.value",
			tag:   reflect.StructTag(`cli.flag.value:""`),
			i:     time.Duration(0),
			value: time.Duration(0),
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can't generate invalid time.Duration value",
			tag:   reflect.StructTag(`cli.flag.value:"invalid"`),
			i:     time.Duration(0),
			value: time.Duration(0),
			errMap: map[string]error{
				"go1.14":    fmt.Errorf(`time: invalid duration invalid`),
				"go1.14.15": fmt.Errorf(`time: invalid duration invalid`),
				"":          fmt.Errorf(`time: invalid duration "invalid"`),
			},
		},

		{
			title: "Can generate time.Time value",
			tag:   reflect.StructTag(`cli.flag.value:"2021-01-01T00:00:00.00Z"`),
			i:     time.Now(),
			value: cli.NewTimestamp(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can generate time.Time value from empty cli.flag.value",
			tag:   reflect.StructTag(`cli.flag.value:""`),
			i:     time.Now(),
			value: nil,
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can't generate invalid time.Time value",
			tag:   reflect.StructTag(`cli.flag.value:"invalid"`),
			i:     time.Now(),
			value: nil,
			errMap: map[string]error{
				"": fmt.Errorf("parsing time \"invalid\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"invalid\" as \"2006\""),
			},
		},

		{
			title: "Can generate []int value",
			tag:   reflect.StructTag(`cli.flag.value:"10,20,30, 40, 50"`),
			i:     []int{},
			value: cli.NewIntSlice(10, 20, 30, 40, 50),
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can generate []int value from empty cli.flag.value",
			tag:   reflect.StructTag(`cli.flag.value:""`),
			i:     []int{},
			value: nil,
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can't generate invalid []int value",
			tag:   reflect.StructTag(`cli.flag.value:"abc,10,20"`),
			i:     []int{},
			value: nil,
			errMap: map[string]error{
				"": fmt.Errorf(`[0]int: strconv.Atoi: parsing "abc": invalid syntax`),
			},
		},

		{
			title: "Can generate []int64 value",
			tag:   reflect.StructTag(`cli.flag.value:"10,20,30, 40, 50"`),
			i:     []int64{},
			value: cli.NewInt64Slice(10, 20, 30, 40, 50),
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can generate []int64 value from empty cli.flag.value",
			tag:   reflect.StructTag(`cli.flag.value:""`),
			i:     []int64{},
			value: nil,
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can't generate invalid []int64 value",
			tag:   reflect.StructTag(`cli.flag.value:"abc,10,20"`),
			i:     []int64{},
			value: nil,
			errMap: map[string]error{
				"": fmt.Errorf(`[0]int64: strconv.ParseInt: parsing "abc": invalid syntax`),
			},
		},

		{
			title: "Can generate []float64 value",
			tag:   reflect.StructTag(`cli.flag.value:"10.20,20.43,30, 40, 50"`),
			i:     []float64{},
			value: cli.NewFloat64Slice(10.2, 20.43, 30, 40, 50),
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can generate []float64 value from empty cli.flag.value",
			tag:   reflect.StructTag(`cli.flag.value:""`),
			i:     []float64{},
			value: nil,
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can't generate invalid []float64 value",
			tag:   reflect.StructTag(`cli.flag.value:"abc,10,20"`),
			i:     []float64{},
			value: nil,
			errMap: map[string]error{
				"": fmt.Errorf(`[0]float64: strconv.ParseFloat: parsing "abc": invalid syntax`),
			},
		},

		{
			title: "Can generate []string value",
			tag:   reflect.StructTag(`cli.flag.value:"string1,string2,string3,"`),
			i:     []string{},
			value: cli.NewStringSlice("string1", "string2", "string3"),
			errMap: map[string]error{
				"": nil,
			},
		},
		{
			title: "Can generate []string value from empty cli.flag.value",
			tag:   reflect.StructTag(`cli.flag.value:""`),
			i:     []string{},
			value: nil,
			errMap: map[string]error{
				"": nil,
			},
		},

		{
			title: "Can't generate unsupported value",
			tag:   reflect.StructTag(`cli.flag.value:"val"`),
			i:     byte('a'),
			value: nil,
			errMap: map[string]error{
				"": fmt.Errorf(`unsupported flag type uint8`),
			},
		},
	}

	for _, table := range tables {
		t.Run(table.title, func(t *testing.T) {
			value, err := generateFlagValue(table.tag, table.i)
			exceptedErr := table.errMap[""]
			if versionedErr, ok := table.errMap[runtime.Version()]; ok {
				exceptedErr = versionedErr
			}
			if fmt.Sprintf("%s", exceptedErr) != fmt.Sprintf("%s", err) {
				t.Fatalf("assert err failed, excepted '%s', actual '%s'", exceptedErr, err)
			}

			if !reflect.DeepEqual(table.value, value) {
				t.Fatalf("assert value failed, excepted %#v, actual %#v", table.value, value)
			}
		})
	}
}

func Test_generateFlagRequired(t *testing.T) {
	tables := []struct {
		title    string
		tag      reflect.StructTag
		required bool
		err      error
	}{
		{
			title:    "Can generate required",
			tag:      reflect.StructTag(`cli.flag.required:"true"`),
			required: true,
			err:      nil,
		},
		{
			title:    "Can generate not required",
			tag:      reflect.StructTag(`cli.flag.required:"false"`),
			required: false,
			err:      nil,
		},
		{
			title:    "Can generate not required with empty",
			tag:      reflect.StructTag(`cli.flag.required:""`),
			required: false,
			err:      nil,
		},
		{
			title:    "Can't generate required invalid bool",
			tag:      reflect.StructTag(`cli.flag.required:"invalid"`),
			required: false,
			err:      fmt.Errorf(`strconv.ParseBool: parsing "invalid": invalid syntax`),
		},
	}

	for _, table := range tables {
		t.Run(table.title, func(t *testing.T) {
			required, err := generateFlagRequired(table.tag)
			if fmt.Sprintf("%s", table.err) != fmt.Sprintf("%s", err) {
				t.Fatalf("assert err failed, excepted '%s', actual '%s'", table.err, err)
			}

			if table.required != required {
				t.Fatalf("assert value failed, excepted %#v, actual %#v", table.required, required)
			}
		})
	}
}

func Test_stringToStringSlice(t *testing.T) {
	tables := []struct {
		title string
		s     string
		ss    []string
	}{
		{
			title: "Can string to string slice",
			s:     "abc,zef, qwerty,lorem , ip sum,",
			ss:    []string{"abc", "zef", "qwerty", "lorem", "ip sum"},
		},
	}

	for _, table := range tables {
		t.Run(table.title, func(t *testing.T) {
			ss := stringToStringSlice(table.s)
			if !reflect.DeepEqual(table.ss, ss) {
				t.Errorf("assert equal failed, excepted %#v, actual %#v", table.ss, ss)
			}
		})
	}
}

func Test_stringToStringMap(t *testing.T) {
	tables := []struct {
		title string
		s     string
		sm    map[string]string
	}{
		{
			title: "Can string to string map",
			s:     "lorem:ipsum,dolor:euro, ruble:rub,    ru:ru, en:eng,,",
			sm: map[string]string{
				"lorem": "ipsum",
				"dolor": "euro",
				"ruble": "rub",
				"ru":    "ru",
				"en":    "eng",
			},
		},
	}

	for _, table := range tables {
		t.Run(table.title, func(t *testing.T) {
			sm := stringToStringMap(table.s)
			if !reflect.DeepEqual(table.sm, sm) {
				t.Errorf("assert equal failed, excepted %#v, actual %#v", table.sm, sm)
			}
		})
	}
}
