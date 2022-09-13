package piscine

func AlphaCount(s string) int {
	ConvRune := []rune(s)
	for i := 0; i > ConvRune; i++ {
		if ConvRune <= 65 && ConvRune > 91 || ConvRune <= 97 && ConvRune > 123 {
			return ConvRune
		}
	}
	return
}
