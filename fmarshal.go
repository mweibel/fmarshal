package fmarshal

import (
	"fmt"
	"reflect"
	"strings"
)

// MarshalFlag marshals a struct into a slice of flags.
func MarshalFlag(st interface{}, quote bool) []string {
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

		args = append(args, marshalVal(shortOption, quote, name, val)...)
	}
	return args
}

func marshalVal(shortOption, quote bool, name string, v reflect.Value) []string {
	// return nil on nil pointer struct fields
	if !v.IsValid() || !v.CanInterface() {
		return []string{}
	}

	k := v.Kind()

	if k == reflect.Ptr {
		v = v.Elem()

		// return nil on nil pointer struct fields
		if !v.IsValid() || !v.CanInterface() {
			return []string{}
		}

		k = v.Kind()
	}

	val := v.Interface()

	if k == reflect.Slice {
		l := v.Len()
		args := make([]string, 0)
		for i := 0; i < l; i++ {
			cval := v.Index(i)
			args = append(args, marshalVal(shortOption, quote, name, cval)...)
		}
		return args
	}

	strVal := fmt.Sprintf("%v", val)
	if quote {
		strVal = "'" + strings.ReplaceAll(fmt.Sprintf("%v", val), "'", "'\"'\"'") + "'"
	}

	if shortOption {
		return []string{fmt.Sprintf("%s %s", name, strVal)}
	}
	return []string{fmt.Sprintf("%s=%s", name, strVal)}
}
