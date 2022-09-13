package piscine

import "fmt"

func main() {
	sentence := []byte("Hello")

	counter := 0

	for index, letter := range sentence {
		counter++
		fmt.Printf("Index: %v Letter: %c\n", index, letter)
	}
	fmt.Printf("Counter value: %v\n", counter)
}

// n rune
func NRune(s string, n int) rune {
	ConvInt := []int(s)
	if n < 0 || n > ConvInt {
		return 0
	} else {
		ConvRune := []rune(s)
		n1 = ConvInt - n
		return ConvRune[n1]
	}
}

// alpha count
func main() {
	ConvRune := []rune(s)
}
