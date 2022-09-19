package piscine

func SplitWhiteSpaces(s string) []string {
	emptyString := []string
	runeString := []rune(s)

	for i := 0; i < len(s); i++ {
		emptyString += runeString[i]
		if runeString[i] == ' ' || runeString[i] == '/' || runeString[i] == '	' {
			break
		} else if i < len(s) {
			return emptyString
		}
	} 
	return emptyString
}
  