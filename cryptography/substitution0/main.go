package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const (
  upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
  lower = "abcdefghijklmnopqrstuvwxyz"
)

func main() {
	f, err := os.Open("message.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scr := bufio.NewScanner(f)
	scr.Scan()
	upkey, lowkey := scr.Text(), strings.ToLower(scr.Text())
	var sb strings.Builder
	for scr.Scan() {
		for _, r := range scr.Text() {
			switch {
			case unicode.IsUpper(r):
				sb.WriteByte(upper[strings.IndexRune(upkey, r)])
			case unicode.IsLower(r):
				sb.WriteByte(lower[strings.IndexRune(lowkey, r)])
			default:
				sb.WriteRune(r)
			}
		}
	}
	text := strings.Split(sb.String(), " ")
	fmt.Println(text[len(text)-1])
}
