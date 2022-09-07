package main

import "github.com/01-edu/z01"

func main() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			z01.PrintRune(a)
			z01.PrintRune(b)
			z01.PrintRune(44)
			z01.PrintRune(32)
			if a == b {
				break
			}
		}
	}
}
