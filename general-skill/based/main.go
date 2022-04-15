package main

// Please give the 01110100 01100101 01110011 01110100 as a word.

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scn = bufio.NewScanner(os.Stdin)
var sb strings.Builder

func main() {
	bi := flag.Bool("b", false, "convert your bytes to string")
	oct := flag.Bool("o", false, "convert your octal to string")
	hx := flag.Bool("h", false, "convert your hex to string")
	flag.Parse()

	switch {
	case *bi:
		fromBytes(&sb)
	case *oct:
		fromOct(&sb)
	case *hx:
		fromHex(&sb)
	}
	fmt.Println(sb.String())
}

func fromHex(sb *strings.Builder) *strings.Builder {
	for scn.Scan() {
		if scn.Text() == "" {
			break
		}
		for i := 0; i < len(scn.Text()); i += 2 {
			b, err := strconv.ParseInt(scn.Text()[i:i+2], 16, 0)
			if err != nil {
				panic(err)
			}
			sb.WriteString(fmt.Sprintf("%c", b))
		}
	}
	return sb
}

func fromOct(sb *strings.Builder) *strings.Builder {
	for scn.Scan() {
		if scn.Text() == "" {
			break
		}
		writeFieldWith(8)
	}
	return sb
}

func fromBytes(sb *strings.Builder) *strings.Builder {
	for scn.Scan() {
		if scn.Text() == "" {
			break
		}
		writeFieldWith(2)
	}
	return sb
}

func writeFieldWith(base int) {
		for _, s := range strings.Fields(scn.Text()) {
			b, err := strconv.ParseInt(s, base, 0)
			if err != nil {
				panic(err)
			}
			sb.WriteString(fmt.Sprintf("%c", b))
		}
}

