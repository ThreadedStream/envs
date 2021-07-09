package envs

import (
	"errors"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	envKw      = "env"
	fallbackKw = "fallback"
)

func Parse(value interface{}) error {
	return parse(value)
}

func parse(value interface{}) error {
	var (
		env      string
		fallback string
		ok       bool
	)

	v := reflect.ValueOf(value).Elem()

	if v.Kind() != reflect.Struct {
		return errors.New("value must be a struct")
	}

	// Iterate through each struct field
	for i := 0; i < v.NumField(); i++ {
		if !v.Field(i).CanAddr() || !v.Field(i).CanSet() {
			log.Printf("field %s is unsettable", v.Type().Field(i).Name)
			continue
		}
		m := parseParams(string(v.Type().Field(i).Tag))
		if env, ok = m[envKw]; !ok {
			// Ignoring this field
			continue
		}
		env = strings.Trim(env, "\"")
		fallback = strings.Trim(m[fallbackKw], "\"")
		switch v.Field(i).Kind() {
		case reflect.String:
			v.Field(i).SetString(envString(env, fallback))
			break
		case reflect.Bool:
			fallbackBool, err := strconv.ParseBool(fallback)
			if err != nil {
				fallbackBool = false
			}

			v.Field(i).SetBool(envBool(env, fallbackBool))
			break
		case reflect.Int:
			fallbackInt, err := strconv.ParseInt(fallback, 10, 64)
			if err != nil {
				log.Println("failed to parse fallback, setting to -1")
				fallbackInt = -1
			}
			v.Field(i).SetInt(envInt(env, fallbackInt))
			break
		}
	}

	return nil
}

func parseParams(tag string) map[string]string {
	params := strings.Split(tag, " ")
	m := make(map[string]string, len(params))
	for _, param := range params {
		paramSet := strings.Split(param, ":")
		var (
			first  = paramSet[0]
			second = paramSet[1]
		)
		m[first] = second
	}

	return m
}

func envString(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func envBool(key string, fallback bool) bool {
	var (
		value   string
		boolVal bool
		err     error
		ok      bool
	)
	if value, ok = os.LookupEnv(key); !ok {
		return fallback
	}

	boolVal, err = strconv.ParseBool(value)
	if err != nil {
		return fallback
	}

	return boolVal
}

func envInt(key string, fallback int64) int64 {
	var (
		value  string
		intVal int64
		err    error
		ok     bool
	)

	if value, ok = os.LookupEnv(key); !ok {
		return fallback
	}
	if intVal, err = strconv.ParseInt(value, 10, 64); err != nil {
		return fallback
	}

	return intVal
}
