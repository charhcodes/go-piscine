package piscine

func Index(s string, toFind string) int {
	ParentString := []rune(s)
	ChildString := []rune(toFind)
	index := -1

	if s == 0 || toFind == 0 {
		return 0
	}
	for i := 0; i < len(ParentString); i++ {
		if ChildString[0] == ParentString[i] {
			index = i
			for j := 0; j < len(ChildString); j++ {
				if !(ChildString[j] == ParentString[i+j]) {
					index = -1
					break
				}
			}
		}
		if index != -1 {
			return index
		}
	}
	return index
}
