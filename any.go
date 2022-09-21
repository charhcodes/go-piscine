// returns true for a string slice []string
// when passed through f function

package piscine

func Any(f func(string) bool, a []string) bool {
	for _, v := range a {
		if f(v) {
			return true
		}
	}
	return false
}
