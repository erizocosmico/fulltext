package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	text := strings.Join(os.Args[1:], " ")

	for _, c := range text {
		if c == 0x20 {
			fmt.Printf("%c", 0x3000)
		} else if c >= 0x21 && c <= 0x7E {
			fmt.Printf("%c", c+0xFEE0)
		} else {
			fmt.Printf("%c", c)
		}
	}

	fmt.Println()
}
