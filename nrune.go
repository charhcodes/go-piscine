package piscine

func NRune(s string, n int) rune {
	ConvInt := []int(s)
	if n < 0 || n > ConvInt {
		return 0
	} else {
		ConvRune := []rune(s)
		newN = ConvInt - n
		return ConvRune[ConvInt-newN]
	}
}
