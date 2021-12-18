package corecliapp

import (
	"github.com/urfave/cli/v2"
	"reflect"
	"testing"
)

func Test_create(t *testing.T) {
	type Flags struct {
		Debug bool `cli.flag.name:"debug" cli.flag.envVars:"DEBUG"`
	}
	type App struct {
		Name  string `cli:"name"`
		Flags *Flags `cli:"flags"`
	}

	appDefine := &App{
		Name: "test",
		Flags: &Flags{
			Debug: false,
		},
	}

	app, err := create(appDefine)
	if err != nil {
		t.Fatalf("app create assert error, excepted no errors, actual: %v", err)
	}

	if app.Name != appDefine.Name {
		t.Errorf("app name assert error, excepted '%s', actual '%s'", appDefine.Name, app.Name)
	}

	if len(app.Flags) != 1 {
		t.Fatalf("app flags count assert error, excepted %d, actual %d", 1, len(app.Flags))
	}

	debugBoolFlag, ok := app.Flags[0].(*cli.BoolFlag)
	if !ok {
		t.Fatalf("app flag debug type assert error, excepted bool, actual %T", app.Flags[0])
	}

	if debugBoolFlag.Name != "debug" {
		t.Fatalf("app flag debug name assert error, excepted 'debug', actual '%s'", debugBoolFlag.Name)
	}

	if debugBoolFlag.Value != appDefine.Flags.Debug {
		t.Fatalf("app flag debug default value assert error, excepted %v, actual %v", appDefine.Flags.Debug, debugBoolFlag.Value)
	}
}

func Test_lookupName(t *testing.T) {
	tables := []struct {
		rAppDefine reflect.Value
		name       string
	}{
		{
			rAppDefine: reflect.ValueOf(struct {
				Name string `cli:"name"`
			}{Name: "test"}),
			name: "test",
		},
	}

	for _, table := range tables {
		name := lookupName(table.rAppDefine)
		if name != table.name {
			t.Errorf("assert name failed, excepted '%s', actual '%s'", table.name, name)
		}
	}
}

func Test_lookupFlags(t *testing.T) {
	tables := []struct {
		rAppDefine reflect.Value
		flags      []cli.Flag
		err        error
	}{
		{
			rAppDefine: reflect.ValueOf(struct {
				Flags struct {
					Token string `cli.flag.name:"token"`
					Debug bool   `cli.flag.name:"debug"`
				} `cli:"flags"`
			}{
				Flags: struct {
					Token string `cli.flag.name:"token"`
					Debug bool   `cli.flag.name:"debug"`
				}{Token: "Test", Debug: true},
			}),
			flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "token",
					Value: "Test",
				},
				&cli.BoolFlag{
					Name:  "debug",
					Value: true,
				},
			},
			err: nil,
		},
	}

	for _, table := range tables {
		flags, err := lookupFlags(table.rAppDefine, []cli.Flag{})
		if err != table.err {
			t.Errorf("assert err of flags failed, excepted %v, actual %v", table.err, err)
		}

		if len(flags) != len(table.flags) {
			t.Errorf("assert len of flags failed, excepted %d, actual %d", len(table.flags), len(flags))
		}

		for i, flag := range flags {
			switch actualFlag := flag.(type) {
			case *cli.StringFlag:
				exceptedFlag, ok := table.flags[i].(*cli.StringFlag)
				if !ok {
					t.Errorf("assert type of flag %d failed, excepted %T, actual %T", i, table.flags[i], actualFlag)
				}
				if exceptedFlag.Name != actualFlag.Name {
					t.Errorf("assert name of flag %d failed, excepted '%s', actual '%s'", i, exceptedFlag.Name, actualFlag.Name)
				}
				if exceptedFlag.Value != actualFlag.Value {
					t.Errorf("assert value of flag %d failed, excepted %v', actual %v", i, exceptedFlag.Value, actualFlag.Value)
				}
			case *cli.BoolFlag:
				exceptedFlag, ok := table.flags[i].(*cli.BoolFlag)
				if !ok {
					t.Errorf("assert type of flag %d failed, excepted %T, actual %T", i, table.flags[i], actualFlag)
				}
				if exceptedFlag.Name != actualFlag.Name {
					t.Errorf("assert name of flag %d failed, excepted '%s', actual '%s'", i, exceptedFlag.Name, actualFlag.Name)
				}
				if exceptedFlag.Value != actualFlag.Value {
					t.Errorf("assert value of flag %d failed, excepted %v', actual %v", i, exceptedFlag.Value, actualFlag.Value)
				}
			default:
				t.Errorf("assert type of flag %d failed, excepted %T, actual %T", i, table.flags[i], actualFlag)
			}
		}
	}
}

type TestApp struct {
	Name     string        `cli:"name"`
	Flags    *TestAppFlags `cli:"flags"`
	Commands []Command     `cli:"commands"`
}

type TestAppFlags struct {
	Debug bool `cli.flag.name:"debug" cli.flag.envVars:"DEBUG"`
}

type TestAppCommand struct {
	Name     string    `cli.command:"name"`
	Commands []Command `cli.command:"commands"`
}

type TestAppCommandFlags struct {
	Addr string `cli.flag.name:"addr" cli.flag.envVars:"ADDR"`
}

func (cmd TestAppCommand) Run() error {
	return nil
}

type TestAppSubCommand struct {
	Name string `cli.command:"name"`
}

func (cmd TestAppSubCommand) Run() error {
	return nil
}

func Test_lookupCommands(t *testing.T) {
	appDefine := TestApp{
		Name: "testApp",
		Commands: []Command{
			&TestAppCommand{
				Name: "testCommand",
				Commands: []Command{
					&TestAppSubCommand{
						Name: "testSubCommand",
					},
				},
			},
		},
	}

	commands, err := lookupCommands(reflect.ValueOf(appDefine))
	if err != nil {
		t.Fatalf("error lookup commands: %v", err)
	}

	if len(commands) != len(appDefine.Commands) {
		t.Fatalf("assert commands len failed, excepted %d, actual %d", len(appDefine.Commands), len(commands))
	}

	exceptedCommand := appDefine.Commands[0].(*TestAppCommand)
	if commands[0].Name != exceptedCommand.Name {
		t.Fatalf("assert command name failed, excepted '%s', actual '%s'", exceptedCommand.Name, commands[0].Name)
	}

	if len(commands[0].Subcommands) != len(exceptedCommand.Commands) {
		t.Fatalf("assert subcommand len failed, excepted %d, actual %d", len(exceptedCommand.Commands), len(commands[0].Subcommands))
	}

	exceptedSubCommand := exceptedCommand.Commands[0].(*TestAppSubCommand)
	if commands[0].Subcommands[0].Name != exceptedSubCommand.Name {
		t.Fatalf("assert subcommand name failed, excepted '%s', actual '%s'", exceptedSubCommand.Name, commands[0].Subcommands[0].Name)
	}
}
