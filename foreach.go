package main

func one(nbr int) int {
	return nbr + 1
}

func two(nbr int) int {
	return nbr + 2
}

func three(nbr int) int {
	return nbr + 3
}

func four(nbr int) int {
	return nbr + 4
}

func five(nbr int) int {
	return nbr + 5
}

func six(nbr int) int {
	return nbr + 6
}

func ForEach(f func(int), a []int) {
	a = []func(int) int{one, two, three, four, five, six}
	return a
}
