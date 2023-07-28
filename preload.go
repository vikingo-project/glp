package glp

import (
	crypto "github.com/vikingo-project/glp/crypto"
	encoding "github.com/vikingo-project/glp/encoding"
	zlib "github.com/vikingo-project/glp/zlib"
	lua "github.com/yuin/gopher-lua"
)

func Preload(L *lua.LState) {
	zlib.Preload(L)
	encoding.Preload(L)
	crypto.Preload(L)
}
