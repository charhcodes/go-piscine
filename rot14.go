package piscine

func Rot14(s string) string {
	runeString := []rune(s)

	for i := 0; i < len(s); i++ {
		// A - L inclusive
		if (runeString[i] >= 65 && runeString[i] <= 76) || (runeString[i] >= 97 && runeString[i] <= 108) {
			runeString[i] += 14
		} else if (runeString[i] > 76 && runeString[i] <= 90) || (runeString[i] > 108 && runeString[i] <= 122) {
			runeString[i] -= 12
		}
	}
	return string(runeString)
}
