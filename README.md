# fmarshal
[![GoDoc](https://godoc.org/github.com/mweibel/fmarshal?status.svg)](https://godoc.org/github.com/mweibel/fmarshal) ![build](https://github.com/mweibel/fmarshal/actions/workflows/go.yml/badge.svg)

```
go get github.com/mweibel/fmarshal
```

Package fmarshal marshals a struct type into a slice of CLI arguments.

**Note:** This package does not do any shell escaping. It just quotes flag values if quote is true.

Example:

```golang
type Args struct {
  Debug   bool   `flag:"--debug"`
  Level   string `flag:"--level"`
  Numeric int    `flag:"-n"`
}

a := Args{
  Debug: true,
  Level: "info",
  Numeric: 42
}
fmt.Println(strings.Join(MarshalFlag(a, true), " "))
// Output: --debug=true --level=info -n 42
```