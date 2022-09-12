package piscine

func RecursiveFactorial(nb int) int {
	factorial := 1
	if nb < 0 || nb > 20 {
		return 0
	} else {
		RecursiveFactorial(factorial*nb - 1)
		if nb == 1 {
			return factorial
		}
	}
	return factorial
}
