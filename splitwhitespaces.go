package piscine

/*func SplitWhiteSpaces(s string) []string {
	stringSlice := []string(s)
	emptyString := []string{}

	for i := range s {
		emptyString = append(emptyString, stringSlice...)
		if s[i] == ' ' || s[i] == '/' || s[i] == '	' {
			break
		} else if i < len(s) {
			return emptyString
		}
	}
	return emptyString
}
*/

func SplitWhiteSpaces(s string) []string {
	var word []rune
	var str []string
	for i, val := range s {
		if val != 32 {
			word = append(word, val)
		}
		if val == 32 || i == len(s)-1 {
			if len(word) > 0 {
				str = append(str, string(word))
				word = []rune{}
			}
		}
	}
	return str
}
