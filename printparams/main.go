package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// program name
	// args1 := os.Args[0]
	// command line arguments
	args1 := os.Args

	// print
	for index, value := range args1 {
		if index != 0 {
			for _, value := range value {
				z01.PrintRune(value)
			}
			z01.PrintRune('\n')
		}
	}
}
