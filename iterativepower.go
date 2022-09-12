package piscine

func IterativePower(nb int, power int) int {
	if nb < 0 || nb > 20 {
		return 0
	} else {
		result := 1
		for power != 0 {
			result *= nb
			power -= 1
		}
		return result
	}
}
