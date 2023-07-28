package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	lua "github.com/yuin/gopher-lua"
)

func Preload(L *lua.LState) {
	L.PreloadModule("crypto", Loader)
}

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"md5":    md5Hash,
		"sha1":   sha1Hash,
		"sha256": sha256Hash,
		"sha512": sha512Hash,
		"xor":    xor,
	})
	L.Push(mod)
	return 1
}

func md5Hash(L *lua.LState) int {
	s := L.CheckString(1)
	hash := md5.Sum([]byte(s))
	L.Push(lua.LString(hex.EncodeToString(hash[:])))
	return 1
}

func sha1Hash(L *lua.LState) int {
	s := L.CheckString(1)
	hash := sha1.Sum([]byte(s))
	L.Push(lua.LString(hex.EncodeToString(hash[:])))
	return 1
}

func sha256Hash(L *lua.LState) int {
	s := L.CheckString(1)
	hash := sha256.Sum256([]byte(s))
	L.Push(lua.LString(hex.EncodeToString(hash[:])))
	return 1
}
func sha512Hash(L *lua.LState) int {
	s := L.CheckString(1)
	hash := sha512.Sum512([]byte(s))
	L.Push(lua.LString(hex.EncodeToString(hash[:])))
	return 1
}

func xorNums(data, key []int64) []int64 {
	r := []int64{}
	for i := 0; i < len(data); i++ {
		r = append(r, data[i]^key[i%len(key)])
	}
	return r
}

func argToNums(L *lua.LState, index int) []int64 {
	d := L.Get(index)
	if d.Type() == lua.LTNil {
		return []int64{}
	}
	switch d.Type() {
	case lua.LTNumber:
		return []int64{int64(d.(lua.LNumber))}
	case lua.LTString:
		nums := []int64{}
		for _, b := range d.String() {
			nums = append(nums, int64(b))
		}
		return nums
	case lua.LTTable:
		nums := []int64{}
		if elems, ok := d.(*lua.LTable); ok {
			elems.ForEach(func(_ lua.LValue, v lua.LValue) {
				nums = append(nums, int64(v.(lua.LNumber)))
			})
		}
		return nums
	}
	return []int64{}
}

func xor(L *lua.LState) int {
	d, k := argToNums(L, 1), argToNums(L, 2)
	if len(d) < 1 || len(k) < 1 {
		L.RaiseError("parameters should be number, table or string")
	}
	res := xorNums(d, k)

	retType := L.Get(1)
	switch retType.Type() {
	case lua.LTNumber:
		L.Push(lua.LNumber(res[0]))
		return 1
	case lua.LTString:
		r := []byte{}
		for _, n := range res {
			r = append(r, byte(n))
		}
		L.Push(lua.LString(r))
		return 1
	case lua.LTTable:
		t := L.NewTable()
		for _, n := range res {
			t.Append(lua.LNumber(n))
		}
		L.Push(t)
		return 1
	}
	return 0
}
