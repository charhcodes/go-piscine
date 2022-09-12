package piscine

func IterativePower(nb int, power int) int {
	if nb > 20 || nb < 0 || power < 0 {
		return 1
	} else {
		result := 1
		for power != 0 {
			result *= nb
			power -= 1
		}
		return result
	}
	return 0
}
