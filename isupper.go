package piscine

func IsUpper(s string) bool {
	runeString := []rune(s)

	for i := 0; i < len(s); i++ {
		a := runeString[i]
		if a >= 0 && a < 65 {
			return false
		}
		if a >= 91 {
			return false
		}
	}
	return true
}
