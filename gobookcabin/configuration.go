package gobookcabin

import (
	"fmt"
	"github.com/joho/godotenv"
	"reflect"
	"strings"
)

const (
	EnvFileName = ".env"
)

type AppConfiguration struct {
	DBString string `env:"DB_STRING"`

	ServerPort string `env:"SERVER_PORT"`
}

// loadConfiguration initializes the environment configuration by loading variables from .env files
// By default, if the key is not found in .env, the loadConfiguration() function will return an error
// The target is an instance of .env
func loadConfiguration(target any) error {
	envMap, err := godotenv.Read(EnvFileName)
	if err != nil {
		return err
	}

	err = mapToObject(envMap, target)

	return err
}

// mapToObject maps a Go map into struct based on "env" struct tags.
func mapToObject(sourceMap map[string]string, target any) error {
	val := reflect.ValueOf(target)
	if val.IsNil() || val.Kind() != reflect.Ptr {
		return fmt.Errorf("target must be a non-nil pointer to a struct")
	}

	elem := val.Elem()
	if elem.Kind() != reflect.Struct {
		return fmt.Errorf("target must be a pointer to a struct")
	}

	typ := elem.Type()
	for i := 0; i < elem.NumField(); i++ {
		fieldType := typ.Field(i)

		tag := fieldType.Tag.Get("env")
		if tag == "" {
			continue
		}

		value, ok := sourceMap[tag]
		if !ok {
			return fmt.Errorf("expected key %s to appear in map, but was null. check your .env file", tag)
		} else {
			if elem.Field(i).Type().Kind() == reflect.Slice || elem.Field(i).Type().Kind() == reflect.Array {
				for _, splitValue := range strings.Split(value, ",") {
					elem.Field(i).Set(reflect.Append(elem.Field(i), reflect.ValueOf(splitValue)))
				}
			} else {
				elem.Field(i).SetString(value)
			}
		}
	}
	return nil
}
