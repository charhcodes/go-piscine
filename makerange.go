package piscine

func MakeRange(min, max int) []int {
	size := make([]int, 10)
	for i := min; i < max; i++ {
		size[i] = i
	}
	return size
}
