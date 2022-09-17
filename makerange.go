package piscine

func MakeRange(min, max int) []int {
	size := make([]int, min)
	for i := min; i < max; i++ {
		size[i] = i
	}
	return size
}
