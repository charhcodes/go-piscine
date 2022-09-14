package piscine

func ToUpper(s string) string {
	ToUpperRune := []rune(s)
	for i := 0; i < len(s); i++ {
		if ToUpperRune[i] <= 97 && ToUpperRune > 123 {
			ToUpperRune -= 32
		}
		return
	}
}
