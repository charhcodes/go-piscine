package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	parameters := os.Args
	counter := 0

	for index := range parameters {
		counter = index
	}

	for i := counter; i >= 1; i-- {
		for _, word := range parameters[i] {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
