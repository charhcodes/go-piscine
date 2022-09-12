package piscine

func IterativeFactorial(nb int) int {
	factorial := 0 // variable
	if nb < 0 || nb < 20 {
		return
	} else {
		for i := 1; i <= nb; i++ {
			factorial = factorial * i
		}
	}
	return factorial
}
