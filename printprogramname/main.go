package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	runeSlice := []rune(arguments[0])
	result := make([]rune, 0)

	for i := len(runeSlice) - 1; i >= 0; i-- {
		if runeSlice[i] != 47 {
			result = append(result, runeSlice[i])
		} else {
			break
		}
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	for i := 0; i < len(result); i++ {
		z01.PrintRune(result[i])
	}
	z01.PrintRune('\n')
}
