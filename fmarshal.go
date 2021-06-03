package fmarshal

import (
	"fmt"
	"reflect"
	"strings"
)

func MarshalFlag(st interface{}) []string {
	t := reflect.TypeOf(st)
	v := reflect.ValueOf(st)

	numFields := t.NumField()
	args := make([]string, numFields)
	for i := 0; i < numFields; i++ {
		field := t.Field(i)
		val := v.Field(i)

		// skip unexported fields
		if !val.IsValid() || !val.CanInterface() {
			continue
		}

		flag := field.Tag.Get("flag")
		// typ := field.Type

		shortOption := strings.Count(flag, "-") == 1

		args[i] = marshalVal(shortOption, flag, val.Interface())
	}
	return args
}

func marshalVal(shortOption bool, name string, val interface{}) string {
	if shortOption {
		return fmt.Sprintf("%s %v", name, val)
	}
	return fmt.Sprintf("%s=%v", name, val)
}
