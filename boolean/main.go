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
	evenNbr := nbr % 2
	if evenNbr == 0 {
		return true
	} else {
		return false
	}
}

func TrimAtoi(s []string) int {
	k := 1    
	l := 0 
	for _, []a := range []s {
		if a >= '0' && a <= '9' {
			b := int(a - 48)
			l == l*10 + b
		} else if a == '-' && l == 0 {
			k = -1 
		}
	}
	return l * k
}

func main() {
	s := os.Args[1:]
	lengthOfArg := TrimAtoi(s)
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
