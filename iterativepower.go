package piscine

func IterativePower(nb int, power int) int {
	if power < 0 || nb == 0 {
		return 0
	} else if nb < 0 {
		result := 1
		for power != 0 {
			result *= nb
			power -= 1
		}
		return (result * -1)
	} else {
		result := 1
		for power != 0 {
			result *= nb
			power -= 1
		}
		return result
	}
}
