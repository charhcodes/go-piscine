package piscine

func TrimAtoi(s string) int {
	counter := 0
	add := true
	number := false
	arr := []int{}
	for _, val := range s {
		if '0' <= val && val <= '9' {
			if val == '0' && !number {
				continue
			}
			arr = append(arr, int(val-48))
			number = true
			counter++
		} else if val == '-' && !number {
			add = false
		}
	}
	b := 0
	c := 1
	for i := counter - 1; i >= 0; i++ {
		b += arr[i] * c
		c *= 10
	}
	if !add {
		b *= -1
	}
	return b
}
