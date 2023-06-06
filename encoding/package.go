package encoding

import (
	"encoding/base64"
	"encoding/hex"
	"net/url"

	lua "github.com/yuin/gopher-lua"
	"golang.org/x/net/idna"
)

func Preload(L *lua.LState) {
	L.PreloadModule("encoding", Loader)
}

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"hex_encode":    hexEncode,
		"hex_decode":    hexDecode,
		"url_encode":    urlEncode,
		"url_decode":    urlDecode,
		"base64_encode": base64Encode,
		"base64_decode": base64Decode,
		"puny_encode":   punyEncode,
		"puny_decode":   punyDecode,
	})
	L.Push(mod)
	return 1
}

func hexEncode(L *lua.LState) int {
	s := L.CheckString(1)
	encoded := hex.EncodeToString([]byte(s))
	L.Push(lua.LString(encoded))
	return 1
}

func hexDecode(L *lua.LState) int {
	s := L.CheckString(1)
	decoded, err := hex.DecodeString(s)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LString(string(decoded)))
	return 1
}

func urlEncode(L *lua.LState) int {
	query := L.CheckString(1)
	escapedUrl := url.QueryEscape(query)
	L.Push(lua.LString(escapedUrl))
	return 1
}

func urlDecode(L *lua.LState) int {
	query := L.CheckString(1)
	url, err := url.QueryUnescape(query)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LString(url))
	return 1
}

func base64Encode(L *lua.LState) int {
	s := L.CheckString(1)
	encoded := base64.StdEncoding.EncodeToString([]byte(s))
	L.Push(lua.LString(encoded))
	return 1
}

func base64Decode(L *lua.LState) int {
	s := L.CheckString(1)
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LString(string(decoded)))
	return 1
}

func punyEncode(L *lua.LState) int {
	s := L.CheckString(1)
	p := idna.New()
	encoded, err := p.ToASCII(s)
	println(encoded, err)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LString(encoded))
	return 1
}

func punyDecode(L *lua.LState) int {
	s := L.CheckString(1)
	p := idna.New()
	decoded, err := p.ToUnicode(s)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LString(decoded))
	return 1
}
