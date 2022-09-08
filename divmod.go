package piscine

func DivMod(a int, b int, div *int, mod *int) {
	a = 3
	b = 4
	d := 2

	c := a + b
	e := c / d
	f := c % d

	*div = e
	*mod = f
}
