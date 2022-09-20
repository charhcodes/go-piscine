package main

import "github.com/01-edu/z01"

func main() {
	x := "x = 42, y = 21"
	for _, v := range x {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
