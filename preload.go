package glp

import (
	encoding "github.com/vikingo-project/glp/encoding"
	zlib "github.com/vikingo-project/glp/zlib"
	lua "github.com/yuin/gopher-lua"
)

func Preload(L *lua.LState) {
	zlib.Preload(L)
	encoding.Preload(L)
}
