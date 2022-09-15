package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// program name
	// args1 := os.Args[0]

	// command line arguments
	args1 := os.Args[1:]
	runeArgs := []rune(args1[0])

	// print
	for i := 0; i < len(args1); i++ {
		z01.PrintRune(runeArgs[i], args1)
		z01.PrintRune('\n')
	}
}
