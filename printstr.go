package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	s = []string{"Hello World!"}
	bob := rune(s)
	for i := 0; i < len(bob); i++ {
		z01.PrintRuneell(bob[i])
	}
	z01.PrintRune('\n')
}
