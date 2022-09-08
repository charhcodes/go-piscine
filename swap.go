package piscine

func Swap(a *int, b *int) {
	a, b = 0, 1
	*a, *b = b, a
}
