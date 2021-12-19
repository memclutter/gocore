package corecliapp

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"reflect"
	"testing"
	"time"
)

func Test_createFlag(t *testing.T) {
	exampleTimeDate := time.Date(2020, 1,1,0,0,0,0, time.UTC)
	exampleUnsupport := struct{A string}{A: "a"}
	tables := []struct {
		structField reflect.StructField
		value       reflect.Value
		flag        cli.Flag
		err         error
	}{
		{
			structField: reflect.StructField{Name: "Token", Type: reflect.TypeOf("token-default-value")},
			value:       reflect.ValueOf("token-default-value"),
			flag:        &cli.StringFlag{Name: "token", Value: "token-default-value", EnvVars: []string{"TOKEN"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "Timeout", Type: reflect.TypeOf(10 * time.Second)},
			value:       reflect.ValueOf(10 * time.Second),
			flag:        &cli.DurationFlag{Name: "timeout", Value: 10 * time.Second, EnvVars: []string{"TIMEOUT"}},
			err: nil,
		},
		{
			structField: reflect.StructField{Name: "Debug", Type: reflect.TypeOf(true)},
			value:       reflect.ValueOf(true),
			flag:        &cli.BoolFlag{Name: "debug", Value: true, EnvVars: []string{"DEBUG"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "MaxPrice", Type: reflect.TypeOf(10.10)},
			value: reflect.ValueOf(10.10),
			flag: &cli.Float64Flag{Name: "maxPrice", Value: 10.10, EnvVars: []string{"MAX_PRICE"}},
			err: nil,
		},
		{
			structField: reflect.StructField{Name: "MaxDate", Type: reflect.TypeOf(exampleTimeDate)},
			value: reflect.ValueOf(exampleTimeDate),
			flag: &cli.TimestampFlag{Name: "maxDate", Value: cli.NewTimestamp(exampleTimeDate), EnvVars: []string{"MAX_DATE"}},
			err: nil,
		},
		{
			structField: reflect.StructField{Name: "Count", Type: reflect.TypeOf(10)},
			value:       reflect.ValueOf(10),
			flag:        &cli.IntFlag{Name: "count", Value: 10, EnvVars: []string{"COUNT"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "Count", Type: reflect.TypeOf(int64(10))},
			value:       reflect.ValueOf(int64(10)),
			flag:        &cli.Int64Flag{Name: "count", Value: 10, EnvVars: []string{"COUNT"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "Count", Type: reflect.TypeOf(uint(10))},
			value:       reflect.ValueOf(uint(10)),
			flag:        &cli.UintFlag{Name: "count", Value: 10, EnvVars: []string{"COUNT"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "Count", Type: reflect.TypeOf(uint64(10))},
			value:       reflect.ValueOf(uint64(10)),
			flag:        &cli.Uint64Flag{Name: "count", Value: 10, EnvVars: []string{"COUNT"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "Count", Type: reflect.TypeOf(10), Tag: reflect.StructTag(`cli.flag.name:"cnt" cli.flag.envVars:"CNT,COUNT"`)},
			value: reflect.ValueOf(10),
			flag: &cli.IntFlag{Name: "cnt", Value: 10, EnvVars: []string{"CNT", "COUNT"}},
			err: nil,
		},
		{
			structField: reflect.StructField{Name: "Count", Type: reflect.TypeOf(10), Tag: reflect.StructTag(`cli.flag.envVars:"-"`)},
			value: reflect.ValueOf(10),
			flag: &cli.IntFlag{Name: "count", Value: 10, EnvVars: []string{}},
			err: nil,
		},
		{
			structField: reflect.StructField{Name: "Count", Type: reflect.TypeOf(10), Tag: reflect.StructTag(`cli.flag.value:"20"`)},
			value: reflect.ValueOf(10),
			flag: &cli.IntFlag{Name: "count", Value: 20, EnvVars: []string{"COUNT"}},
			err: nil,
		},
		{
			structField: reflect.StructField{Name: "internal", Type: reflect.TypeOf(exampleUnsupport)},
			value: reflect.ValueOf(exampleUnsupport),
			flag: nil,
			err: fmt.Errorf("unsupport flag type '%T' for field 'internal'", exampleUnsupport),
		},
	}

	for _, table := range tables {
		flag, err := createFlag(table.structField, table.value)
		if fmt.Sprintf("%v", table.err) != fmt.Sprintf("%v", err) {
			t.Fatalf("assert error failed, excepted '%s', actual '%s'", table.err, err)
		}

		var exceptedName string
		var exceptedEnvVars []string
		var exceptedValue reflect.Value
		var name string
		var envVars []string
		var value reflect.Value

		switch exceptedFlag := table.flag.(type) {
		case *cli.StringFlag:
			exceptedName = exceptedFlag.Name
			exceptedEnvVars = exceptedFlag.EnvVars
			exceptedValue = reflect.ValueOf(exceptedFlag.Value)
			actualFlag, ok := flag.(*cli.StringFlag)
			if !ok {
				t.Fatalf("assert flag type failed, excepted %T, actual %T", exceptedFlag, actualFlag)
			}
			name = actualFlag.Name
			envVars = actualFlag.EnvVars
			value = reflect.ValueOf(actualFlag.Value)
		case *cli.DurationFlag:
			exceptedName = exceptedFlag.Name
			exceptedEnvVars = exceptedFlag.EnvVars
			exceptedValue = reflect.ValueOf(exceptedFlag.Value)
			actualFlag, ok := flag.(*cli.DurationFlag)
			if !ok {
				t.Fatalf("assert flag type failed, excepted %T, actual %T", exceptedFlag, actualFlag)
			}
			name = actualFlag.Name
			envVars = actualFlag.EnvVars
			value = reflect.ValueOf(actualFlag.Value)
		case *cli.BoolFlag:
			exceptedName = exceptedFlag.Name
			exceptedEnvVars = exceptedFlag.EnvVars
			exceptedValue = reflect.ValueOf(exceptedFlag.Value)
			actualFlag, ok := flag.(*cli.BoolFlag)
			if !ok {
				t.Fatalf("assert flag type failed, excepted %T, actual %T", exceptedFlag, actualFlag)
			}
			name = actualFlag.Name
			envVars = actualFlag.EnvVars
			value = reflect.ValueOf(actualFlag.Value)
		case *cli.Float64Flag:
			exceptedName = exceptedFlag.Name
			exceptedEnvVars = exceptedFlag.EnvVars
			exceptedValue = reflect.ValueOf(exceptedFlag.Value)
			actualFlag, ok := flag.(*cli.Float64Flag)
			if !ok {
				t.Fatalf("assert flag type failed, excepted %T, actual %T", exceptedFlag, actualFlag)
			}
			name = actualFlag.Name
			envVars = actualFlag.EnvVars
			value = reflect.ValueOf(actualFlag.Value)
		case *cli.TimestampFlag:
			exceptedName = exceptedFlag.Name
			exceptedEnvVars = exceptedFlag.EnvVars
			exceptedValue = reflect.ValueOf(exceptedFlag.Value)
			actualFlag, ok := flag.(*cli.TimestampFlag)
			if !ok {
				t.Fatalf("assert flag type failed, excepted %T, actual %T", exceptedFlag, actualFlag)
			}
			name = actualFlag.Name
			envVars = actualFlag.EnvVars
			value = reflect.ValueOf(actualFlag.Value)
		case *cli.IntFlag:
			exceptedName = exceptedFlag.Name
			exceptedEnvVars = exceptedFlag.EnvVars
			exceptedValue = reflect.ValueOf(exceptedFlag.Value)
			actualFlag, ok := flag.(*cli.IntFlag)
			if !ok {
				t.Fatalf("assert flag type failed, excepted %T, actual %T", exceptedFlag, actualFlag)
			}
			name = actualFlag.Name
			envVars = actualFlag.EnvVars
			value = reflect.ValueOf(actualFlag.Value)
		case *cli.Int64Flag:
			exceptedName = exceptedFlag.Name
			exceptedEnvVars = exceptedFlag.EnvVars
			exceptedValue = reflect.ValueOf(exceptedFlag.Value)
			actualFlag, ok := flag.(*cli.Int64Flag)
			if !ok {
				t.Fatalf("assert flag type failed, excepted %T, actual %T", exceptedFlag, actualFlag)
			}
			name = actualFlag.Name
			envVars = actualFlag.EnvVars
			value = reflect.ValueOf(actualFlag.Value)
		case *cli.UintFlag:
			exceptedName = exceptedFlag.Name
			exceptedEnvVars = exceptedFlag.EnvVars
			exceptedValue = reflect.ValueOf(exceptedFlag.Value)
			actualFlag, ok := flag.(*cli.UintFlag)
			if !ok {
				t.Fatalf("assert flag type failed, excepted %T, actual %T", exceptedFlag, actualFlag)
			}
			name = actualFlag.Name
			envVars = actualFlag.EnvVars
			value = reflect.ValueOf(actualFlag.Value)
		case *cli.Uint64Flag:
			exceptedName = exceptedFlag.Name
			exceptedEnvVars = exceptedFlag.EnvVars
			exceptedValue = reflect.ValueOf(exceptedFlag.Value)
			actualFlag, ok := flag.(*cli.Uint64Flag)
			if !ok {
				t.Fatalf("assert flag type failed, excepted %T, actual %T", exceptedFlag, actualFlag)
			}
			name = actualFlag.Name
			envVars = actualFlag.EnvVars
			value = reflect.ValueOf(actualFlag.Value)
		case nil:
			return
		default:
			t.Fatalf("unimplement flag type %T", table.flag)
		}

		if exceptedName != name {
			t.Errorf("assert flag name failed, excepted '%s', actual '%s'", exceptedName, name)
		} else if !reflect.DeepEqual(exceptedEnvVars, envVars) {
			t.Errorf("assert flag envVars failed, excepted %#v, actual %#v", exceptedEnvVars, envVars)
		} else if !reflect.DeepEqual(exceptedValue.Interface(), value.Interface()) {
			t.Errorf("assert flag value failed, excepted %#v, actual %#v", exceptedValue, value)
		}
	}
}

func Test_setFlags(t *testing.T) {
	type Flags struct {
		Debug      bool  `cli.flag.name:"debug" cli.flag.envVars:"DEBUG"`
		MaxIndex   int   `cli.flag.name:"maxIndex"`
		MaxIndex64 int64 `cli.flag.name:"maxIndex64"`
	}
	type ServerFlags struct {
		Flags
		Addr string `cli.flag.name:"addr"`
	}
	app := cli.NewApp()
	app.Name = "testApp"
	app.Flags = []cli.Flag{
		&cli.BoolFlag{Name: "debug", Value: false},
		&cli.IntFlag{Name: "maxIndex", Value: 0},
		&cli.Int64Flag{Name: "maxIndex64", Value: 0},
	}
	app.Commands = cli.Commands{
		&cli.Command{
			Name: "server",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "addr", Value: ":9000"},
			},
			Action: func(c *cli.Context) error {
				flags := &ServerFlags{}
				setFlags(c, reflect.ValueOf(flags))

				if !flags.Debug {
					t.Errorf("assert server command debug flag failed")
				}

				if flags.Addr != ":3000" {
					t.Errorf("assert server command addr flag failed")
				}

				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		flags := &Flags{Debug: false}
		setFlags(c, reflect.ValueOf(flags))

		if !flags.Debug {
			t.Errorf("assert set flags type bool failed")
		}

		if flags.MaxIndex != 10 {
			t.Fatalf("assert set flags type int failed")
		}

		if flags.MaxIndex64 != 20 {
			t.Fatalf("assert set flags type int64 failed")
		}

		return nil
	}

	if err := app.Run([]string{"testApp", "--debug", "--maxIndex", "10", "--maxIndex64", "20"}); err != nil {
		t.Fatalf("error run test app: %v", err)
	}

	if err := app.Run([]string{"testApp", "--debug", "--maxIndex", "10", "server", "--addr", ":3000"}); err != nil {
		t.Fatalf("error run server app command: %v", err)
	}
}

type TestCallRunCommand struct {
}

func (cmd TestCallRunCommand) Run() error {
	return nil
}

func Test_callRun(t *testing.T) {
	err := callRun(reflect.ValueOf(TestCallRunCommand{}))
	if err != nil {
		t.Fatalf("error call run, excepted no errors, actual have error: %v", err)
	}
}

type TestCallRunCommandWithError struct {
}

func (cmd TestCallRunCommandWithError) Run() error {
	return fmt.Errorf("test")
}

func Test_callRunWithError(t *testing.T) {
	exceptedErr := fmt.Errorf("test")
	err := callRun(reflect.ValueOf(TestCallRunCommandWithError{}))
	if err.Error() != exceptedErr.Error() {
		t.Fatalf("error call run, excepted have error '%s', actual '%s'", exceptedErr, err)
	}
}
