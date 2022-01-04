package corecli

import (
	"github.com/urfave/cli/v2"
	"reflect"
	"testing"
)


type GenerateCommands_CommandFlags struct {
	Bool bool
	String string
}

type GenerateCommands_Command struct {
	Name string `cli.command.name:"command"`
	Usage string `cli.command.usage:"Usage of command"`
	Flags GenerateCommands_CommandFlags `cli.command.flags:"*"`
}

func (cmd GenerateCommands_Command) Init()error  {

	return nil
}

func (cmd GenerateCommands_Command) Run() error {

	return nil
}

//func (cmd GenerateCommands_Command) Clear() error {
//	return nil
//}

func TestGenerateCommands(t *testing.T) {
	tables := []struct{
		title string
		i []interface{}
		commands cli.Commands
		err error
	}{
		{
			title: "Can generate commands 1 levels",
			i: []interface{}{
				&GenerateCommands_Command{},
			},
			commands: cli.Commands{
				&cli.Command{
					Name: "command",
					Usage: "Usage of command",
					Flags: []cli.Flag{
						&cli.BoolFlag{Name: "bool", EnvVars: []string{"BOOL"}},
						&cli.StringFlag{Name: "string", EnvVars: []string{"STRING"}},
					},
					Action: func(c *cli.Context) error {

						return nil
					},
				},
			},
			err: nil,
		},
	}


	for _, table := range tables {
		t.Run(table.title, func(t *testing.T) {
			commands, err := GenerateCommands(table.i)
			if !reflect.DeepEqual(table.err, err) {
				t.Fatalf("assert err failed, excepted '%s', actual '%s'", table.err, err)
			}

			if !reflect.DeepEqual(table.commands, commands) {
				t.Fatalf("assert commands failed\n\texcepted: %+v\n\tactual:   %+v", table.commands, commands)
			}
		})
	}
}
