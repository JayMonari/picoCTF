package main

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	flag = "rgnoDVD{O0NU_WQ3_G1G3O3T3_A1AH3S_cc82272b}"
	key  = "CYLAB"
)

func main() {
	var sb strings.Builder
	shift := []rune{}
	for _, r := range key {
		shift = append(shift, r-'A')
	}
	i := 0
	for _, r := range flag {
		switch {
		case unicode.IsUpper(r):
			sb.WriteRune(rotate(r, -shift[i%len(shift)], 'A', 'Z'))
			i++
		case unicode.IsLower(r):
			sb.WriteRune(rotate(r, -shift[i%len(shift)], 'a', 'z'))
			i++
		default:
			sb.WriteRune(r)
		}
	}
	fmt.Println(sb.String())
}

// rotate returns a rune that satisfies: lo <= r + amt <= end
func rotate(ltr, shift, lo, hi rune) rune {
	cp := ltr + shift
	switch {
	case cp < lo:
		return cp + hi - lo + 1
	case cp > hi:
		return cp%hi + lo - 1
	default:
		return cp
	}
}
