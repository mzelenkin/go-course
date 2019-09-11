package unpackstring

import (
	"testing"
)

func TestUnpackString(t *testing.T) {
	variants := map[string]string{
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
			t.Errorf("%s = %s", param, ret)
			t.Errorf("не работает, ожидалось %s, получено %s", ethalon, ret)
		}
	}
}
