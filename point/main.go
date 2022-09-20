/*package piscine

import (
	"fmt"
)

type PointR struct {
	ptr *int
}

var x *PointR
var y *PointR

func setPoint(ptr *int) {
	x = PointR{ptr: ptr}
	y = PointR{ptr: ptr}
}

func main() {
	points := &ptr{}
	setPoint(points)

	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}*/

/*package piscine

import (
	"fmt"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) *point {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}*/

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