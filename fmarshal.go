package fmarshal

import (
	"fmt"
	"reflect"
	"strings"
)

// MarshalFlag marshals a struct into a slice of flags.
//
// - quote will quote each value
// - separateKeyVal will make sure key and value are separate items in the returned slice.
func MarshalFlag(st interface{}, quote, separateKeyVal bool) []string {
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

		tag := strings.Split(field.Tag.Get("flag"), ",")
		name := tag[0]
		omitempty := false
		if len(tag) == 2 && tag[1] == "omitempty" {
			omitempty = true
		}

		shortOption := strings.Count(name, "-") == 1

		args = append(args, marshalVal(shortOption, quote, separateKeyVal, name, val, omitempty)...)
	}
	return args
}

func marshalVal(shortOption, quote, separateKeyVal bool, name string, v reflect.Value, omitempty bool) []string {
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

	if v.IsZero() && omitempty {
		return []string{}
	}

	val := v.Interface()

	if k == reflect.Slice {
		l := v.Len()
		args := make([]string, 0)
		for i := 0; i < l; i++ {
			cval := v.Index(i)
			args = append(args, marshalVal(shortOption, quote, separateKeyVal, name, cval, omitempty)...)
		}
		return args
	}

	strVal := fmt.Sprintf("%v", val)
	if strVal == "" && omitempty {
		return []string{}
	}
	if quote {
		strVal = "'" + strings.ReplaceAll(fmt.Sprintf("%v", val), "'", "'\"'\"'") + "'"
	}

	if separateKeyVal {
		return []string{name, strVal}
	}
	if shortOption {
		return []string{fmt.Sprintf("%s %s", name, strVal)}
	}
	return []string{fmt.Sprintf("%s=%s", name, strVal)}
}
