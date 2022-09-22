package piscine

func CollatzCountdown(start int) int {
	startSlice := []int{start}
	counter := 0
	for i := start; i < start; i-- {
		if start <= 0 {
			return -1
		} else if start % 2 = 0 {
			counter++
			startSlice[i] = startSlice[i] / 2
		} else if start % 2 = 1 {
			counter++
			startSlice[i] = startSlice[i]*3 + 1
		} else if start[i] == 1 {
			return
		}
	}
	return counter
}
