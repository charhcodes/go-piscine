package piscine

func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		for i = min; i < max; i++ {
			a[i] = min + i
		}
	}
	return a
}
