package corecliapp

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"reflect"
	"runtime"
	"testing"
	"time"
)

func Test_createFlag(t *testing.T) {
	exampleTimeDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	exampleUnsupported := struct{ A string }{A: "a"}
	tables := []struct {
		structField reflect.StructField
		value       reflect.Value
		flag        cli.Flag
		err         error
		errVersions map[string]error
	}{
		{
			structField: reflect.StructField{Name: "Token", Type: reflect.TypeOf("token-default-value")},
			value:       reflect.ValueOf("token-default-value"),
			flag:        &cli.StringFlag{Name: "token", Value: "token-default-value", EnvVars: []string{"TOKEN"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "Token", Type: reflect.TypeOf("token-default-value"), Tag: reflect.StructTag(`cli.flag.value:"override"`)},
			value:       reflect.ValueOf("token-default-value"),
			flag:        &cli.StringFlag{Name: "token", Value: "override", EnvVars: []string{"TOKEN"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "Timeout", Type: reflect.TypeOf(10 * time.Second)},
			value:       reflect.ValueOf(10 * time.Second),
			flag:        &cli.DurationFlag{Name: "timeout", Value: 10 * time.Second, EnvVars: []string{"TIMEOUT"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "Timeout", Type: reflect.TypeOf(10 * time.Second), Tag: reflect.StructTag(`cli.flag.value:"30s"`)},
			value:       reflect.ValueOf(10 * time.Second),
			flag:        &cli.DurationFlag{Name: "timeout", Value: 30 * time.Second, EnvVars: []string{"TIMEOUT"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "Timeout", Type: reflect.TypeOf(10 * time.Second), Tag: reflect.StructTag(`cli.flag.value:"invalid30s"`)},
			value:       reflect.ValueOf(10 * time.Second),
			flag:        nil,
			err:         fmt.Errorf(`error parse duration 'invalid30s': time: invalid duration "invalid30s"`),
			errVersions: map[string]error{
				"go1.14": fmt.Errorf(`error parse duration 'invalid30s': time: invalid duration invalid30s`),
			},
		},
		{
			structField: reflect.StructField{Name: "Debug", Type: reflect.TypeOf(true)},
			value:       reflect.ValueOf(true),
			flag:        &cli.BoolFlag{Name: "debug", Value: true, EnvVars: []string{"DEBUG"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "Debug", Type: reflect.TypeOf(true), Tag: reflect.StructTag(`cli.flag.value:"false"`)},
			value:       reflect.ValueOf(true),
			flag:        &cli.BoolFlag{Name: "debug", Value: false, EnvVars: []string{"DEBUG"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "Debug", Type: reflect.TypeOf(true), Tag: reflect.StructTag(`cli.flag.value:"invalid"`)},
			value:       reflect.ValueOf(true),
			flag:        nil,
			err:         fmt.Errorf(`error parse bool 'invalid': strconv.ParseBool: parsing "invalid": invalid syntax`),
		},
		{
			structField: reflect.StructField{Name: "MaxPrice", Type: reflect.TypeOf(10.1)},
			value:       reflect.ValueOf(10.1),
			flag:        &cli.Float64Flag{Name: "maxPrice", Value: 10.10, EnvVars: []string{"MAX_PRICE"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "MaxPrice", Type: reflect.TypeOf(10.1), Tag: reflect.StructTag(`cli.flag.value:"30.1"`)},
			value:       reflect.ValueOf(10.1),
			flag:        &cli.Float64Flag{Name: "maxPrice", Value: 30.1, EnvVars: []string{"MAX_PRICE"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "MaxPrice", Type: reflect.TypeOf(10.1), Tag: reflect.StructTag(`cli.flag.value:"invalid"`)},
			value:       reflect.ValueOf(10.1),
			flag:        nil,
			err:         fmt.Errorf(`error parse float 'invalid': strconv.ParseFloat: parsing "invalid": invalid syntax`),
		},
		{
			structField: reflect.StructField{Name: "MaxDate", Type: reflect.TypeOf(exampleTimeDate)},
			value:       reflect.ValueOf(exampleTimeDate),
			flag:        &cli.TimestampFlag{Name: "maxDate", Value: cli.NewTimestamp(exampleTimeDate), EnvVars: []string{"MAX_DATE"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "MaxDate", Type: reflect.TypeOf(exampleTimeDate), Tag: reflect.StructTag(`cli.flag.value:"2021-02-01T00:00:10Z"`)},
			value:       reflect.ValueOf(exampleTimeDate),
			flag:        &cli.TimestampFlag{Name: "maxDate", Value: cli.NewTimestamp(time.Date(2021, 2, 1, 0, 0, 10, 0, time.UTC)), EnvVars: []string{"MAX_DATE"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "MaxDate", Type: reflect.TypeOf(exampleTimeDate), Tag: reflect.StructTag(`cli.flag.value:"invalid"`)},
			value:       reflect.ValueOf(exampleTimeDate),
			flag:        nil,
			err:         fmt.Errorf(`error parse time.Time 'invalid' (layout '2006-01-02T15:04:05Z07:00'): parsing time "invalid" as "2006-01-02T15:04:05Z07:00": cannot parse "invalid" as "2006"`),
		},
		{
			structField: reflect.StructField{Name: "Count", Type: reflect.TypeOf(10)},
			value:       reflect.ValueOf(10),
			flag:        &cli.IntFlag{Name: "count", Value: 10, EnvVars: []string{"COUNT"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "Count", Type: reflect.TypeOf(10), Tag: reflect.StructTag(`cli.flag.value:"20"`)},
			value:       reflect.ValueOf(10),
			flag:        &cli.IntFlag{Name: "count", Value: 20, EnvVars: []string{"COUNT"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "Count", Type: reflect.TypeOf(10), Tag: reflect.StructTag(`cli.flag.value:"invalid"`)},
			value:       reflect.ValueOf(10),
			flag:        nil,
			err:         fmt.Errorf(`error parse int 'invalid': strconv.ParseInt: parsing "invalid": invalid syntax`),
		},
		{
			structField: reflect.StructField{Name: "Count", Type: reflect.TypeOf(int64(10))},
			value:       reflect.ValueOf(int64(10)),
			flag:        &cli.Int64Flag{Name: "count", Value: 10, EnvVars: []string{"COUNT"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "Count", Type: reflect.TypeOf(int64(10)), Tag: reflect.StructTag(`cli.flag.value:"40"`)},
			value:       reflect.ValueOf(int64(10)),
			flag:        &cli.Int64Flag{Name: "count", Value: 40, EnvVars: []string{"COUNT"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "Count", Type: reflect.TypeOf(int64(10)), Tag: reflect.StructTag(`cli.flag.value:"invalid"`)},
			value:       reflect.ValueOf(int64(10)),
			flag:        nil,
			err:         fmt.Errorf(`error parse int64 'invalid': strconv.ParseInt: parsing "invalid": invalid syntax`),
		},
		{
			structField: reflect.StructField{Name: "Count", Type: reflect.TypeOf(uint(10))},
			value:       reflect.ValueOf(uint(10)),
			flag:        &cli.UintFlag{Name: "count", Value: 10, EnvVars: []string{"COUNT"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "Count", Type: reflect.TypeOf(uint(10)), Tag: reflect.StructTag(`cli.flag.value:"40"`)},
			value:       reflect.ValueOf(uint(10)),
			flag:        &cli.UintFlag{Name: "count", Value: 40, EnvVars: []string{"COUNT"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "Count", Type: reflect.TypeOf(uint(10)), Tag: reflect.StructTag(`cli.flag.value:"invalid"`)},
			value:       reflect.ValueOf(uint(10)),
			flag:        nil,
			err:         fmt.Errorf(`error parse uint 'invalid': strconv.ParseUint: parsing "invalid": invalid syntax`),
		},
		{
			structField: reflect.StructField{Name: "Count", Type: reflect.TypeOf(uint64(10))},
			value:       reflect.ValueOf(uint64(10)),
			flag:        &cli.Uint64Flag{Name: "count", Value: 10, EnvVars: []string{"COUNT"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "Count", Type: reflect.TypeOf(uint64(10)), Tag: reflect.StructTag(`cli.flag.value:"90"`)},
			value:       reflect.ValueOf(uint64(10)),
			flag:        &cli.Uint64Flag{Name: "count", Value: 90, EnvVars: []string{"COUNT"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "Count", Type: reflect.TypeOf(uint64(10)), Tag: reflect.StructTag(`cli.flag.value:"invalid"`)},
			value:       reflect.ValueOf(uint64(10)),
			flag:        nil,
			err:         fmt.Errorf(`error parse uint64 'invalid': strconv.ParseUint: parsing "invalid": invalid syntax`),
		},
		{
			structField: reflect.StructField{Name: "Count", Type: reflect.TypeOf(10), Tag: reflect.StructTag(`cli.flag.name:"cnt" cli.flag.envVars:"CNT,COUNT"`)},
			value:       reflect.ValueOf(10),
			flag:        &cli.IntFlag{Name: "cnt", Value: 10, EnvVars: []string{"CNT", "COUNT"}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "Count", Type: reflect.TypeOf(10), Tag: reflect.StructTag(`cli.flag.envVars:"-"`)},
			value:       reflect.ValueOf(10),
			flag:        &cli.IntFlag{Name: "count", Value: 10, EnvVars: []string{}},
			err:         nil,
		},
		{
			structField: reflect.StructField{Name: "internal", Type: reflect.TypeOf(exampleUnsupported)},
			value:       reflect.ValueOf(exampleUnsupported),
			flag:        nil,
			err:         fmt.Errorf("unsupport flag type '%T' for field 'internal'", exampleUnsupported),
		},
	}

	for _, table := range tables {
		flag, err := createFlag(table.structField, table.value)
		exceptedErr := table.err
		if versionedErr, ok := table.errVersions[runtime.Version()]; ok {
			exceptedErr = versionedErr
		}
		if fmt.Sprintf("%v", exceptedErr) != fmt.Sprintf("%v", err) {
			t.Errorf("assert error failed, excepted '%s', actual '%s'", exceptedErr, err)
		}

		if !reflect.DeepEqual(table.flag, flag) {
			t.Errorf("assert flag failed, excepted %#v, actual %#v", table.flag, flag)
		}
	}
}

func Test_createFlags(t *testing.T) {
	type exampleEmbedded struct {
		Token string
		Debug bool
	}
	tables := []struct {
		i     interface{}
		flags []cli.Flag
		err   error
	}{
		{
			i: struct {
				Token   string
				Debug   bool
				Count   int
				Timeout time.Duration
			}{
				Token:   "default-token",
				Debug:   true,
				Count:   200,
				Timeout: 10 * time.Second,
			},
			flags: []cli.Flag{
				&cli.StringFlag{Name: "token", Value: "default-token", EnvVars: []string{"TOKEN"}},
				&cli.BoolFlag{Name: "debug", Value: true, EnvVars: []string{"DEBUG"}},
				&cli.IntFlag{Name: "count", Value: 200, EnvVars: []string{"COUNT"}},
				&cli.DurationFlag{Name: "timeout", Value: 10 * time.Second, EnvVars: []string{"TIMEOUT"}},
			},
			err: nil,
		},
		{
			i: struct {
				exampleEmbedded
				Count   int
				Timeout time.Duration
			}{
				exampleEmbedded: exampleEmbedded{
					Token: "default-token",
					Debug: true,
				},
				Count:   200,
				Timeout: 10 * time.Second,
			},
			flags: []cli.Flag{
				&cli.IntFlag{Name: "count", Value: 200, EnvVars: []string{"COUNT"}},
				&cli.DurationFlag{Name: "timeout", Value: 10 * time.Second, EnvVars: []string{"TIMEOUT"}},
			},
			err: nil,
		},
		{
			i:     []string{"abc"},
			flags: []cli.Flag{},
			err:   nil,
		},
		{
			i: struct {
				Count   int `cli.flag.value:"abc"`
				Timeout time.Duration
			}{
				Count:   200,
				Timeout: 10 * time.Second,
			},
			flags: []cli.Flag{},
			err:   fmt.Errorf(`error parse int 'abc': strconv.ParseInt: parsing "abc": invalid syntax`),
		},
	}

	for _, table := range tables {
		flags, err := createFlags(table.i, []cli.Flag{})
		if fmt.Sprintf("%v", table.err) != fmt.Sprintf("%v", err) {
			t.Errorf("assert error failed, excepted '%s', actual '%s'", table.err, err)
		}

		if !reflect.DeepEqual(table.flags, flags) {
			t.Errorf("assert flags failed, excepted %#v, actual %#v", table.flags, flags)
		}
	}
}

func Test_setFlags(t *testing.T) {
	exampleTime := time.Now().UTC()

	tables := []struct {
		runArgs     []string
		runErr      error
		appName     string
		appFlags    []cli.Flag
		appCommands cli.Commands
	}{
		{
			runArgs: []string{
				"testApp",
				"--token", "app-token",
				"--debug",
				"testCommand",
				"--addr", ":8000",
				"--maxPrice", "2000.5",
				"--maxDate", "2021-02-02T00:00:00Z",
				"--retries", "32",
				"--maxId", "1000",
				"--count", "500",
				"--countReverse", "400",
				"--timeout", "30s",
			},
			runErr:  nil,
			appName: "testApp",
			appFlags: []cli.Flag{
				&cli.BoolFlag{Name: "debug", Value: false, EnvVars: []string{"DEBUG"}},
				&cli.StringFlag{Name: "token", Value: "default-token", EnvVars: []string{"TOKEN"}},
			},
			appCommands: cli.Commands{
				&cli.Command{
					Name: "testCommand",
					Flags: []cli.Flag{
						&cli.StringFlag{Name: "addr", Value: ":9000", EnvVars: []string{"ADDR"}},
						&cli.Float64Flag{Name: "maxPrice", Value: 0},
						&cli.TimestampFlag{Name: "maxDate", Value: cli.NewTimestamp(exampleTime), Layout: time.RFC3339},
						&cli.IntFlag{Name: "retries", Value: 10},
						&cli.Int64Flag{Name: "maxId", Value: 0},
						&cli.UintFlag{Name: "count", Value: 10},
						&cli.Uint64Flag{Name: "countReverse", Value: 100},
						&cli.DurationFlag{Name: "timeout", Value: 10*time.Second},
					},
					Action: func(c *cli.Context) error {
						type Flags struct {
							Debug    bool
							ApiToken string `cli.flag.name:"token"`
						}
						type CommandFlags struct {
							Flags
							Addr         string
							MaxPrice     float64 `cli.flag.name:"maxPrice"`
							MaxDate      time.Time
							Retries      int
							MaxID        int64 `cli.flag.name:"maxId"`
							Count        uint
							CountReverse uint64
							Timeout time.Duration
						}
						flags := &CommandFlags{}
						if err := setFlags(c, reflect.ValueOf(flags)); err != nil {
							return err
						}

						if flags.Debug != c.Bool("debug") {
							t.Errorf("assert root bool flag --debug failed, excepted %v, actual %v", c.Bool("debug"), flags.Debug)
						}

						if flags.ApiToken != c.String("token") {
							t.Errorf("assert root string flag --token failed, excepted %v, actual %v", c.String("token"), flags.ApiToken)
						}

						if flags.Addr != c.String("addr") {
							t.Errorf("assert string flag --addr failed, excepted %v, actual %v", c.String("addr"), flags.Addr)
						}

						if flags.MaxPrice != c.Float64("maxPrice") {
							t.Errorf("assert float64 flag --maxPrice failed, excepted %v, actual %v", c.Float64("maxPrice"), flags.MaxPrice)
						}

						if !flags.MaxDate.Equal(*c.Timestamp("maxDate")) {
							t.Errorf("assert timestamp flag --maxDate failed, excepted %v, actual %v", c.Timestamp("maxDate"), flags.MaxDate)
						}

						if flags.Retries != c.Int("retries") {
							t.Errorf("assert int flag --retries failed, excepted %v, actual %v", c.Int("retries"), flags.Retries)
						}

						if flags.MaxID != c.Int64("maxId") {
							t.Errorf("assert int64 flag --maxId failed, excepted %v, actual %v", c.Int("maxId"), flags.MaxID)
						}

						if flags.Count != c.Uint("count") {
							t.Errorf("assert uint flag --count failed, excepted %v, actual %v", c.Int("count"), flags.Count)
						}

						if flags.CountReverse != c.Uint64("countReverse") {
							t.Errorf("assert uint64 flag --countReverse failed, excepted %v, actual %v", c.Int("countReverse"), flags.CountReverse)
						}

						if flags.Timeout != c.Duration("timeout") {
							t.Errorf("assert duration flag --timeout failed, excepted %v, actual %v", c.Int("timeout"), flags.Timeout)
						}

						return nil
					},
				},
			},
		},
	}

	for _, table := range tables {
		app := cli.NewApp()
		app.Name = table.appName
		app.Flags = table.appFlags
		app.Commands = table.appCommands

		if err := app.Run(table.runArgs); fmt.Sprintf("%s", err) != fmt.Sprintf("%s", table.runErr) {
			t.Errorf("assert run err failed, excepted '%s', actual '%s'", table.runErr, err)
		}
	}
}
