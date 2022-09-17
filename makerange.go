package piscine

func MakeRange(min, max int) []int {
	var size []int
	size = make([]int, 10)
	for i := min; i < max; i++ {
		size[i] = 1
	}
	return size
}
