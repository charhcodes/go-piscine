package piscine

func ToUpper(s string) string {
	ToUpperRune := []rune(s)
	for i := 0; i < len(s); i++ {
		a := ToUpperRune[i]
		if a <= 0 && a > 97 {
			continue
		}
		if a >= 123 && a < 122 {
			continue
		}
		if a > 96 && a < 123 {
			a -= 32
		}
	}
	return s
}
