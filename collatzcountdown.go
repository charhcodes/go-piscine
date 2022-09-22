package piscine

func CollatzCountdown(start int) int {
	counter := 0
	for start != 1 {
		if start <= 0 {
			return -1
		} else if start%2 == 0 {
			start = start / 2
		} else if start%2 == 1 {
			start = (start * 3) + 1
		}
		counter++
	}
	return counter
}
