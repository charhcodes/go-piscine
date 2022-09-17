package piscine

func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = i
		if i == max {
			break
		}
	}
	return a
}
