package corecli

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"reflect"
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
			if fmt.Sprintf("%s",table.err) != fmt.Sprintf("%s", err) {
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

func Test_generateFlagName(t *testing.T)  {
	tables := []struct{
		title string
		field reflect.StructField
		name string
	}{
		{
			title: "Can generate name from specified `cli.flag.name` tag",
			field: reflect.StructField{Name: "String", Tag: reflect.StructTag(`cli.flag.name:"fieldString"`)},
			name: "fieldString",
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
	tables := []struct{
		title string
		tag reflect.StructTag
		name string
		envVars []string
	}{
		{
			title: "Can generate envVars from name",
			tag: reflect.StructTag(``),
			name: "Int64",
			envVars: []string{"INT64"},
		},
		{
			title: "Can generate envVars from cli.flag.envVars tag",
			tag: reflect.StructTag(`cli.flag.envVars:"FLAG_INT64,INT64"`),
			name: "Int64",
			envVars: []string{"FLAG_INT64", "INT64"},
		},
		{
			title: "Can't generate envVars when cli.flag.envVars == '-'",
			tag: reflect.StructTag(`cli.flag.envVars:"-"`),
			name: "Int64",
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
	tables := []struct{
		title string
		tag reflect.StructTag
		i interface{}
		value interface{}
		err error
	}{
		{
			title: "Can generate bool value",
			tag: reflect.StructTag(`cli.flag.value:"true"`),
			i: false,
			value: true,
			err: nil,
		},
		{
			title: "Can generate bool value from empty cli.flag.value",
			tag: reflect.StructTag(`cli.flag.value:""`),
			i: false,
			value: false,
			err: nil,
		},
		{
			title: "Can't generate invalid bool value",
			tag: reflect.StructTag(`cli.flag.value:"invalid"`),
			i: false,
			value: false,
			err: fmt.Errorf("strconv.ParseBool: parsing \"invalid\": invalid syntax"),
		},

		{
			title: "Can generate int value",
			tag: reflect.StructTag(`cli.flag.value:"10"`),
			i: int(0),
			value: int(10),
			err: nil,
		},
		{
			title: "Can generate int value from empty cli.flag.value",
			tag: reflect.StructTag(`cli.flag.value:""`),
			i: int(0),
			value: int(0),
			err: nil,
		},
		{
			title: "Can't generate invalid int value",
			tag: reflect.StructTag(`cli.flag.value:"invalid"`),
			i: int(0),
			value: int(0),
			err: fmt.Errorf("strconv.Atoi: parsing \"invalid\": invalid syntax"),
		},

		{
			title: "Can generate int64 value",
			tag: reflect.StructTag(`cli.flag.value:"10"`),
			i: int64(0),
			value: int64(10),
			err: nil,
		},
		{
			title: "Can generate int64 value from empty cli.flag.value",
			tag: reflect.StructTag(`cli.flag.value:""`),
			i: int64(0),
			value: int64(0),
			err: nil,
		},
		{
			title: "Can't generate invalid int64 value",
			tag: reflect.StructTag(`cli.flag.value:"invalid"`),
			i: int64(0),
			value: int64(0),
			err: fmt.Errorf("strconv.ParseInt: parsing \"invalid\": invalid syntax"),
		},


		{
			title: "Can generate float64 value",
			tag:   reflect.StructTag(`cli.flag.value:"100.50"`),
			i:     float64(0),
			value: 100.5,
			err:   nil,
		},
		{
			title: "Can generate float64 value from empty cli.flag.value",
			tag: reflect.StructTag(`cli.flag.value:""`),
			i: float64(0),
			value: float64(0),
			err: nil,
		},
		{
			title: "Can't generate invalid float64 value",
			tag: reflect.StructTag(`cli.flag.value:"invalid"`),
			i: float64(0),
			value: float64(0),
			err: fmt.Errorf("strconv.ParseFloat: parsing \"invalid\": invalid syntax"),
		},

		{
			title: "Can generate uint value",
			tag: reflect.StructTag(`cli.flag.value:"10"`),
			i: uint(0),
			value: uint(10),
			err: nil,
		},
		{
			title: "Can generate uint value from empty cli.flag.value",
			tag: reflect.StructTag(`cli.flag.value:""`),
			i: uint(0),
			value: uint(0),
			err: nil,
		},
		{
			title: "Can't generate invalid uint value",
			tag: reflect.StructTag(`cli.flag.value:"invalid"`),
			i: uint(0),
			value: uint(0),
			err: fmt.Errorf("strconv.ParseUint: parsing \"invalid\": invalid syntax"),
		},


		{
			title: "Can generate uint64 value",
			tag: reflect.StructTag(`cli.flag.value:"10"`),
			i: uint64(0),
			value: uint64(10),
			err: nil,
		},
		{
			title: "Can generate uint64 value from empty cli.flag.value",
			tag: reflect.StructTag(`cli.flag.value:""`),
			i: uint64(0),
			value: uint64(0),
			err: nil,
		},
		{
			title: "Can't generate invalid uint64 value",
			tag: reflect.StructTag(`cli.flag.value:"invalid"`),
			i: uint64(0),
			value: uint64(0),
			err: fmt.Errorf("strconv.ParseUint: parsing \"invalid\": invalid syntax"),
		},

		{
			title: "Can generate string value",
			tag: reflect.StructTag(`cli.flag.value:"string"`),
			i: "",
			value: "string",
			err: nil,
		},

		{
			title: "Can generate time.Duration value",
			tag: reflect.StructTag(`cli.flag.value:"10s"`),
			i: time.Duration(0),
			value: 10 * time.Second,
			err: nil,
		},
		{
			title: "Can generate time.Duration value from empty cli.flag.value",
			tag: reflect.StructTag(`cli.flag.value:""`),
			i: time.Duration(0),
			value: time.Duration(0),
			err: nil,
		},
		{
			title: "Can't generate invalid time.Duration value",
			tag: reflect.StructTag(`cli.flag.value:"invalid"`),
			i: time.Duration(0),
			value: time.Duration(0),
			err: fmt.Errorf("time: invalid duration \"invalid\""),
		},
	}

	for _, table := range tables {
		t.Run(table.title, func(t *testing.T) {
			value, err := generateFlagValue(table.tag, table.i)
			if fmt.Sprintf("%s",table.err) != fmt.Sprintf("%s", err) {
				t.Fatalf("assert err failed, excepted '%s', actual '%s'", table.err, err)
			}

			if !reflect.DeepEqual(table.value, value) {
				t.Fatalf("assert value failed, excepted %#v, actual %#v", table.value, value)
			}
		})
	}
}