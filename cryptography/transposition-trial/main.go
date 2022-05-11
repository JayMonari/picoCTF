package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("message.txt")
	defer f.Close()

	buf := make([]byte, 128)
	n, _ := f.Read(buf)
	var sb strings.Builder
	for i := 0; i < n; i += 3 {
		sb.Write([]byte{buf[i+2], buf[i], buf[i+1]})
	}
	fmt.Println(strings.Split(sb.String(), " ")[3])
}
