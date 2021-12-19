package corecliapp

import (
	"fmt"
	"github.com/memclutter/gocore/pkg/corereflect"
	"github.com/urfave/cli/v2"
	"reflect"
)

// create godoc
//
// Create app. Use app define for create urfave/cli/v2 app.
func create(appDefine interface{}) (*cli.App, error) {
	var err error

	// Prepare reflection type of app define
	rAppDefine := reflect.ValueOf(appDefine)

	app := cli.NewApp()
	app.Name = lookupName(rAppDefine)
	app.Flags, err = lookupFlags(rAppDefine, app.Flags)
	if err != nil {
		return nil, fmt.Errorf("error lookup flags: %v", err)
	}
	app.Commands, err = lookupCommands(rAppDefine)
	if err != nil {
		return nil, fmt.Errorf("error lookup commands: %v", err)
	}

	return app, err
}

// lookupName godoc
//
// Lookup and parse app name from app define struct. App name contains in struct field with tag `cli:"name"`
func lookupName(rAppDefine reflect.Value) string {
	rAppDefine = corereflect.PtrValueOf(rAppDefine)
	rtAppDefine := corereflect.PtrTypeOf(rAppDefine)

	// @only struct processing
	if rtAppDefine.Kind() != reflect.Struct {
		return ""
	}

	for i := 0; i < rtAppDefine.NumField(); i++ {
		field := rtAppDefine.Field(i)
		if field.Type.Kind() != reflect.String {
			continue
		}
		if field.Tag.Get("cli") != "name" {
			continue
		}
		return rAppDefine.Field(i).String()
	}

	return ""
}

// lookupFlags godoc
//
// Lookup and parse app flags from app define struct.
func lookupFlags(rAppDefine reflect.Value, flags []cli.Flag) ([]cli.Flag, error) {
	var err error
	rAppDefine = corereflect.PtrValueOf(rAppDefine)
	rtAppDefine := corereflect.PtrTypeOf(rAppDefine)

	// @only struct processing
	if rtAppDefine.Kind() != reflect.Struct {
		return flags, nil
	}

	for i := 0; i < rtAppDefine.NumField(); i++ {
		field := rtAppDefine.Field(i)

		// Flags struct
		if field.Tag.Get("cli") == "flags" {
			flags, err = createFlags(rAppDefine.Field(i).Interface(), flags)
			if err != nil {
				return nil, err
			}
			continue
		}
	}

	return flags, nil
}

// lookupCommands godoc
//
// Lookup and parse app commands from app define struct.
func lookupCommands(rAppDefine reflect.Value) (cli.Commands, error) {
	var err error
	commands := cli.Commands{}
	rAppDefine = corereflect.PtrValueOf(rAppDefine)
	rtAppDefine := corereflect.PtrTypeOf(rAppDefine)

	// Root commands lookup, search `cli:"commands"` struct tag
	if rtAppDefine.Kind() == reflect.Struct {
		for i := 0; i < rtAppDefine.NumField(); i++ {
			field := rtAppDefine.Field(i)

			if field.Tag.Get("cli") == "commands" {
				commands, err = lookupCommands(rAppDefine.Field(i))
				if err != nil {
					return nil, fmt.Errorf("error lookup root commands: %v", err)
				}
				break
			}
		}
	}

	// Iterate over slice of Commands
	if rtAppDefine.Kind() == reflect.Slice {
		for i := 0; i < rAppDefine.Len(); i++ {
			// @TODO check Command interface
			rCommand := corereflect.PtrValueOf(reflect.ValueOf(rAppDefine.Index(i).Interface()))
			rtCommand := corereflect.PtrTypeOf(rCommand)
			command := &cli.Command{}
			flagsIndex := -1
			for j := 0; j < rtCommand.NumField(); j++ {
				field := rtCommand.Field(j)
				if field.Tag.Get("cli.command") == "name" {
					command.Name = rCommand.Field(j).String()
				} else if field.Tag.Get("cli.command") == "commands" {
					command.Subcommands, err = lookupCommands(rCommand.Field(j))
					if err != nil {
						return nil, fmt.Errorf("error lookup subcommands: %v", err)
					}
				} else if field.Tag.Get("cli.command") == "flags" {
					flagsIndex = j
					command.Flags, err = createFlags(rCommand.Field(j).Interface(), []cli.Flag{})
					if err != nil {
						return nil, fmt.Errorf("error lookup command flags: %v", err)
					}
				}
			}

			// Register command function
			command.Action = func(c *cli.Context) error {

				// Preset app flags
				rFlags := corereflect.PtrValueOf(rCommand.Field(flagsIndex))
				if err := setFlags(c, rFlags); err != nil {
					return fmt.Errorf("set flags error: %v", err)
				}

				// Preset app services
				if err := setServices(c, rCommand); err != nil {
					return fmt.Errorf("error set service: %v", err)
				}

				// Call run method
				return callRun(rCommand)
			}

			commands = append(commands, command)
		}
	}

	return commands, err
}
