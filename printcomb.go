package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for a := 0; a <= 9; a++ {
		for b := a + 1; b <= 9; b++ {
			for c := b + 1; c <= 9; c++ {
				print(a)
				print(b)
				print(c)
				if a < 55 {
					z01.PrintRune(44)
					z01.PrintRune(32)
				}
			}
		}
	}
}