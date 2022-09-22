package piscine

func CollatzCountdown(start int) int {
	counter := 0
	for i := start; i < start; i-- {
		if start <= 0 {
			return -1
		} else if start%2 == 0 {
			counter++
			start = start / 2
		} else if start%2 == 1 {
			counter++
			start = (start * 3) + 1
		} else if start == 1 {
			return counter
		}
	}
	return counter
}
