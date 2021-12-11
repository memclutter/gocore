package corecliapp

import (
	"github.com/urfave/cli/v2"
	"testing"
)

type testConfig struct {
	ApiKey string `value:"default-api-key"`
	Debug  bool   `value:"false"`
}
type testServer struct {
	config testConfig
}
func (srv *testServer) Run(c *cli.Context) error {
	if err := ContextToConfig(c, &srv.config); err != nil {
		return err
	}

	return nil
}

func TestNewApp(t *testing.T) {
	app := NewApp().
		CoreConfig(testConfig{}).
		CoreCommands([]CoreCommand{
			&testServer{},
	})

	if err := app.Run([]string{"gocore", "testServer"}); err != nil {
		t.Errorf("error run test server: %v", err)
	}
}
