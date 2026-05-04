package intermediate

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r *Rectangle) Scale(factor float64){
	r.length *= factor
	r.width *= factor
}

func main() {
	r := Rectangle{
		length: 10,
		width:  10,
	}

	fmt.Println("Area before scale",r.Area())
	r.Scale(5)
	fmt.Println("Area after scale",r.Area())

}