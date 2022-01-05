package corecli

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/streadway/amqp"
	"github.com/urfave/cli/v2"
	"reflect"
	"testing"
)

const (
	dsnAmqp  = "amqp://gocore:gocore@127.0.0.1:5672/gocore"
	dsnCache = "redis://127.0.0.1:6379/1"
	dsnDb    = "postgres://gocore:gocore@127.0.0.1:5432/gocore?sslmode=disable"

	dsnAmqpInvalid  = "invalid+amqp://gocore:gocore@127.0.0.1:5672/gocore"
	dsnCacheInvalid = "invalid+redis://127.0.0.1:6379/1"
	dsnDbInvalid    = "invalid+postgres://gocore:gocore@127.0.0.1:5432/gocore?sslmode=disable"

	dsnAmqpConnectError  = "amqp://gocore:gocore@127.0.0.1:2765/gocore"
	dsnCacheConnectError = "redis://127.0.0.1:9736/1"
	dsnDbConnectError    = "postgres://gocore:gocore@127.0.0.1:2345/gocore?sslmode=disable"
)

func TestLoadDependencies(t *testing.T) {
	tables := []struct {
		title string
		i     interface{}
		flags []cli.Flag
		args  []string
		err   error
	}{
		{
			title: "Can load dependency",
			i: &struct {
				DB   *pg.DB `cli.command.dependency:"dsn:dsnDb"`
				Pass *pg.DB
			}{},
			flags: []cli.Flag{
				&cli.StringFlag{Name: "dsnDb"},
			},
			args: []string{
				"gocore",
				"--dsnDb", dsnDb,
			},
			err: nil,
		},
		{
			title: "Can't load dependency, because unsupported pgkName",
			i: &struct {
				DB   reflect.Kind `cli.command.dependency:"dsn:dsnDb"`
			}{},
			flags: []cli.Flag{
				&cli.StringFlag{Name: "dsnDb"},
			},
			args: []string{
				"gocore",
				"--dsnDb", dsnDb,
			},
			err: fmt.Errorf(`unsupported depdendency package reflect::Kind`),
		},
		{
			title: "Can't load dependency, because error in loader",
			i: &struct {
				DB   *pg.DB `cli.command.dependency:"dsn:dsnDb"`
			}{},
			flags: []cli.Flag{
				&cli.StringFlag{Name: "dsnDb"},
			},
			args: []string{
				"gocore",
				"--dsnDb", dsnDbInvalid,
			},
			err: fmt.Errorf(`github.com/go-pg/pg/v10::DB: error parse data source name: pg: invalid scheme: invalid+postgres`),
		},
	}

	for _, table := range tables {
		t.Run(table.title, func(t *testing.T) {
			app := cli.NewApp()
			app.Name = table.args[0]
			app.Flags = table.flags
			app.Action = func(c *cli.Context) error {

				err := LoadDependencies(reflect.ValueOf(table.i), c)
				if !reflect.DeepEqual(table.err, err) {
					t.Errorf("assert err failed, excepted '%s', actual '%s'", table.err, err)
				}

				return nil
			}

			if err := app.Run(table.args); err != nil {
				t.Fatalf("assert app run failed: %v", err)
			}
		})
	}
}

func Test_loadDependencyGoPgV10(t *testing.T) {
	tables := []struct {
		title   string
		v       reflect.Value
		options map[string]string
		app     *cli.App
		args    []string
		isNil   bool
		err     error
	}{
		{
			title:   "Can load go-pg v10 dependency correctly",
			v:       reflect.ValueOf(struct{}{}),
			options: map[string]string{"dsn": "dsnDb"},
			app: &cli.App{
				Name: "gocore",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "dsnDb"},
				},
			},
			args: []string{
				"gocore",
				"--dsnDb", dsnDb,
			},
			isNil: false,
			err:   nil,
		},
		{
			title:   "Can't load go-pg v10 dependency correctly, because invalid data source name",
			v:       reflect.ValueOf(struct{}{}),
			options: map[string]string{"dsn": "dsnDb"},
			app: &cli.App{
				Name: "gocore",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "dsnDb"},
				},
			},
			args: []string{
				"gocore",
				"--dsnDb", dsnDbInvalid,
			},
			isNil: true,
			err:   fmt.Errorf(`error parse data source name: pg: invalid scheme: invalid+postgres`),
		},
		{
			title:   "Can load and ping go-pg v10 dependency correctly",
			v:       reflect.ValueOf(struct{}{}),
			options: map[string]string{"dsn": "dsnDb", "ping": "true"},
			app: &cli.App{
				Name: "gocore",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "dsnDb"},
				},
			},
			args: []string{
				"gocore",
				"--dsnDb", dsnDb,
			},
			isNil: false,
			err:   nil,
		},
		{
			title:   "Can't load and ping go-pg v10 dependency correctly",
			v:       reflect.ValueOf(struct{}{}),
			options: map[string]string{"dsn": "dsnDb", "ping": "true"},
			app: &cli.App{
				Name: "gocore",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "dsnDb"},
				},
			},
			args: []string{
				"gocore",
				"--dsnDb", dsnDbConnectError,
			},
			isNil: true,
			err:   fmt.Errorf(`error connect to database: dial tcp 127.0.0.1:2345: connect: connection refused`),
		},
	}

	for _, table := range tables {
		t.Run(table.title, func(t *testing.T) {
			table.app.Action = func(c *cli.Context) error {
				dep, err := loadDependencyGoPgV10(table.v, table.options, c)
				if !reflect.DeepEqual(table.err, err) {
					t.Errorf("assert err failed, excepted '%s', actual '%s'", table.err, err)
				}

				if table.isNil && dep != nil {
					t.Errorf("assert is nil failed, excepted return nil, but %T received", dep)
				}

				return nil
			}

			if err := table.app.Run(table.args); err != nil {
				t.Fatalf("assert run app failed, error returned %v", err)
				return
			}
		})
	}
}

func Test_loadDependencyGoRedisV8(t *testing.T) {
	tables := []struct {
		title   string
		v       reflect.Value
		options map[string]string
		app     *cli.App
		args    []string
		isNil   bool
		err     error
	}{
		{
			title:   "Can load go-redis v8 dependency correctly",
			v:       reflect.ValueOf(struct{}{}),
			options: map[string]string{"dsn": "dsnCache"},
			app: &cli.App{
				Name: "gocore",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "dsnCache"},
				},
			},
			args: []string{
				"gocore",
				"--dsnCache", dsnCache,
			},
			isNil: false,
			err:   nil,
		},
		{
			title:   "Can't load go-redis v8 dependency correctly, because invalid data source name",
			v:       reflect.ValueOf(struct{}{}),
			options: map[string]string{"dsn": "dsnCache"},
			app: &cli.App{
				Name: "gocore",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "dsnCache"},
				},
			},
			args: []string{
				"gocore",
				"--dsnCache", dsnCacheInvalid,
			},
			isNil: true,
			err:   fmt.Errorf(`error parse data source name: redis: invalid URL scheme: invalid+redis`),
		},
		{
			title:   "Can load and ping go-redis v8 dependency correctly",
			v:       reflect.ValueOf(struct{}{}),
			options: map[string]string{"dsn": "dsnCache", "ping": "true"},
			app: &cli.App{
				Name: "gocore",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "dsnCache"},
				},
			},
			args: []string{
				"gocore",
				"--dsnCache", dsnCache,
			},
			isNil: false,
			err:   nil,
		},
		{
			title:   "Can't load and ping go-redis v8 dependency correctly",
			v:       reflect.ValueOf(struct{}{}),
			options: map[string]string{"dsn": "dsnCache", "ping": "true"},
			app: &cli.App{
				Name: "gocore",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "dsnCache"},
				},
			},
			args: []string{
				"gocore",
				"--dsnCache", dsnCacheConnectError,
			},
			isNil: true,
			err:   fmt.Errorf(`error connect to redis: dial tcp 127.0.0.1:9736: connect: connection refused`),
		},
	}

	for _, table := range tables {
		t.Run(table.title, func(t *testing.T) {
			table.app.Action = func(c *cli.Context) error {
				dep, err := loadDependencyGoRedisV8(table.v, table.options, c)
				if !reflect.DeepEqual(table.err, err) {
					t.Errorf("assert err failed, excepted '%s', actual '%s'", table.err, err)
				}

				if table.isNil && dep != nil {
					t.Errorf("assert is nil failed, excepted return nil, but %T received", dep)
				}

				return nil
			}

			if err := table.app.Run(table.args); err != nil {
				t.Fatalf("assert run app failed, error returned %v", err)
				return
			}
		})
	}
}

func Test_loadDependencyStreadwayAmqpConnection(t *testing.T) {
	tables := []struct {
		title   string
		v       reflect.Value
		options map[string]string
		app     *cli.App
		args    []string
		isNil   bool
		err     error
	}{
		{
			title:   "Can load streadway/amqp connection dependency correctly",
			v:       reflect.ValueOf(struct{}{}),
			options: map[string]string{"dsn": "dsnAmqp"},
			app: &cli.App{
				Name: "gocore",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "dsnAmqp"},
				},
			},
			args: []string{
				"gocore",
				"--dsnAmqp", dsnAmqp,
			},
			isNil: false,
			err:   nil,
		},
		{
			title:   "Can't load go-redis v8 dependency correctly, because invalid data source name",
			v:       reflect.ValueOf(struct{}{}),
			options: map[string]string{"dsn": "dsnAmqp"},
			app: &cli.App{
				Name: "gocore",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "dsnAmqp"},
				},
			},
			args: []string{
				"gocore",
				"--dsnAmqp", dsnAmqpInvalid,
			},
			isNil: true,
			err:   fmt.Errorf(`error dial connect to amqp: AMQP scheme must be either 'amqp://' or 'amqps://'`),
		},
	}

	for _, table := range tables {
		t.Run(table.title, func(t *testing.T) {
			table.app.Action = func(c *cli.Context) error {
				dep, err := loadDependencyStreadwayAmqpConnection(table.v, table.options, c)
				if !reflect.DeepEqual(table.err, err) {
					t.Errorf("assert err failed, excepted '%s', actual '%s'", table.err, err)
				}

				if table.isNil && dep != nil {
					t.Errorf("assert is nil failed, excepted return nil, but %T received", dep)
				}

				return nil
			}

			if err := table.app.Run(table.args); err != nil {
				t.Fatalf("assert run app failed, error returned %v", err)
				return
			}
		})
	}
}

func Test_loadDependencyStreadwayAmqpChannel(t *testing.T) {
	mqConnection, err := amqp.Dial(dsnAmqp)
	if err != nil {
		t.Fatalf("error connect to amqp: %v", err)
		return
	}

	mqClosedConnection, err := amqp.Dial(dsnAmqp)
	if err != nil {
		t.Fatalf("error connect to amqp: %v", err)
		return
	}
	if err := mqClosedConnection.Close(); err != nil {
		t.Fatalf("error close connect to amqp: %v", err)
		return
	}

	tables := []struct {
		title   string
		v       reflect.Value
		options map[string]string
		isNil   bool
		err     error
	}{
		{
			title: "Can load streadway/amqp channel dependency correctly",
			v: reflect.ValueOf(struct {
				MQConn *amqp.Connection
			}{
				MQConn: mqConnection,
			}),
			options: map[string]string{"connField": "MQConn"},
			isNil:   false,
			err:     nil,
		},
		{
			title: "Can't load streadway/amqp channel dependency correctly, because nil connection struct",
			v: reflect.ValueOf(struct {
				MQConn *amqp.Connection
			}{}),
			options: map[string]string{"connField": "MQConn"},
			isNil:   true,
			err:     fmt.Errorf(`invalid connField, must be initialized *amqp.Connection instance, but zero`),
		},
		{
			title: "Can't load streadway/amqp channel dependency correctly, because incorrect connField type",
			v: reflect.ValueOf(struct {
				MQConn string
			}{MQConn: "invalid"}),
			options: map[string]string{"connField": "MQConn"},
			isNil:   true,
			err:     fmt.Errorf(`invalid connection, excepted *amqp.Connection, actual string`),
		},
		{
			title: "Can't load streadway/amqp channel dependency correctly, because error in method",
			v: reflect.ValueOf(struct {
				MQConn *amqp.Connection
			}{MQConn: mqClosedConnection}),
			options: map[string]string{"connField": "MQConn"},
			isNil:   true,
			err:     fmt.Errorf(`failed init channel Exception (504) Reason: "channel/connection is not open"`),
		},
	}

	for _, table := range tables {
		t.Run(table.title, func(t *testing.T) {
			app := cli.NewApp()
			app.Name = "gocore"
			app.Action = func(c *cli.Context) error {
				dep, err := loadDependencyStreadwayAmqpChannel(table.v, table.options, c)
				if !reflect.DeepEqual(table.err, err) {
					t.Errorf("assert err failed, excepted '%s', actual '%s'", table.err, err)
				}

				if table.isNil && dep != nil {
					t.Errorf("assert is nil failed, excepted return nil, but %T received", dep)
				}

				return nil
			}

			if err := app.Run([]string{"gocore"}); err != nil {
				t.Fatalf("assert run app failed, error returned %v", err)
				return
			}
		})
	}
}
