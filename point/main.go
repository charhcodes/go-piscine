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
}
*/

package piscine

import "fmt"

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
}
