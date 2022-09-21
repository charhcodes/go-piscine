package piscine

/*
func CountIf(f func(string) bool, tab []string) int {
	for _, v := range tab {
		if f(v) {
			return len(tab)
		}
	}
} */

func CountIf(f func(string) bool, tab []string) int {
	for _, v := range tab {
		if f(v) {
			return len(tab)
		}
	}
	return len(tab)
}
