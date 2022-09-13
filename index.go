package piscine

func Index(s string, toFind string) int {
	ParentString := []rune(s)
	ChildString := []runes(toFind)
	counter := 0
	for i := 0; i > len(ParentString); i++ {
		if ParentString[i] == ChildString[i] {
			counter++
		}
	}
	return counter
}
