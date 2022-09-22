package piscine

import "github.com/01-edu/z01"

func Rot14(s string) string {
	runeString := []rune(s)
	emptyString := ""

	for i := 0; i < len(s); i++ {
		if (runeString[i] >= 65 && runeString[i] <= 76) || (runeString[i] >= 97 && runeString[i] <= 108) {
			runeString[i] += 14
		}
		if (runeString[i]) < 76 && runeString[i] > 90) || (runeString[i] < 108 && runeString[i] > 122) {
			runeString[i] += 14
			runeString[i] -= 26
		} 
		if runeString[i] == 90 {
			runeString[i] == 78
		}
		if runeString[i] == 122 {
			runeString[i] == 110
		}
	}
	return emptyString += runeString[i]
}
