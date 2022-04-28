package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("message.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, 64)
	n, err := f.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, string(buf))
	fmt.Println(decryptRailFence(string(buf[:n]), 4))
}

// decryptRailFence takes in a cipher and outputs a decrypted string based on
// the number of supplied rails.
func decryptRailFence(cipher string, rails int) string {
	// initialize rail fence as matrix
	rail := make([][]byte, rails)
	for i := 0; i < rails; i++ {
		rail[i] = make([]byte, len(cipher))
		for j := 0; j < len(cipher); j++ {
			rail[i][j] = '.'
		}
	}
	for _, bb := range rail {
		fmt.Println(string(bb))
	}
	fmt.Println()

	// Mark where to make rail fence
	goDown := false
	row, col := 0, 0
	for range cipher {
		switch row {
		case 0:
			goDown = true
		case rails - 1:
			goDown = false
		}
		rail[row][col] = '*'
		col++
		if goDown {
			row++
		} else {
			row--
		}
	}
	for _, bb := range rail {
		fmt.Println(string(bb))
	}
	fmt.Println()

	// Fill in rail fence bytes from markers
	idx := 0
	for i := 0; i < rails; i++ {
		for j := 0; j < len(cipher); j++ {
			if rail[i][j] == '*' && idx < len(cipher) {
				rail[i][j] = cipher[idx]
				idx++
			}
		}
	}
	for _, bb := range rail {
		fmt.Println(string(bb))
	}
	fmt.Println()

	// Build decrypted string
	var res strings.Builder
	row, col = 0, 0
	for range cipher {
		switch row {
		case 0:
			goDown = true
		case rails - 1:
			goDown = false
		}
		if rail[row][col] != '*' {
			res.WriteByte(rail[row][col])
			col++
		}
		if goDown {
			row++
		} else {
			row--
		}
	}
	return res.String()
}
