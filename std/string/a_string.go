// This file is generated by generate-std.joke script. Do not edit manually!

package string

import (
	. "github.com/candid82/joker/core"
	"strings"
	"unicode"
)

var stringNamespace = GLOBAL_ENV.EnsureNamespace(MakeSymbol("joker.string"))



var isblank_ Proc

func __isblank_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		s := ExtractObject(_args, 0)
		_res := isBlank(s)
		return MakeBoolean(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var capitalize_ Proc

func __capitalize_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		s := ExtractStringable(_args, 0)
		_res := capitalize(s)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var isends_with_ Proc

func __isends_with_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		s := ExtractString(_args, 0)
		substr := ExtractStringable(_args, 1)
		_res := strings.HasSuffix(s, substr)
		return MakeBoolean(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var escape_ Proc

func __escape_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		s := ExtractString(_args, 0)
		cmap := ExtractCallable(_args, 1)
		_res := escape(s, cmap)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var isincludes_ Proc

func __isincludes_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		s := ExtractString(_args, 0)
		substr := ExtractStringable(_args, 1)
		_res := strings.Contains(s, substr)
		return MakeBoolean(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var index_of_ Proc

func __index_of_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		s := ExtractString(_args, 0)
		value := ExtractObject(_args, 1)
		_res := indexOf(s, value, 0)
		return _res

	case _c == 3:
		s := ExtractString(_args, 0)
		value := ExtractObject(_args, 1)
		from := ExtractInt(_args, 2)
		_res := indexOf(s, value, from)
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var join_ Proc

func __join_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		coll := ExtractSeqable(_args, 0)
		_res := join("", coll)
		return MakeString(_res)

	case _c == 2:
		separator := ExtractStringable(_args, 0)
		coll := ExtractSeqable(_args, 1)
		_res := join(separator, coll)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var last_index_of_ Proc

func __last_index_of_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		s := ExtractString(_args, 0)
		value := ExtractObject(_args, 1)
		_res := lastIndexOf(s, value, 0)
		return _res

	case _c == 3:
		s := ExtractString(_args, 0)
		value := ExtractObject(_args, 1)
		from := ExtractInt(_args, 2)
		_res := lastIndexOf(s, value, from)
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var lower_case_ Proc

func __lower_case_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		s := ExtractStringable(_args, 0)
		_res := strings.ToLower(s)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var pad_left_ Proc

func __pad_left_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 3:
		s := ExtractString(_args, 0)
		pad := ExtractStringable(_args, 1)
		n := ExtractInt(_args, 2)
		_res := padLeft(s, pad, n)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var pad_right_ Proc

func __pad_right_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 3:
		s := ExtractString(_args, 0)
		pad := ExtractStringable(_args, 1)
		n := ExtractInt(_args, 2)
		_res := padRight(s, pad, n)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var replace_ Proc

func __replace_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 3:
		s := ExtractString(_args, 0)
		match := ExtractObject(_args, 1)
		repl := ExtractStringable(_args, 2)
		_res := replace(s, match, repl)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var replace_first_ Proc

func __replace_first_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 3:
		s := ExtractString(_args, 0)
		match := ExtractObject(_args, 1)
		repl := ExtractStringable(_args, 2)
		_res := replaceFirst(s, match, repl)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var reverse_ Proc

func __reverse_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		s := ExtractString(_args, 0)
		_res := reverse(s)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var split_ Proc

func __split_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		s := ExtractString(_args, 0)
		re := ExtractRegex(_args, 1)
		_res := split(s, re, 0)
		return _res

	case _c == 3:
		s := ExtractString(_args, 0)
		re := ExtractRegex(_args, 1)
		n := ExtractInt(_args, 2)
		_res := split(s, re, n)
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var split_lines_ Proc

func __split_lines_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		s := ExtractString(_args, 0)
		_res := split(s, newLine, 0)
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var isstarts_with_ Proc

func __isstarts_with_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		s := ExtractString(_args, 0)
		substr := ExtractStringable(_args, 1)
		_res := strings.HasPrefix(s, substr)
		return MakeBoolean(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var trim_ Proc

func __trim_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		s := ExtractString(_args, 0)
		_res := strings.TrimSpace(s)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var trim_left_ Proc

func __trim_left_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		s := ExtractString(_args, 0)
		_res := strings.TrimLeftFunc(s, unicode.IsSpace)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var trim_newline_ Proc

func __trim_newline_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		s := ExtractString(_args, 0)
		_res := strings.TrimRight(s, "\n\r")
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var trim_right_ Proc

func __trim_right_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		s := ExtractString(_args, 0)
		_res := strings.TrimRightFunc(s, unicode.IsSpace)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var triml_ Proc

func __triml_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		s := ExtractString(_args, 0)
		_res := strings.TrimLeftFunc(s, unicode.IsSpace)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var trimr_ Proc

func __trimr_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		s := ExtractString(_args, 0)
		_res := strings.TrimRightFunc(s, unicode.IsSpace)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var upper_case_ Proc

func __upper_case_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		s := ExtractStringable(_args, 0)
		_res := strings.ToUpper(s)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

func Init() {

	isblank_ = __isblank_
	capitalize_ = __capitalize_
	isends_with_ = __isends_with_
	escape_ = __escape_
	isincludes_ = __isincludes_
	index_of_ = __index_of_
	join_ = __join_
	last_index_of_ = __last_index_of_
	lower_case_ = __lower_case_
	pad_left_ = __pad_left_
	pad_right_ = __pad_right_
	replace_ = __replace_
	replace_first_ = __replace_first_
	reverse_ = __reverse_
	split_ = __split_
	split_lines_ = __split_lines_
	isstarts_with_ = __isstarts_with_
	trim_ = __trim_
	trim_left_ = __trim_left_
	trim_newline_ = __trim_newline_
	trim_right_ = __trim_right_
	triml_ = __triml_
	trimr_ = __trimr_
	upper_case_ = __upper_case_

	stringNamespace.ResetMeta(MakeMeta(nil, `Implements simple functions to manipulate strings.`, "1.0"))

	
	stringNamespace.InternVar("blank?", isblank_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`True if s is nil, empty, or contains only whitespace.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Boolean"}))

	stringNamespace.InternVar("capitalize", capitalize_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Converts first character of the string to upper-case, all other
  characters to lower-case.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("ends-with?", isends_with_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("substr"))),
			`True if s ends with substr.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Boolean"}))

	stringNamespace.InternVar("escape", escape_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("cmap"))),
			`Return a new string, using cmap to escape each character ch
  from s as follows:

  If (cmap ch) is nil, append ch to the new string.
  If (cmap ch) is non-nil, append (str (cmap ch)) instead.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("includes?", isincludes_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("substr"))),
			`True if s includes substr.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Boolean"}))

	stringNamespace.InternVar("index-of", index_of_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("value")), NewVectorFrom(MakeSymbol("s"), MakeSymbol("value"), MakeSymbol("from"))),
			`Return index of value (string or char) in s, optionally searching
  forward from from or nil if not found.`, "1.0"))

	stringNamespace.InternVar("join", join_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("coll")), NewVectorFrom(MakeSymbol("separator"), MakeSymbol("coll"))),
			`Returns a string of all elements in coll, as returned by (seq coll), separated by an optional separator.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("last-index-of", last_index_of_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("value")), NewVectorFrom(MakeSymbol("s"), MakeSymbol("value"), MakeSymbol("from"))),
			`Return last index of value (string or char) in s, optionally
  searching backward from from or nil if not found.`, "1.0"))

	stringNamespace.InternVar("lower-case", lower_case_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Converts string to all lower-case.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("pad-left", pad_left_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("pad"), MakeSymbol("n"))),
			`Returns s padded with pad at the beginning to length n.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("pad-right", pad_right_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("pad"), MakeSymbol("n"))),
			`Returns s padded with pad at the end to length n.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("replace", replace_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("match"), MakeSymbol("repl"))),
			`Replaces all instances of match (String or Regex) with string repl in string s.

  If match is Regex, $1, $2, etc. in the replacement string repl are
  substituted with the string that matched the corresponding
  parenthesized group in the pattern.
  `, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("replace-first", replace_first_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("match"), MakeSymbol("repl"))),
			`Replaces the first instance of match (String or Regex) with string repl in string s.

  If match is Regex, $1, $2, etc. in the replacement string repl are
  substituted with the string that matched the corresponding
  parenthesized group in the pattern.
  `, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("reverse", reverse_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Returns s with its characters reversed.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("split", split_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("re")), NewVectorFrom(MakeSymbol("s"), MakeSymbol("re"), MakeSymbol("n"))),
			`Splits string on a regular expression. Returns vector of the splits.`, "1.0"))

	stringNamespace.InternVar("split-lines", split_lines_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Splits string on \n or \r\n. Returns vector of the splits.`, "1.0"))

	stringNamespace.InternVar("starts-with?", isstarts_with_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("substr"))),
			`True if s starts with substr.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Boolean"}))

	stringNamespace.InternVar("trim", trim_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Removes whitespace from both ends of string.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("trim-left", trim_left_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Removes whitespace from the left side of string.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("trim-newline", trim_newline_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Removes all trailing newline \n or return \r characters from string.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("trim-right", trim_right_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Removes whitespace from the right side of string.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("triml", triml_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Removes whitespace from the left side of string.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("trimr", trimr_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Removes whitespace from the right side of string.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("upper-case", upper_case_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Converts string to all upper-case.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

}

func init() {
	stringNamespace.Lazy = Init
}
