package piscine

func LastRune(s string) rune {
	ConvRune := []rune(s)
	return ConvRune[len(ConvRune)-1]
}
