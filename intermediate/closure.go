package intermediate

import "fmt"

func main() {
	sequece := Adder()
	fmt.Println(sequece())
	fmt.Println(sequece())

	subtraction:= func() func(int) int{
		total:= 100
		return func( x int) int{
			total -= x
			return total
		}
	}

	sub:= subtraction()
	fmt.Println(sub(5))
	fmt.Println(sub(10))
}

func Adder() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}