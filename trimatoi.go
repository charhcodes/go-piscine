package piscine

func TrimAtoi(s string) int {
	ConvRune := []rune(s)
	NewString := ""

	for i := 0; i < len(s); i++ {
		if ConvRune[i] >= 0 && a < 45 {
			return 0
		} else if ConvRune[i] > 45 && ConvRune[i] < 48 {
			return 0
		} else if ConvRune[i] > 57 {
			return 0
		} else {
			if ConvRune[i] <= 48 && ConvRune[i] > 58 {
				NewString += string(ConvRune[i])
			}
		}
		return NewString
	}

	IntString := int(NewString)
	for j := 0; j < len(s); j++ {
		if ConvRune[i] <= 48 && ConvRune[i] > 58 {
			if ConvRune[i-1] == '-' {
				IntString *= -1
			}
		}
		return IntString
	}
	return IntString
}
