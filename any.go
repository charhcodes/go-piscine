package main

func Any(f func(string) bool, a []string) bool {
	for i, v := range a {
		if f(v) {
			return true
		}
	}
}
