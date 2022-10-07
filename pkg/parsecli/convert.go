package parsecli

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/serenize/snaker"
	"github.com/urfave/cli/v2"
)

func Convert(src interface{}, prefixEnv string) []cli.Flag {
	configValue := reflect.Indirect(reflect.ValueOf(src))
	var flags []cli.Flag
	for i := 0; i < configValue.NumField(); i++ {
		fieldValue := configValue.Field(i)
		fieldType := configValue.Type().Field(i)
		name := snaker.CamelToSnake(fieldType.Name)
		flagName := fieldType.Tag.Get("cli")
		if flagName == "" {
			flagName = name
		}
		envName := fieldType.Tag.Get("env")
		if envName == "" {
			envName = strings.ToUpper(flagName)
		}
		envName = prefixEnv + envName
		switch fieldType.Type.Kind() {
		case reflect.Struct:
			flags = append(flags, Convert(fieldValue.Interface(), envName+"_")...)
		case reflect.String:
			flag := &cli.StringFlag{
				Name:    flagName,
				EnvVars: []string{envName},
				Value:   fieldType.Tag.Get("default"),
			}
			flags = append(flags, flag)
		case reflect.Bool:
			flag := &cli.BoolFlag{
				Name:    flagName,
				EnvVars: []string{envName},
				Value:   boolFromString(fieldType.Tag.Get("default")),
			}
			flags = append(flags, flag)
		case reflect.Int:
			flag := &cli.IntFlag{
				Name:    flagName,
				EnvVars: []string{envName},
				Value:   intFromString(fieldType.Tag.Get("default")),
			}
			flags = append(flags, flag)
		case reflect.Int64:
			flag := &cli.Int64Flag{
				Name:    flagName,
				EnvVars: []string{envName},
				Value:   int64FromString(fieldType.Tag.Get("default")),
			}
			flags = append(flags, flag)

		case reflect.Slice:
			if fieldType.Type.Elem().Kind() == reflect.String {
				values := strings.Split(fieldType.Tag.Get("default"), ",")
				fieldValue.Set(reflect.ValueOf(values))
				flag := &cli.StringSliceFlag{
					Name:    flagName,
					EnvVars: []string{envName},
					Value:   cli.NewStringSlice(values...),
				}
				flags = append(flags, flag)
			}
			if fieldType.Type.Elem().Kind() == reflect.Int {
				values := strings.Split(fieldType.Tag.Get("default"), ",")
				var intValues []int
				for _, v := range values {
					intValues = append(intValues, intFromString(v))
				}
				fieldValue.Set(reflect.ValueOf(intValues))
				flag := &cli.IntSliceFlag{
					Name:    flagName,
					EnvVars: []string{envName},
					Value:   cli.NewIntSlice(intValues...),
				}
				flags = append(flags, flag)
			}
		}
	}
	return flags
}
func int64FromString(s string) int64 {
	val, _ := strconv.ParseInt(s, 10, 64)
	return val
}

func intFromString(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func boolFromString(s string) bool {
	val, _ := strconv.ParseBool(s)
	return val
}
