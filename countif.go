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
	counter := 0
	for _, v := range tab {
		if f(v) {
			counter++
		}
	}
	return counter
}
