package corecliapp

import (
	"github.com/memclutter/gocore/pkg/corestrings"
	"github.com/urfave/cli/v2"
	"reflect"
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
	app.Flags = ConfigToFlags(config)
	return app
}

type CoreCommand interface {
	Run(c *cli.Context) error
	Flags() []cli.Flag
}

func (app *App) CoreCommands(commands []CoreCommand) *App {
	app.Commands = cli.Commands{}
	for _, c := range commands {
		name := corestrings.ToLowerFirst(reflect.ValueOf(c).Type().Name())

		app.Commands = append(app.Commands, &cli.Command{
			Name:   name,
			Flags:  c.Flags(),
			Action: c.Run,
		})
	}
	return app
}
