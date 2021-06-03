package fmarshal

import (
	"fmt"
	"reflect"
	"strings"
)

// MarshalFlag marshals a struct into a slice of flags.
func MarshalFlag(st interface{}) []string {
	t := reflect.TypeOf(st)
	v := reflect.ValueOf(st)

	numFields := t.NumField()
	args := make([]string, 0)
	for i := 0; i < numFields; i++ {
		field := t.Field(i)
		val := v.Field(i)

		// skip unexported fields
		if !val.IsValid() || !val.CanInterface() {
			continue
		}

		name := field.Tag.Get("flag")

		shortOption := strings.Count(name, "-") == 1

		args = append(args, marshalVal(shortOption, name, val)...)
	}
	return args
}

func marshalVal(shortOption bool, name string, val reflect.Value) []string {
	k := val.Kind()
	if k == reflect.Slice {
		l := val.Len()
		args := make([]string, 0)
		for i := 0; i < l; i++ {
			cval := val.Index(i)
			args = append(args, marshalVal(shortOption, name, cval)...)
		}
		return args
	}

	strVal := "'" + strings.ReplaceAll(fmt.Sprintf("%v", val), "'", "'\"'\"'") + "'"

	if shortOption {
		return []string{fmt.Sprintf("%s %s", name, strVal)}
	}
	return []string{fmt.Sprintf("%s=%s", name, strVal)}
}
