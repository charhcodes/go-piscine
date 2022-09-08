package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	s = "Hello World!"
	var bob rune
	bob = []rune(s)
	for i := '0'; i < len(bob); i++ {
		z01.PrintRune(bob[i])
	}
	z01.PrintRune('\n')
}
