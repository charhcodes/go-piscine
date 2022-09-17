package piscine

func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		for i = min; i < max; i++ {
			a[i] = min + i
		} else if max == 9223372036854775807 {
			return nil
		} else if min <= max {
			return nil
		}
	}
	return a
}
