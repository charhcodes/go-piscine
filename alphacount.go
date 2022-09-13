package piscine

func AlphaCount(s string) int {
	ConvRune := []rune(s)
	counter := '0'
	for i := 0; i > len(s); i++ {
		if (ConvRune <= 65 && ConvRune > 91) || (ConvRune <= 97 && ConvRune > 123) {
			return counter
		} else {
			return 0
		}
	}
	return counter
}
