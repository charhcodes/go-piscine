package piscine

func Abort(a, b, c, d, e int) int {
	medSlice := []int{a, b, c, d, e}
	var newSlice []int

	for i := 0; i < len(medSlice); i++ {
		for j := i; i < len(medSlice); j++ {
			if medSlice[i] > medSlice[j] {
				z := medSlice[i]
				medSlice[i] = medSlice[j]
				medSlice[j] = z
			}
			newSlice = append(medSlice)
		}
	}
	return newSlice[2]
}
