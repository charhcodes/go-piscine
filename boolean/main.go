package piscine

import (
	"os"

	"github.com/01-edu/z01"

	"piscine"
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

func TrimAtoi(s string) {
	counter := 0
	add := true
	number := false
	arr := []int{}
	for _, val := range s {
		if '0' <= val && val <= '9' {
			if val == '0' && !number {
				continue
			}
			arr = append(arr, int(val-48))
			number = true
			counter++
		} else if val == '-' && !number {
			add = false
		}
	}
	b := 0
	c := 1
	for i := counter - 1; i >= 0; i-- {
		b += arr[i] * c
		c *= 10
	}
	if !add {
		b *= -1
	}
}

func main() {
	a := os.Args[1:]
	lengthOfArg := piscine.TrimAtoi(a)
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
