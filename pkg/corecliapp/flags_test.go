package corecliapp

import (
	"github.com/memclutter/gocore/pkg/corestrings"
	"github.com/urfave/cli/v2"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestConfigToFlags(t *testing.T) {
	tables := []struct {
		config interface{}
		flags  []cli.Flag
	}{
		{
			config: struct {
				ApiKey  string        `value:"default-api-key"`
				Debug   bool          `value:"0"`
				Timeout time.Duration `value:"10s"`
			}{},
			flags: []cli.Flag{
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
				&cli.DurationFlag{
					Name:    "timeout",
					Value:   10 * time.Second,
					EnvVars: []string{"TIMEOUT"},
				},
			},
		},
	}

	for _, table := range tables {
		flags := ConfigToFlags(table.config)

		if len(table.flags) != len(flags) {
			t.Fatalf("excepted app flags and parsed app flags mismatch %d != %d", len(table.flags), len(flags))
		}

		for i, exceptedCliFlag := range table.flags {
			cliFlag := flags[i]
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

func TestContextToConfig(t *testing.T) {
	tables := []struct {
		cliFlags []cli.Flag
		envs     map[string]string
		config   struct {
			ApiKey  string        `value:"default-api-key"`
			Debug   bool          `value:"false"`
			Timeout time.Duration `value:"10s"`
		}
	}{
		{
			cliFlags: []cli.Flag{
				&cli.StringFlag{
					Name:    "apiKey",
					Value:   "default-api-key",
					EnvVars: []string{"API_KEY"},
				},
				&cli.BoolFlag{
					Name:    "debug",
					Value:   true,
					EnvVars: []string{"DEBUG"},
				},
				&cli.DurationFlag{
					Name:    "timeout",
					Value:   10 * time.Second,
					EnvVars: []string{"TIMEOUT"},
				},
			},
			envs: map[string]string{
				"API_KEY": "api-key",
				"DEBUG":   "1",
				"TIMEOUT": "30m",
			},
			config: struct {
				ApiKey  string        `value:"default-api-key"`
				Debug   bool          `value:"false"`
				Timeout time.Duration `value:"10s"`
			}{
				ApiKey:  "api-key",
				Debug:   true,
				Timeout: 30 * time.Minute,
			},
		},
	}

	for _, table := range tables {
		app := cli.NewApp()
		app.Flags = table.cliFlags
		app.Action = func(c *cli.Context) error {
			exceptedConfig := table.config
			if err := ContextToConfig(c, &table.config); err != nil {
				t.Fatalf("error parse app context: %v", err)
			}

			refConfig := reflect.ValueOf(exceptedConfig)
			refConfigType := refConfig.Type()

			for i := 0; i < refConfigType.NumField(); i++ {
				field := refConfigType.Field(i)
				fieldValue := refConfig.Field(i)
				name := strings.TrimSpace(field.Tag.Get("name"))

				// Set name as struct name lowerCamelCase
				if len(name) == 0 {
					name = corestrings.ToLowerFirst(field.Name)
				}

				switch v := fieldValue.Interface().(type) {
				case bool:
					if c.Bool(name) != fieldValue.Bool() {
						t.Errorf("field %s not equal %v != %v", name, c.Bool(name), fieldValue.Bool())
					}
				case int:
					if c.Int(name) != int(fieldValue.Int()) {
						t.Errorf("field %s not equal %v != %v", name, c.Int(name), fieldValue.Int())
					}
				case string:
					if c.String(name) != fieldValue.String() {
						t.Errorf("field %s not equal %v != %v", name, c.String(name), fieldValue.String())
					}
				case time.Duration:
					if c.Duration(name) != fieldValue.Interface().(time.Duration) {
						t.Errorf("field %s not equal %v != %v", name, c.Duration(name), fieldValue.Interface().(time.Duration))
					}
				default:
					t.Fatalf("unsuported config type %T for config param '%s'", v, field.Name)
				}
			}

			return nil
		}

		for name, value := range table.envs {
			if err := os.Setenv(name, value); err != nil {
				t.Fatalf("error set test env %s=%s: %v", name, value, err)
			}
		}

		if err := app.Run([]string{"gocore"}); err != nil {
			t.Fatalf("error run test app: %v", err)
		}

		for name := range table.envs {
			if err := os.Unsetenv(name); err != nil {
				t.Fatalf("error unset test env %s: %v", name, err)
			}
		}
	}
}
