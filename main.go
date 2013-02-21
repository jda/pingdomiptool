package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s TEMPLATE.tmpl OUTPUT\n", os.Args[0])
	}
}