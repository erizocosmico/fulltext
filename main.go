package main

import (
	"fmt"
	"os"
	"strings"
	
	"golang.org/x/text/width"
)

func main() {
	text := strings.Join(os.Args[1:], " ")
	fmt.Println(width.Widen.String(text))
}
