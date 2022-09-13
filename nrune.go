package piscine

func NRune(s string, n int) rune {
	if n < 0 || n > len(s) {
		return 0
	} else {
		ConvRune := []rune(s)
		newN := len(s) - n
		return ConvRune[len(s)-newN]
	}
}
