package piscine

func Compare(a, b string) int {
	if len(a) == len(b) {
		return 0
	}
	if len(a) > len(b) {
		return -1
	}
	if len(a) < len(b) {
		return 1
	}
}
