package corecliapp

import (
	"github.com/urfave/cli/v2"
	"reflect"
	"testing"
)

func Test_setFlags(t *testing.T) {
	type Flags struct {
		Debug    bool `cli.flag.name:"debug" cli.flag.envVars:"DEBUG"`
		MaxIndex int  `cli.flag.name:"maxIndex"`
	}
	type ServerFlags struct {
		Flags
		Addr string `cli.flag.name:"addr"`
	}
	app := cli.NewApp()
	app.Name = "testApp"
	app.Flags = []cli.Flag{
		&cli.BoolFlag{Name: "debug", Value: false},
		&cli.IntFlag{Name: "maxIndex", Value: 0},
	}
	app.Commands = cli.Commands{
		&cli.Command{
			Name: "server",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "addr", Value: ":9000"},
			},
			Action: func(c *cli.Context) error {
				flags := &ServerFlags{}
				setFlags(c, reflect.ValueOf(flags))

				if !flags.Debug {
					t.Errorf("assert server command debug flag failed")
				}

				if flags.Addr != ":3000" {
					t.Errorf("assert server command addr flag failed")
				}

				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		flags := &Flags{Debug: false}
		setFlags(c, reflect.ValueOf(flags))

		if !flags.Debug {
			t.Errorf("assert set flags type bool failed")
		}

		if flags.MaxIndex != 10 {
			t.Fatalf("assert set flags type int failed")
		}

		return nil
	}

	if err := app.Run([]string{"testApp", "--debug", "--maxIndex", "10"}); err != nil {
		t.Fatalf("error run test app: %v", err)
	}

	if err := app.Run([]string{"testApp", "--debug", "--maxIndex", "10", "server", "--addr", ":3000"}); err != nil {
		t.Fatalf("error run server app command: %v", err)
	}
}

func Test_callRun(t *testing.T) {

}
