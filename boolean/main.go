package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	nbr -= 1
	evenNbr := nbr % 2
	if evenNbr == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	a := os.Args[1:]
	lengthOfArg := TrimAtoi(a)
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
