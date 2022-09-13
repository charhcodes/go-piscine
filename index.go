package piscine

func Index(s string, toFind string) int {
	ParentString := []rune(s)
	ChildString := []rune(toFind)
	counter := 0

	/*if ParentString > '0' && ChildString > '0' {
		for i := 0; i < len(ParentString); i++ {
			if ChildString[0] == ParentString[i] {
				for j := 0; len(ChildString); j++ {
					if ChildString[j] == ParentString[i+j] {
						counter++
					} else {
						return counter
					}
				}
			} else {
				return -1
			}
		}
	}*/
	for i := 0; i < len(ParentString); i++ {
		if ChildString[0] == ParentString[i] {
			counter++
			for j := 0; len(ChildString); j++ {
				if ChildString[j] == ParentString[i+j] {
					counter++
				} else {
					return counter
				}
			}
		} else {
			return -1
		}
	}
}
