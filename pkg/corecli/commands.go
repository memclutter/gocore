package corecli

import (
	"fmt"
	"github.com/memclutter/gocore/pkg/corestrings"
	"github.com/urfave/cli/v2"
	"reflect"
	"strings"
	"time"
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

		// el - Slice element, user defined command
		el := valueOf.Index(i).Interface()
		elValueOf := reflect.ValueOf(el)
		elTypeOf := reflect.Indirect(elValueOf).Type()

		command, flagsFieldValueOf, period, err := initCommand(elTypeOf, elValueOf)
		if err != nil {
			return nil, fmt.Errorf("error init command: %v", err)
		}

		elInitable, isImplementInit := el.(Initable)
		elRunnable, isImplementRun := el.(Runnable)
		elClearable, isImplementClear := el.(Clearable)

		command.Before = func(c *cli.Context) error {
			// Default initialization
			if !flagsFieldValueOf.IsZero() && !flagsFieldValueOf.IsNil() && flagsFieldValueOf.IsValid() {
				if err := LoadFlags(flagsFieldValueOf, c); err != nil {
					return err
				}
			}

			if err := LoadDependencies(elValueOf, c); err != nil {
				return err
			}

			if isImplementInit {
				return elInitable.Init()
			}
			return nil
		}

		command.Action = func(c *cli.Context) error {
			if isImplementRun {
				if period == 0 {
					return elRunnable.Run()
				}

				// Periodical command
				for {
					if err := elRunnable.Run(); err != nil {
						return err
					}
					time.Sleep(period)
				}
			}
			return nil
		}

		command.After = func(c *cli.Context) error {
			if isImplementClear {
				return elClearable.Clear()
			}

			if err := CloseDependencies(elValueOf); err != nil {
				return err
			}

			return nil
		}

		commands = append(commands, command)
	}

	return commands, err
}

func initCommand(typeOf reflect.Type, valueOf reflect.Value) (*cli.Command, reflect.Value, time.Duration, error) {
	var flagsFieldValueOf reflect.Value
	var period time.Duration
	var err error

	// Init command and set default name
	command := &cli.Command{
		Name: corestrings.ToLowerFirst(typeOf.Name()),
	}

	for j := 0; j < typeOf.NumField(); j++ {
		field := typeOf.Field(j)
		fieldValueOf := reflect.Indirect(valueOf).Field(j)

		// Command name
		name := strings.TrimSpace(field.Tag.Get("cli.command.name"))
		if len(name) > 0 {
			command.Name = name
			fieldValueOf.SetString(name)
		}

		// Command usage
		usage := strings.TrimSpace(field.Tag.Get("cli.command.usage"))
		if len(usage) > 0 {
			command.Usage = usage
			fieldValueOf.SetString(usage)
		}

		// Command flags
		flags := strings.TrimSpace(field.Tag.Get("cli.command.flags"))
		if flags == "*" {
			flagsFieldValueOf = fieldValueOf
			command.Flags, err = GenerateFlags(fieldValueOf.Interface())
			if err != nil {
				return nil, flagsFieldValueOf, period, fmt.Errorf("generate flags error: %v", err)
			}
		}

		// Periodical command
		if periodTag, ok := field.Tag.Lookup("cli.command.period"); ok {
			period, err = time.ParseDuration(periodTag)
			if err != nil {
				return nil, flagsFieldValueOf, period, fmt.Errorf("parse period error: %v", err)
			}
		}

		// Command subcommands
		subcommands := strings.TrimSpace(field.Tag.Get("cli.command.subcommands"))
		if subcommands == "*" {
			command.Subcommands, err = GenerateCommands(fieldValueOf.Interface())
			if err != nil {
				return nil, flagsFieldValueOf, period, fmt.Errorf("generate subcommands error: %v", err)
			}
		}
	}

	return command, flagsFieldValueOf, period, nil
}
