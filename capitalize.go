package piscine

func Capitalize(s string) string {
	StringRune := []rune(s)
	converter := 'A' - 'a'

	if len(StringRune) > 0 && IsLower(string(StringRune[0])) {
		StringRune[0] = StringRune[0] + converter
	}

	for i := 1; i < len(StringRune); i++ {
		if StringRune[i-1] == 32 || !IsAlpha(string(StringRune[i-1])) {
			if IsLower(string(StringRune[i])) {
				StringRune[i] = StringRune[i] + converter
			}
		} else {
			if IsUpper(string(StringRune[i])) {
				StringRune[i] = StringRune[i] - converter
			}
		}
	}
	return string(StringRune)
}
