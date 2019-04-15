package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(2)
	}
	in := os.Args[1]
	out := url.QueryEscape(in)
	fmt.Println(out)
	os.Exit(0)
}

const usage = `urlify

Usage:
    urlify INPUT`
