package piscine

func IsNumeric(s string) bool {
	runeString := []rune(s)

	for i := 0; i < len(s); i++ {
		a := runeString[i]
		if a >= 0 && a < 32 {
			return false
		}
		if a >= 127 {
			return false
		}
	}
	return true
}
