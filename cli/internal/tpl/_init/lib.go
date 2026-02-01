package rolib

import (
	rt "github.com/arnodel/golua/runtime"
)

func Init() *rt.Table {
	pkg := rt.NewTable()

	pkg.Set(
		rt.StringValue("name"),
		rt.StringValue("{{.LibName}}"),
	)

	return pkg
}
