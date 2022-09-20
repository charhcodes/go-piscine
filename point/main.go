package main

import "github.com/01-edu/z01"

type point struct {
    x int
    y int
}

func main() {
    ptr := point{
        x: 1,
        y: 2,
    }

    ptr.setPoint()

    printStr("x = ")
    for , curr := range getDigits(ptr.x) {
        z01.PrintRune(rune(curr + 48))
    }
    printStr(", y = ")
    for , curr := range getDigits(ptr.y) {
        z01.PrintRune(rune(curr + 48))
    }
    z01.PrintRune('\n')
}

func printStr(s string) {
    for _, curr := range s {
        z01.PrintRune(rune(curr))
    }
}

func getDigits(n int) []int {
    if n < 10 {
        return []int{n}
    } else {
        return append(getDigits(n/10), (n % 10))
    }
}

func (d point) setPoint() {
    (d).x = 42
    (*d).y = 21
}