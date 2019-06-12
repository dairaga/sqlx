package sqlx

import "strings"

func toCamel(v string) string {
	ret := &strings.Builder{}
	ret.Grow(len(v))

	needUpper := true

	for _, x := range v {

		if needUpper {
			if 'a' <= x && x <= 'z' {
				x -= 'a' - 'A'
			}
		}
		if x == '_' || x == '.' {
			needUpper = true
			continue
		}

		ret.WriteByte(byte(x))
		needUpper = false
	}

	return ret.String()
}

func toSnake(v string) string {
	needUnderLine := false

	ret := &strings.Builder{}
	ret.Grow(len(v))

	for _, x := range v {
		if 'A' <= x && x <= 'Z' {
			if needUnderLine {
				ret.WriteByte(byte('_'))
			}
			needUnderLine = true
			x += 'a' - 'A'
		}
		ret.WriteByte(byte(x))
	}
	return ret.String()
}
