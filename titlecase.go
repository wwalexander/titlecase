package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const usage = `usage: titlecase string

titlecase maps all Unicode letters in string to their title case.`

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
	}
	flag.Parse()
	for _, arg := range flag.Args() {
		title := strings.Title(strings.ToLower(arg))
		fmt.Printf("%s ", title)
	}
	fmt.Println()
}
