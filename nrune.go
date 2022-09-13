package piscine

func NRune(s string, n int) rune {
	ConvRune := []rune(s)
	if n > ConvRune || n < 0 {
		return 0
	} else {
		n1 := ConvRune - n
		return ConvRune[len(ConvRune)-n1]
	}
}
