package unpackstring

import (
	"testing"
)

func TestUnpackString(t *testing.T) {
	variants := map[string]string{
		"я5Му2":    "яяяяяМуу",
		"a2":       "aa",
		"a2b5":     "aabbbbb",
		"a4bc2d5e": "aaaabccddddde",
		"abcd":     "abcd",
		"45":       "",
		`qwe\4\5`:  `qwe45`,
		`qwe\45`:   `qwe44444`,
		`qwe\\5`:   `qwe\\\\\`,
	}

	for param, ethalon := range variants {
		if ret := UnpackString(param); ret != ethalon {
			t.Errorf("Передано '%s', ожидалось '%s', получено '%s'", param, ethalon, ret)
		}
	}
}
