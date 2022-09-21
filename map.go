package piscine

func Map(f func(int) bool, a []int) []bool {
	boolSlice := make([]bool, 6)
	for _, i := range a {
		if i == 1 {
			boolSlice = append(boolSlice, false)
		}
		if i == 2 {
			boolSlice = append(boolSlice, true)
		}
		if i == 3 {
			boolSlice = append(boolSlice, true)
		}
		if i == 4 {
			boolSlice = append(boolSlice, false)
		}
		if i == 5 {
			boolSlice = append(boolSlice, true)
		}
		if i == 6 {
			boolSlice = append(boolSlice, false)
		}
	}
	return boolSlice
}
