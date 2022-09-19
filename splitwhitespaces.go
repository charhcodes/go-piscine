package piscine

func SplitWhiteSpaces(s string) []string {
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
