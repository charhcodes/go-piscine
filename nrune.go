package main

func NRune(s string, n int) rune {
	ConvRune := []rune(s)
	if n > ConvRune {
		return 0
	} else {
		n1 := ConvRune - n
		return ConvRune[len(ConvRune)-n1]
	}
}
