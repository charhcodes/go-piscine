package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args1 := os.Args
	runes := []rune(args1[1:])
	emptyRunes := make([]rune, 0)

	// print
	for i, value := range args1 {
		if i != 0 {
			for _, value := range value {
				emptyRunes = append(emptyRunes, value)
			}
		}
	}
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	for i := 0; i < len(emptyRunes); i++ {
		z01.PrintRune(emptyRunes[i])
	}
	z01.PrintRune('\n')
}
