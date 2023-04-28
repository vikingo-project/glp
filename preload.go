package glp

import (
	zlib "github.com/vikingo-project/glp/zlib"
	lua "github.com/yuin/gopher-lua"
)

func Preload(L *lua.LState) {
	zlib.Preload(L)
}
