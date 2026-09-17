package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run . <username>")
		os.Exit(2)
	}

	fmt.Printf("Hello, %s! Welcome to Go.\n", os.Args[1])
}
