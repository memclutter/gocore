package corecliapp

import (
	"fmt"
	"github.com/memclutter/gocore/pkg/coreslices"
	"github.com/urfave/cli/v2"
	"reflect"
	"strings"
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
	rAppDefine = valueOfPtr(rAppDefine)
	rtAppDefine := typeOfPtr(rAppDefine)

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
	rAppDefine = valueOfPtr(rAppDefine)
	rtAppDefine := typeOfPtr(rAppDefine)

	// @only struct processing
	if rtAppDefine.Kind() != reflect.Struct {
		return flags, nil
	}

	for i := 0; i < rtAppDefine.NumField(); i++ {
		field := rtAppDefine.Field(i)
		fieldType := field.Type
		fieldTypeKind := fieldType.Kind()

		// Flags struct
		if field.Tag.Get("cli") == "flags" {
			flags, err = lookupFlags(rAppDefine.Field(i), flags)
			if err != nil {
				return nil, err
			}
			continue
		}

		// Current struct flags
		name := strings.TrimSpace(field.Tag.Get("cli.flag.name"))
		if len(name) == 0 {
			continue
		}

		// EnvVars "NAME,VAR_NAME, " -> []string{"NAME", "VAR_NAME"}
		envVars := strings.Split(field.Tag.Get("cli.flag.envVars"), ",")
		envVars = coreslices.StringApply(envVars, func(i int, s string) string { return strings.TrimSpace(s) })
		envVars = coreslices.StringFilter(envVars, func(i int, s string) bool { return len(s) > 0 })

		// Default values
		var flag cli.Flag
		switch fieldTypeKind {
		case reflect.Bool:
			flag = &cli.BoolFlag{
				Name:    name,
				Value:   rAppDefine.Field(i).Bool(),
				EnvVars: envVars,
			}
		case reflect.String:
			flag = &cli.StringFlag{
				Name:    name,
				Value:   rAppDefine.Field(i).String(),
				EnvVars: envVars,
			}
		case reflect.Int:
			flag = &cli.IntFlag{
				Name:    name,
				Value:   int(rAppDefine.Field(i).Int()),
				EnvVars: envVars,
			}
		default:
			return nil, fmt.Errorf("unsupport kind of type %s", fieldTypeKind)
		}

		flags = append(flags, flag)
	}

	return flags, nil
}

// lookupCommands godoc
//
// Lookup and parse app commands from app define struct.
func lookupCommands(rAppDefine reflect.Value) (cli.Commands, error) {
	var err error
	commands := make(cli.Commands, 0)



	return commands, err
}