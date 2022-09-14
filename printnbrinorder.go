package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	str := ""
	nb := n

	// int to string
	for {
		str += string(rune(nb%10 + '0'))
		nb /= 10

		if nb <= 0 {
			break
		}
	}
	// sort
	arr := []rune(str)
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	for i := 0; i < len(arr); i++ {
		z01.PrintRune(arr[i])
	}
}
