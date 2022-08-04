package fmarshal

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FlagMarshal(t *testing.T) {
	foo := struct {
		Debug     bool     `flag:"--debug"`
		Level     string   `flag:"--level"`
		Numeric   int      `flag:"-n"`
		SliceTest []string `flag:"--test"`
	}{
		Debug:   true,
		Level:   "info",
		Numeric: 42,
		SliceTest: []string{
			"foo",
			"bar",
			"baz",
		},
	}

	assert.Equal(t, []string{"--debug='true'", "--level='info'", "-n '42'", "--test='foo'", "--test='bar'", "--test='baz'"}, MarshalFlag(foo, true))
	assert.Equal(t, []string{"--debug=true", "--level=info", "-n 42", "--test=foo", "--test=bar", "--test=baz"}, MarshalFlag(foo, false))
}

func Test_FlagMarshalPtr(t *testing.T) {
	debug := true
	level := "info"
	numeric := 42
	sliceTest := []string{"foo", "bar", "baz"}
	foo := struct {
		Debug     *bool     `flag:"--debug"`
		Level     *string   `flag:"--level"`
		Numeric   *int      `flag:"-n"`
		SliceTest *[]string `flag:"--test"`
		Nil       *string   `flag:"--nil"`
	}{
		Debug:     &debug,
		Level:     &level,
		Numeric:   &numeric,
		SliceTest: &sliceTest,
	}

	assert.Equal(t, []string{"--debug='true'", "--level='info'", "-n '42'", "--test='foo'", "--test='bar'", "--test='baz'"}, MarshalFlag(foo, true))
}

func Test_FlagMarshalOmitIfNotSet(t *testing.T) {
	foo := struct {
		Debug   bool   `flag:"--debug,omitempty"`
		Level   string `flag:"--level,omitempty"`
		Numeric int    `flag:"-n,omitempty"`
	}{}

	assert.Equal(t, []string{}, MarshalFlag(foo, true))
}

type Args struct {
	Debug     bool     `flag:"--debug"`
	Level     string   `flag:"--level"`
	Numeric   int      `flag:"-n"`
	SliceTest []string `flag:"--test"`
}

func ExampleMarshalFlag() {
	a := Args{
		Debug:   true,
		Level:   "info",
		Numeric: 42,
		SliceTest: []string{
			"foo",
			"bar",
			"baz",
		},
	}
	fmt.Println(strings.Join(MarshalFlag(a, true), " "))
	// Output: --debug='true' --level='info' -n '42' --test='foo' --test='bar' --test='baz'
}
