package corecli

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
	"github.com/urfave/cli/v2"
	"io"
	"reflect"
	"strconv"
	"strings"
)

// LoadDependencies godoc
//
// Load dependencies like as postgres database, rabbitmq connection or redis client
//
// type CommandFlags struct {
//   DsnCache string
// }
// type Command struct {
//   Flags  CommandFlags 	 `cli.command.flags:"*"`
//   DB     *pg.DB           `cli.command.dependency:"dsn:dsnDb"
//   MQConn *amqp.Connection `cli.command.dependency:"dsn:dsnAmqp"
//   MQCh   *amqp.Channel    `cli.command.dependency:"connField:MQConn"
//   Cache  *redis.Client    `cli.command.dependency:"dsn:dsnCache"
// }
func LoadDependencies(v reflect.Value, c *cli.Context) error {
	valueOf := reflect.Indirect(v)
	typeOf := valueOf.Type()

	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		fieldValueOf := valueOf.Field(i)
		dependencyTag, ok := field.Tag.Lookup(`cli.command.dependency`)
		if !ok {
			continue
		}
		dependencyOptions := stringToStringMap(dependencyTag)

		pkgPath := field.Type.PkgPath()
		typeName := field.Type.Name()
		if field.Type.Kind() == reflect.Ptr {
			pkgPath = field.Type.Elem().PkgPath()
			typeName = field.Type.Elem().Name()
		}
		pkgPath = strings.Join([]string{pkgPath, typeName}, "::")

		loader, ok := map[string]dependencyLoader{
			"github.com/go-pg/pg/v10::DB":           loadDependencyGoPgV10,
			"github.com/go-redis/redis/v8::Client":  loadDependencyGoRedisV8,
			"github.com/streadway/amqp::Connection": loadDependencyStreadwayAmqpConnection,
			"github.com/streadway/amqp::Channel":    loadDependencyStreadwayAmqpChannel,
		}[pkgPath]
		if !ok {
			return fmt.Errorf("unsupported depdendency package %s", pkgPath)
		}

		dependency, err := loader(valueOf, dependencyOptions, c)
		if err != nil {
			return fmt.Errorf("%s: %v", pkgPath, err)
		}
		fieldValueOf.Set(reflect.ValueOf(dependency))
	}

	return nil
}

type dependencyLoader func(v reflect.Value, options map[string]string, c *cli.Context) (interface{}, error)

func loadDependencyGoPgV10(v reflect.Value, options map[string]string, c *cli.Context) (interface{}, error) {
	opt, err := pg.ParseURL(c.String(options["dsn"]))
	if err != nil {
		return nil, fmt.Errorf("error parse data source name: %v", err)
	}
	if option, ok := options["poolSize"]; ok {
		poolSize, err := strconv.Atoi(option)
		if err != nil {
			return nil, fmt.Errorf("invalid parse poolSize option: %v", err)
		}
		opt.PoolSize = poolSize
	}
	db := pg.Connect(opt)
	if _, ok := options["ping"]; ok {
		if err := db.Ping(c.Context); err != nil {
			return nil, fmt.Errorf("error connect to database: %v", err)
		}
	}
	return db, nil
}

func loadDependencyGoRedisV8(v reflect.Value, options map[string]string, c *cli.Context) (interface{}, error) {
	optRd, err := redis.ParseURL(c.String(options["dsn"]))
	if err != nil {
		return nil, fmt.Errorf("error parse data source name: %v", err)
	}
	cache := redis.NewClient(optRd)
	if _, ok := options["ping"]; ok {
		if err := cache.Ping(c.Context).Err(); err != nil {
			return nil, fmt.Errorf("error connect to redis: %v", err)
		}
	}
	return cache, nil
}

func loadDependencyStreadwayAmqpConnection(v reflect.Value, options map[string]string, c *cli.Context) (interface{}, error) {
	conn, err := amqp.Dial(c.String(options["dsn"]))
	if err != nil {
		return nil, fmt.Errorf("error dial connect to amqp: %v", err)
	}
	return conn, nil
}

func loadDependencyStreadwayAmqpChannel(v reflect.Value, options map[string]string, c *cli.Context) (interface{}, error) {
	connField := v.FieldByName(options["connField"])
	if connField.IsZero() || !connField.IsValid() {
		return nil, fmt.Errorf("invalid connField, must be initialized *amqp.Connection instance, but zero")
	}
	conn, ok := connField.Interface().(*amqp.Connection)
	if !ok {
		return nil, fmt.Errorf("invalid connection, excepted *amqp.Connection, actual %T", connField.Interface())
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed init channel %v", err)
	}
	return ch, nil
}

func CloseDependencies(v reflect.Value) error {
	valueOf := reflect.Indirect(v)
	typeOf := valueOf.Type()

	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		fieldValueOf := valueOf.Field(i)
		if _, ok := field.Tag.Lookup(`cli.command.dependency`); !ok {
			continue
		}

		pkgPath := field.Type.PkgPath()
		typeName := field.Type.Name()
		if field.Type.Kind() == reflect.Ptr {
			pkgPath = field.Type.Elem().PkgPath()
			typeName = field.Type.Elem().Name()
		}
		pkgPath = strings.Join([]string{pkgPath, typeName}, "::")

		isCloser := map[string]bool{
			"github.com/go-pg/pg/v10::DB":           true,
			"github.com/go-redis/redis/v8::Client":  true,
			"github.com/streadway/amqp::Connection": true,
			"github.com/streadway/amqp::Channel":    false,
		}[pkgPath]

		if isCloser {
			if closer, ok := fieldValueOf.Interface().(io.Closer); ok && closer != nil {
				if err := closer.Close(); err != nil {
					return fmt.Errorf("%s: error close %v", pkgPath, err)
				}
			}
		}
	}

	return nil
}
