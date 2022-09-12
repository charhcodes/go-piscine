package piscine

func IterativePower(nb int, power int) int {
	if power < 0 || nb == 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if nb < 0 {
		return nb * IterativePower(nb, power-1)
	}
	return 0
}
