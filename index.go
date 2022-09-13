package piscine

func Index(s string, toFind string) int {
	ParentString := []rune(s)
	ChildString := []rune(toFind)
	counter := 0

	if ParentString > '0' && ChildString > '0' {
		for i := 0; i < len(ParentString); i++ {
			if ParentString[i] != ChildString[i] {
				continue
			}
			if ParentString[i] == ChildString[i] {
				counter++
			} else if ParentString[i] != ChildString[i] {
				return counter
			}
		}
		if counter == 0 {
			return -1
		}
	}
	return counter
}
