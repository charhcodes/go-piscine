package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	bob := []rune(s)

	for i := 0; i < len(bob); i++ {
		z01.PrintRune(bob[i])
	}
}
