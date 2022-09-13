package piscine

func AlphaCount(s string) int {
	ConvRune := []rune(s)
	counter := 0
	for _, letters := range ConvRune {
		if letters >= 'a' && letters <= 'z' || letters >= 'A' && letters <= 'Z' {
			counter++
		}
	}
	return counter
}
