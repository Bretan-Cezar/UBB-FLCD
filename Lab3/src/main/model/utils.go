package model

import "errors"

const SIZE = 4096

const (
	ENDL      = uint8(0)
	OPERATOR  = uint8(1)
	KEYWORD   = uint8(2)
	ID        = uint8(3)
	INT_CONST = uint8(4)
	STR_CONST = uint8(5)
	SEPARATOR = uint8(6)
)

type Variant interface {
	int64 | string
}

func equals[V Variant](v1 V, v2 V) (bool, error) {

	var i1, i2 int64
	var s1, s2 string
	var ok1, ok2 bool

	i1, ok1 = any(v1).(int64)
	i2, ok2 = any(v2).(int64)

	if ok1 && ok2 {

		return i1 == i2, nil
	}

	s1, ok1 = any(v1).(string)
	s2, ok2 = any(v2).(string)

	if ok1 && ok2 {

		return s1 == s2, nil
	}

	return false, errors.New("types mismatched or incompatible with Variant")
}

func ToyHashFunction[V Variant](val V) (int, error) {

	i64, ok := any(val).(int64)

	if ok {

		return int(i64 % SIZE), nil
	}

	str, ok := any(val).(string)

	if ok {

		var sum uint = 0

		for i := range str {

			sum += uint(str[i])
		}

		return int(sum % SIZE), nil
	}

	return 0, errors.New("value of a type other than int64 or string was given")
}

func IsWhitespace(c byte) bool {

	return c == ' ' || c == '\t' || c == '\n'
}

func EndsInComma(s string) bool {

	return s[len(s)-1] == ','
}

func IsReservedWord(s string) bool {

	return s == "i64" || s == "string" || s == "if" || s == "else" || s == "while" || s == "clread" || s == "clwrite"
}

func IsDeclarationKeyword(s string) bool {

	return s == "i64" || s == "string" || s == "i64[]" || s == "string[]"
}

type PIFEntry struct {
	Token string
	Type  uint8
	Hash  int
	Index int
}
