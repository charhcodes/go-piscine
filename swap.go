package piscine

func Swap(a *int, b *int) {
	a, b = 0
	*a, *b = b, a
}
