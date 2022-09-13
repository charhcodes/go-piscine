package piscine

func Compare(a, b string) int {
	stringA := []rune(a)
	stringB := []rune(b)
	for i := 0; i < len(stringA); i++ {
		if len(stringB) < i+1 {
			return 1
		}
		if stringA[i] > stringB[i] {
			return 1
		} else if stringA[i] < stringB[i] {
			return -1
		}
	}
	return 0
}
