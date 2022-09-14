package piscine

func IsLower(s string) bool {
	runeString := []rune(s)

	for i := 0; i < len(s); i++ {
		a := runeString[i]
		if a >= 0 && a <= 96 {
			return false
		}
		if a >= 123 {
			return false
		}
	}
	return true
}
