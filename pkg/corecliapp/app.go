package corecliapp

import "fmt"

type Command interface {
	Run() error
}

// Run godoc
//
// Create and run app. Use app define for create urfave/cli/v2 app and run it.
func Run(appDefine interface{}, arguments []string) error {
	app, err := create(appDefine)
	if err != nil {
		return fmt.Errorf("error create urfave/cli/v2 app: %v", err)
	}
	return app.Run(arguments)
}
