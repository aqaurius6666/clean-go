package parsecli

import (
	"fmt"
	"reflect"

	"github.com/urfave/cli/v2"
)

func Parse(c *cli.Context, dest interface{}) error {
	cfg := reflect.Indirect(reflect.ValueOf(dest))
	for i := 0; i < cfg.NumField(); i++ {
		field := cfg.Field(i)
		if field.Kind() == reflect.Struct {
			if err := Parse(c, field.Addr().Interface()); err != nil {
				return err
			}
			continue
		}
		cliTag := cfg.Type().Field(i).Tag.Get("cli")
		if cliTag == "" {
			continue
		}
		switch field.Kind() {

		case reflect.String:
			field.SetString(c.String(cliTag))
		case reflect.Bool:
			field.SetBool(c.Bool(cliTag))
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			field.SetInt(int64(c.Int(cliTag)))
		case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			field.SetUint(uint64(c.Uint(cliTag)))
		case reflect.Float32, reflect.Float64:
			field.SetFloat(float64(c.Float64(cliTag)))
		case reflect.Array, reflect.Slice:
			g := c.Generic(cliTag)
			switch g.(type) {
			case *cli.StringSlice:
				field.Set(reflect.ValueOf(c.StringSlice(cliTag)))
			case *cli.Int64Slice:
				field.Set(reflect.ValueOf(c.Int64Slice(cliTag)))
			case *cli.Float64Slice:
				field.Set(reflect.ValueOf(c.Float64Slice(cliTag)))
			case *cli.Uint64Slice:
				field.Set(reflect.ValueOf(c.Uint64Slice(cliTag)))
			case *cli.IntSlice:
				field.Set(reflect.ValueOf(c.IntSlice(cliTag)))
			default:
				return fmt.Errorf("unsupported type %s", field.Kind())
			}
		default:
			return fmt.Errorf("unsupported type: %v", field.Kind())
		}
	}
	return nil
}
