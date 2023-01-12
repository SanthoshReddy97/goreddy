package main

import (
	"fmt"
	"os"
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	if args[0] == "version" {
		fmt.Println("v0.0.1")
	}
	fmt.Fprintln(os.Stderr, "usage")
	return 2
}
