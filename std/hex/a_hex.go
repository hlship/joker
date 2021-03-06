// This file is generated by generate-std.joke script. Do not edit manually!

package hex

import (
	. "github.com/candid82/joker/core"
	"encoding/hex"
)

var hexNamespace = GLOBAL_ENV.EnsureNamespace(MakeSymbol("joker.hex"))



var decode_string_ Proc

func __decode_string_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		s := ExtractString(_args, 0)
		 t, err := hex.DecodeString(s)
		PanicOnErr(err)
		_res := string(t)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var encode_string_ Proc

func __encode_string_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		s := ExtractString(_args, 0)
		_res := hex.EncodeToString([]byte(s))
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

func Init() {

	decode_string_ = __decode_string_
	encode_string_ = __encode_string_

	hexNamespace.ResetMeta(MakeMeta(nil, `Implements hexadecimal encoding and decoding.`, "1.0"))

	
	hexNamespace.InternVar("decode-string", decode_string_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Returns the bytes represented by the hexadecimal string s.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	hexNamespace.InternVar("encode-string", encode_string_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Returns the hexadecimal encoding of s.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

}

func init() {
	hexNamespace.Lazy = Init
}
