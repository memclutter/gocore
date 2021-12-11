package corecliapp

import (
	"github.com/memclutter/gocore/pkg/coreslices"
	"github.com/memclutter/gocore/pkg/corestrings"
	"github.com/urfave/cli/v2"
	"log"
	"reflect"
	"strconv"
	"strings"
)

// App godoc
//
// urfave/cli/v2 application wrapper.
type App struct {
	*cli.App
}

// NewApp godoc
//
// Create new instance of app
func NewApp() *App {
	return &App{
		cli.NewApp(),
	}
}

// CoreConfig godoc
//
// Parse and set up application config
func (app *App) CoreConfig(config interface{}) *App {

	refConfig := reflect.ValueOf(config)
	refConfigType := refConfig.Type()

	for i := 0; i < refConfigType.NumField(); i++ {
		field := refConfigType.Field(i)
		name := strings.TrimSpace(field.Tag.Get("name"))
		value := field.Tag.Get("value")
		envVars := strings.Split(field.Tag.Get("envVars"), ",")
		envVars = coreslices.StringApply(envVars, func(i int, s string) string { return strings.ToLower(s) })
		envVars = coreslices.StringFilter(envVars, func(i int, s string) bool { return len(s) > 0 })

		// Set name as struct name lowerCamelCase
		if len(name) == 0 {
			name = corestrings.ToLowerFirst(field.Name)
		}

		// Set env vars as struct field name converted from CamelCase -> snake_case -> UPPER_SNAKE_CASE
		if len(envVars) == 0 {
			envVars = []string{
				strings.ToUpper(corestrings.CamelToSnake(field.Name)),
			}
		}

		switch v := refConfig.Field(i).Interface().(type) {
		case bool:
			defaultValue, err := strconv.ParseBool(value)
			if err != nil {
				log.Fatalf("invalid boolean value '%s' for config param '%s'", value, field.Name)
			}
			app.Flags = append(app.Flags, &cli.BoolFlag{
				Name:    name,
				Value:   defaultValue,
				EnvVars: envVars,
			})
		case int:
			defaultValue, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				log.Fatalf("invalid int value '%s' for config param '%s'", value, field.Name)
			}
			app.Flags = append(app.Flags, &cli.IntFlag{
				Name:    name,
				Value:   int(defaultValue),
				EnvVars: envVars,
			})
		case string:
			app.Flags = append(app.Flags, &cli.StringFlag{
				Name:    name,
				Value:   value,
				EnvVars: envVars,
			})
		default:
			log.Fatalf("unsuported config type %T for config param '%s'", v, field.Name)
		}
	}

	return app
}
