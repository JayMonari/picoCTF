package main

import (
	"fmt"
	"strings"
	"unicode"
)

const rot13 = "cvpbPGS{P7e1S_54I35_71Z3}"

func main() {
	var sb strings.Builder
	var dec rune
	for _, enc := range rot13 {
		switch {
		case unicode.IsUpper(enc):
			dec = rotate13(enc, 'A', 'Z')
		case unicode.IsLower(enc):
			dec = rotate13(enc, 'a', 'z')
		default:
			dec = enc
		}
		sb.WriteRune(dec)
	}
	fmt.Println(sb.String())
}

// rotate13 is the OG caeser cipher
func rotate13(r, start, end rune) rune {
	cp := r + 13
	switch {
	case cp < start:
		return cp + end - start + 1
	case cp > end:
		return cp%end + start - 1
	default:
		return cp
	}
}
