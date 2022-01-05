package corecli

import (
	"fmt"
	"github.com/memclutter/gocore/pkg/corestrings"
	"github.com/urfave/cli/v2"
	"reflect"
	"strings"
)

type Runnable interface {
	Run() error
}

type Initable interface {
	Init() error
}

type Clearable interface {
	Clear() error
}

// GenerateCommands godoc
//
// Define urfave/cli Command from some golang struct. Support tags
// - `cli.command.name` name of command
// - `cli.command.usage` usage of command
// - `cli.command.flags` flags for command
// - `cli.command.subcommands` subcommands
func GenerateCommands(i interface{}) (cli.Commands, error) {
	var err error
	commands := make(cli.Commands, 0)
	valueOf := reflect.ValueOf(i)

	for i := 0; i < valueOf.Len(); i++ {
		command := &cli.Command{}
		el := valueOf.Index(i).Interface()
		elValueOf := reflect.ValueOf(el)
		elTypeOf := reflect.Indirect(elValueOf).Type()
		elInitable, isImplementInit := el.(Initable)
		elRunnable, isImplementRun := el.(Runnable)
		elClearable, isImplementClear := el.(Clearable)

		// Name by default, name of struct
		command.Name = corestrings.ToLowerFirst(elTypeOf.Name())

		var elFlags reflect.Value
		for j := 0; j < elTypeOf.NumField(); j++ {
			elField := elTypeOf.Field(j)
			elFieldValue := reflect.Indirect(elValueOf).Field(j)
			name := strings.TrimSpace(elField.Tag.Get("cli.command.name"))
			if len(name) > 0 {
				command.Name = name
				elFieldValue.SetString(name)
			}
			usage := strings.TrimSpace(elField.Tag.Get("cli.command.usage"))
			if len(usage) > 0 {
				command.Usage = usage
				elFieldValue.SetString(usage)
			}
			flags := strings.TrimSpace(elField.Tag.Get("cli.command.flags"))
			if flags == "*" {
				elFlags = elFieldValue
				command.Flags, err = GenerateFlags(elFieldValue.Interface())
				if err != nil {
					return nil, fmt.Errorf("generate command flags error: %v", err)
				}
			}
			subcommands := strings.TrimSpace(elField.Tag.Get("cli.command.subcommands"))
			if subcommands == "*" {
				command.Subcommands, err = GenerateCommands(elFieldValue.Interface())
				if err != nil {
					return nil, fmt.Errorf("generate subcommands error: %v", err)
				}
			}
		}

		command.Before = func(c *cli.Context) error {
			// Default initialization
			if err := LoadFlags(elFlags, c); err != nil {
				return err
			} else if err := LoadDependencies(elValueOf, c); err != nil {
				return err
			}

			if isImplementInit {
				return elInitable.Init()
			}
			return nil
		}

		command.Action = func(c *cli.Context) error {
			if isImplementRun {
				return elRunnable.Run()
			}
			return nil
		}

		command.After = func(c *cli.Context) error {
			if isImplementClear {
				return elClearable.Clear()
			}
			return nil
		}

		commands = append(commands, command)
	}

	return commands, err
}
