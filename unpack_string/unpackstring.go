package unpackString

import (
	"strings"
	"unicode"
)

// Распаковывает строку s
func UnpackString(s string) string {
	var prevChar rune = 0
	ret := strings.Builder{}
	isEscape := false

	for _, r := range s {
		if r == '\\' && !isEscape {
			isEscape = true
			continue
		}

		if !unicode.IsDigit(r) || isEscape {
			if prevChar != 0 {
				ret.WriteRune(prevChar)
			}
			prevChar = r
		} else {
			if unicode.IsDigit(r) && prevChar != 0 {
				for i := 0; i < int(r-'0'); i++ {
					ret.WriteRune(prevChar)
				}
				prevChar = 0
			}
		}

		isEscape = false
	}

	if prevChar != 0 {
		ret.WriteRune(prevChar)
	}

	return ret.String()
}
