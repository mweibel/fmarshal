package fmarshal

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FlagMarshal(t *testing.T) {
	foo := struct {
		Debug   bool   `flag:"--debug"`
		Level   string `flag:"--level"`
		Numeric int    `flag:"-n"`
	}{
		Debug:   true,
		Level:   "info",
		Numeric: 42,
	}

	assert.Equal(t, []string{"--debug=true", "--level=info", "-n 42"}, MarshalFlag(foo))
}

type Args struct {
	Debug   bool   `flag:"--debug"`
	Level   string `flag:"--level"`
	Numeric int    `flag:"-n"`
}

func ExampleMarshal() {
	a := Args{
		Debug:   true,
		Level:   "info",
		Numeric: 42,
	}
	fmt.Println(strings.Join(MarshalFlag(a), " "))
	// Output: --debug=true --level=info -n 42
}
