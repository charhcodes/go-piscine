package piscine

func Abort(a, b, c, d, e int) int {
	medSlice := []int{a, b, c, d, e}

	for i := 0; i < len(medSlice); i++ {
		for j := i; i < len(medSlice); j++ {
			if medSlice[i] > medSlice[j] {
				medSlice[i], medSlice[j] = medSlice[j], medSlice[i]
			}
		}
	}
	return medSlice[2]
}