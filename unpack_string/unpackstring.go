package unpackString

import (
	"strings"
	"unicode"
)

var prevChar rune
var ret strings.Builder

// Распаковывает строку s
func UnpackString(s string) string {
	prevChar = 0
	isEscape := false
	ret = strings.Builder{}

	for _, r := range s {
		if r == '\\' && !isEscape {
			isEscape = true
			continue
		}

		if !unicode.IsDigit(r) || isEscape {
			writeCharIfNotEmpty()
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

	writeCharIfNotEmpty()

	return ret.String()
}

func writeCharIfNotEmpty() {
	if prevChar != 0 {
		ret.WriteRune(prevChar)
	}
}
