package corecliapp

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-redis/redis/v8"
	"github.com/memclutter/gocore/pkg/corereflect"
	"github.com/memclutter/gocore/pkg/coreslices"
	"github.com/streadway/amqp"
	"github.com/urfave/cli/v2"
	"reflect"
	"strings"
)

func setFlags(c *cli.Context, rFlags reflect.Value) {
	rFlags = corereflect.PtrValueOf(rFlags)
	rtFlags := corereflect.PtrTypeOf(rFlags)
	for j := 0; j < rtFlags.NumField(); j++ {
		rfField := rtFlags.Field(j)
		rField := rFlags.Field(j)

		if rfField.Anonymous {
			setFlags(c, rField)
			continue
		}

		name := rfField.Tag.Get("cli.flag.name")
		switch rfField.Type.Kind() {
		case reflect.String:
			rField.SetString(c.String(name))
		case reflect.Bool:
			rField.SetBool(c.Bool(name))
		case reflect.Int:
			rField.SetInt(int64(c.Int(name)))
		case reflect.Int64:
			rField.SetInt(int64(c.Int64(name)))
		}
	}
}


func setServices(c *cli.Context, rCommand reflect.Value) error {
	rCommand = corereflect.PtrValueOf(rCommand)
	rtCommand := corereflect.PtrTypeOf(rCommand)

	for j := 0; j < rtCommand.NumField(); j++ {
		rfField := rtCommand.Field(j)
		rftField := rfField.Type
		if rftField.Kind() == reflect.Ptr {
			rftField = rftField.Elem()
		}
		service := strings.Split(rfField.Tag.Get("cli.service"), ",")
		service = coreslices.StringApply(service, func(i int, s string) string {return strings.TrimSpace(s)})
		service = coreslices.StringFilter(service, func(i int, s string) bool {return len(s) > 0})

		if len(service) == 0 {
			continue
		}

		serviceOptions := map[string]string{}
		for _, s := range service {
			kv := strings.Split(s, "=")
			if len(kv) != 2 {
				continue
			}
			serviceOptions[kv[0]] = kv[1]
		}

		switch rftField.PkgPath() {
		case "github.com/go-pg/pg/v10":
			opt, err := pg.ParseURL(c.String(serviceOptions["dsnFromFlags"]))
			if err != nil {
				return fmt.Errorf("github.com/go-pg/pg/v10: %v", err)
			}
			db := pg.Connect(opt)
			rCommand.Field(j).Set(reflect.ValueOf(db))
		case "github.com/go-redis/redis/v8":
			optRd, err := redis.ParseURL(c.String(serviceOptions["dsnFromFlags"]))
			if err != nil {
				return fmt.Errorf("github.com/go-redis/redis/v8: %v", err)
			}
			cache := redis.NewClient(optRd)
			rCommand.Field(j).Set(reflect.ValueOf(cache))
		case "github.com/streadway/amqp":
			conn, err := amqp.Dial(c.String(serviceOptions["dsnFromFlags"]))
			if err != nil {
				return fmt.Errorf("github.com/streadway/amqp: %v", err)
			}
			rCommand.Field(j).Set(reflect.ValueOf(conn))
		default:
			return fmt.Errorf("unsupport cli.service '%s'", rftField.PkgPath())
		}
	}

	return nil
}

func callRun(rCommand reflect.Value) error {
	rResult := rCommand.MethodByName("Run").Call([]reflect.Value{})
	if len(rResult) == 0 || rResult[0].IsNil() {
		return nil
	}
	return rResult[0].Interface().(error)
}
