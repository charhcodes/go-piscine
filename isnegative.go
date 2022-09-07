package main

import "github.com/01-edu/z01"

func main() {
	integer := -5

	if integer < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
}
