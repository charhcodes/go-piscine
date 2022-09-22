package piscine

func Abort(a, b, c, d, e int) int {
	medSlice := []int{a, b, c, d, e}
	SortSlice(medSlice)
	return medSlice[2]
}

func SortSlice(medSlice []int) {
	for i := 0; i < len(medSlice); i++ {
		for j := 0; j < len(medSlice); j++ {
			if medSlice[j] > medSlice[i] {
				medSlice[i], medSlice[j] = medSlice[j], medSlice[i]
			}
		}
	}
}
