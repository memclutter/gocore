package corecli

import (
	"github.com/go-pg/pg/v10"
	"github.com/urfave/cli/v2"
	"reflect"
	"testing"
)

func TestLoadDependencies(t *testing.T) {
	//tables := []struct{
	//	title string
	//	i interface{}
	//
	//}{
	//
	//}

	type Command struct {
		DB *pg.DB `cli.command.dependency:"dsn:dsnDb"`
	}

	app := cli.NewApp()
	app.Name = "gocore"
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "dsnDb"},
	}
	app.Action = func(c *cli.Context) error {

		cmd := &Command{}

		if err := LoadDependencies(reflect.ValueOf(cmd), c); err != nil {

		}

		return nil
	}

	args := []string{
		"gocore",
		"--dsnDb", "postgres://gocore:gocore@localhost:5432/gocore?sslmode=disable",
	}
	if err := app.Run(args); err != nil {
		t.Fatalf("assert app run failed %v", err)
		return
	}
}
