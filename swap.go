package piscine

func Swap(a *int, b *int) {
	a = 0
	b = 1

	*a, *b = b, a
}
