package piscine

func CountIf(f func(string) bool, tab []string) int {
	for _, v := range tab {
		if f(v) {
			return len(tab)
		}
	}
}
