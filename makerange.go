package piscine

func MakeRange(min, max int) []int {
	if max-min <= 0 {
		return nil
	} else {
		size := make([]int, max-min)
		for i := min; i < max; i++ {
			size[i-min] = i
		}
		return size
	}
}
