package corecliapp

import (
	"github.com/urfave/cli/v2"
	"testing"
)

func TestNewApp(t *testing.T) {
	tables := []struct {
		config      interface{}
		cliAppFlags []cli.Flag
	}{
		{
			config: struct {
				ApiKey string `value:"default-api-key"`
				Debug  bool   `value:"0"`
			}{},
			cliAppFlags: []cli.Flag{
				&cli.StringFlag{
					Name:    "apiKey",
					Value:   "default-api-key",
					EnvVars: []string{"API_KEY"},
				},
				&cli.BoolFlag{
					Name:    "debug",
					Value:   false,
					EnvVars: []string{"DEBUG"},
				},
			},
		},
	}

	for _, table := range tables {
		app := NewApp().CoreConfig(table.config)

		if len(table.cliAppFlags) != len(app.Flags) {
			t.Fatalf("excepted app flags and parsed app flags mismatch %d != %d", len(table.cliAppFlags), len(app.Flags))
		}

		for i, exceptedCliFlag := range table.cliAppFlags {
			cliFlag := app.Flags[i]
			switch exceptedCliFlag := exceptedCliFlag.(type) {
			case *cli.StringFlag:
				stringCliFlag, ok := cliFlag.(*cli.StringFlag)
				if !ok {
					t.Errorf("excepted %d flag as type string", i)
				}
				if stringCliFlag.Name != exceptedCliFlag.Name {
					t.Errorf("excepted %d flag name is '%s, actual '%s'", i, exceptedCliFlag.Name, stringCliFlag.Name)
				}
				if stringCliFlag.Value != exceptedCliFlag.Value {
					t.Errorf("excepted %d flag default value is '%s', actual '%s'", i, exceptedCliFlag.Value, stringCliFlag.Value)
				}
				// @TODO slice compare methods need
				if len(stringCliFlag.EnvVars) != len(exceptedCliFlag.EnvVars) {
					t.Errorf("excepted %d flag contain correct count of env vars, excepted %v, actual %v", i, exceptedCliFlag.EnvVars, stringCliFlag.EnvVars)
				}
				for j, exceptedEnvVar := range exceptedCliFlag.EnvVars {
					if stringCliFlag.EnvVars[j] != exceptedEnvVar {
						t.Errorf("%d flag, %d env, excepted '%s', actual '%s'", i, j, exceptedEnvVar, stringCliFlag.EnvVars[j])
					}
				}
			case *cli.BoolFlag:
				actualCliFlag, ok := cliFlag.(*cli.BoolFlag)
				if !ok {
					t.Errorf("excepted %d flag as type bool", i)
				}
				if actualCliFlag.Name != exceptedCliFlag.Name {
					t.Errorf("excepted %d flag name is '%s, actual '%s'", i, exceptedCliFlag.Name, actualCliFlag.Name)
				}
				if actualCliFlag.Value != exceptedCliFlag.Value {
					t.Errorf("excepted %d flag default value is '%T', actual '%T'", i, exceptedCliFlag.Value, actualCliFlag.Value)
				}
				// @TODO slice compare methods need
				if len(actualCliFlag.EnvVars) != len(exceptedCliFlag.EnvVars) {
					t.Errorf("excepted %d flag contain correct count of env vars, excepted %v, actual %v", i, exceptedCliFlag.EnvVars, actualCliFlag.EnvVars)
				}
				for j, exceptedEnvVar := range exceptedCliFlag.EnvVars {
					if actualCliFlag.EnvVars[j] != exceptedEnvVar {
						t.Errorf("%d flag, %d env, excepted '%s', actual '%s'", i, j, exceptedEnvVar, actualCliFlag.EnvVars[j])
					}
				}
			case *cli.IntFlag:
				actualCliFlag, ok := cliFlag.(*cli.IntFlag)
				if !ok {
					t.Errorf("excepted %d flag as type int", i)
				}
				if actualCliFlag.Name != exceptedCliFlag.Name {
					t.Errorf("excepted %d flag name is '%s, actual '%s'", i, exceptedCliFlag.Name, actualCliFlag.Name)
				}
				if actualCliFlag.Value != exceptedCliFlag.Value {
					t.Errorf("excepted %d flag default value is '%T', actual '%T'", i, exceptedCliFlag.Value, actualCliFlag.Value)
				}
				// @TODO slice compare methods need
				if len(actualCliFlag.EnvVars) != len(exceptedCliFlag.EnvVars) {
					t.Errorf("excepted %d flag contain correct count of env vars, excepted %v, actual %v", i, exceptedCliFlag.EnvVars, actualCliFlag.EnvVars)
				}
				for j, exceptedEnvVar := range exceptedCliFlag.EnvVars {
					if actualCliFlag.EnvVars[j] != exceptedEnvVar {
						t.Errorf("%d flag, %d env, excepted '%s', actual '%s'", i, j, exceptedEnvVar, actualCliFlag.EnvVars[j])
					}
				}
			case *cli.DurationFlag:
				actualCliFlag, ok := cliFlag.(*cli.DurationFlag)
				if !ok {
					t.Errorf("excepted %d flag as type duration", i)
				}
				if actualCliFlag.Name != exceptedCliFlag.Name {
					t.Errorf("excepted %d flag name is '%s, actual '%s'", i, exceptedCliFlag.Name, actualCliFlag.Name)
				}
				if actualCliFlag.Value != exceptedCliFlag.Value {
					t.Errorf("excepted %d flag default value is '%T', actual '%T'", i, exceptedCliFlag.Value, actualCliFlag.Value)
				}
				// @TODO slice compare methods need
				if len(actualCliFlag.EnvVars) != len(exceptedCliFlag.EnvVars) {
					t.Errorf("excepted %d flag contain correct count of env vars, excepted %v, actual %v", i, exceptedCliFlag.EnvVars, actualCliFlag.EnvVars)
				}
				for j, exceptedEnvVar := range exceptedCliFlag.EnvVars {
					if actualCliFlag.EnvVars[j] != exceptedEnvVar {
						t.Errorf("%d flag, %d env, excepted '%s', actual '%s'", i, j, exceptedEnvVar, actualCliFlag.EnvVars[j])
					}
				}
			}
		}
	}

}
