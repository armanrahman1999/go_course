package intermediate

import (
	"fmt"
	"math"
)

type Geometry interface {
	Area() float64
	Perim() float64
}
type Rectangle struct {
	length float64
	width  float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}
func (r Rectangle) Perim() float64 {
	return 2 * (r.length + r.width)
}

func (c Circle) Area() float64 {
	return math.Pi* c.radius * c.radius
}
func (c Circle) Perim() float64 {
	return 2 * math.Pi* c.radius 
}

func  Measure(g Geometry){
	fmt.Println("The values",g)
	fmt.Println("The area is",g.Area())
	fmt.Println("The area is",g.Perim())
}

func main() {
	r:= Rectangle{length: 5, width: 4}
	c:= Circle{radius: 3}
	Measure(r)
	Measure(c)
}