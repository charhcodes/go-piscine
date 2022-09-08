package piscine

func DivMod(a int, b int, div *int, mod *int) {
	a = 4
	b = 6
	c := a + b
	*div = c / 2
	*mod = c % 2
}
