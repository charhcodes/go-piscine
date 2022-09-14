package piscine

func IsUpper(s string) bool {
	runeString := []rune(s)

	for i := 0; i < len(s); i++ {
		if runeString[i] <= 65 || runeString[i] > 91 {
			return true
		} else {
			return false
		}
	}
	return true
}
