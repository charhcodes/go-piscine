package piscine

func MakeRange(min, max int) []int {
	size := make([]int, 1)
	for i := min; i < max; i++ {
		size[i] = i
	}
	return size
}
