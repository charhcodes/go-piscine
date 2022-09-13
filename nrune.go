package piscine

func NRune(s string, n int) rune {
	ConvRune := []rune(s)
	counter := 0
	for n := range ConvRune {
		counter++
		n1 := ConvRune[range-n]
		return ConvRune[n1]
	}
}
