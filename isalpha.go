package piscine

func IsAlpha(s string) bool {
	runeString := []rune(s)

	for i := 0; i < len(s); i++ {
		a := runeString[i]
		if a >= 32 && a < 48 {
			return false
		}
		if a >= 58 && a < 65 {
			return false
		}
		if a >= 91 && a < 97 {
			return false
		}
		if a >= 123 {
			return false
		}
	}
	return true
}
