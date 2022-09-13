package piscine

func FirstRune(s string) rune {
	s = "Hello"
	ConvRune := []rune(s)
	ConvString := string(ConvRune[0:1])
	return ConvString
}
