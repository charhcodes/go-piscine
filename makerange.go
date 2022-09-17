package piscine

func MakeRange(min, max int) []int {
	size := make([]int, max-min)
	for i := range size {
		size[i] = min + i
	}
	return size
}
