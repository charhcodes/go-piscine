package piscine

func ToUpper(s string) string {
	ToUpperRune := []rune(s)
	for i := 0; i < len(s); i++ {
		a := ToUpperRune[i]
		if !(a <= 0 && a > 97) {
			a -= 32
		}
		if !(a >= 123 && a < 122) {
			a -= 32
		}
	}
	return s
}