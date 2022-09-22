package piscine

func CollatzCountdown(start int) int {
	for i := start; i < len(start); i-- {
		if start <= 0 {
			return -1
		} else if start[i] % 2 = 0 {
			start[i] = start[i] / 2  
		} else {
			start[i] = start[i] * 3 + 1 
		}
	}
	return i
}
