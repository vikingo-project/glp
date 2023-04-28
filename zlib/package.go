package lzlib

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"

	lua "github.com/yuin/gopher-lua"
)

func Preload(L *lua.LState) {
	L.PreloadModule("zlib", Loader)
}

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"compress":   compress,
		"decompress": decompress,
	})
	L.Push(mod)
	return 1
}

func compress(L *lua.LState) int {
	s := L.CheckString(1)
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write([]byte(s))
	w.Close()
	L.Push(lua.LString(b.Bytes()))
	return 1
}

func decompress(L *lua.LState) int {
	s := L.CheckString(1)
	b := bytes.NewReader([]byte(s))
	r, err := zlib.NewReader(b)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
	}
	defer r.Close()
	data, err := ioutil.ReadAll(r)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
	}
	L.Push(lua.LString(data))
	return 1
}
