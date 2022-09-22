package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

/*
func main() {
	for len(os.Args) == 2 { // works only when terminal input = 2
		numbers, err := strconv.Atoi(os.Args[1]) // strconv takes terminal input (minus program name), returns numbers + error
		if err != nil {                          // if error status is true aka if error status is not zero (nil acts as a non-bool false)
			panic(err) // panic(err) checks for errors, basically cancelling the program
		}
		newnum := numbers  // reassigning numbers, keeps old numbers value
		counts := 0        // counter
		for numbers != 1 { // when numbers not equal to 1
			if newnum%2 != 0 { // if remainder of newnum not equal to 0
			} else {
				newnum = newnum / 2 // newnum divided by 2
			}
			counts++ // counter up (after if else statements)
		}
		var x int = 2 ^ counts // x = 2 to the power of counter
		if x == numbers {      // if x equals numbers
			return
		}
	}
}
*/

func main() {
	for len(os.Args) == 2 {
		number, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		test := IsPowerOfTwo(number)
		if test == 0 {
			PrintStr("true\n")
			break
		} else {
			PrintStr("false\n")
			break
		}
	}
}

func IsPowerOfTwo(n int) int {
	if n == 0 {
		return 1
	}
	return n & (n - 1)
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
