package corecliapp

import (
	"github.com/urfave/cli/v2"
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
