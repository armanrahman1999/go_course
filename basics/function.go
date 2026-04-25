package basics

import "fmt"

func main() {

	sum:= addNumbers(3, 5)
	fmt.Println(addNumbers(2, 4), sum)

	greet := func(){
		fmt.Println("Radd dee")
	}
	greet()

	result1 := applyOperation( 5, 5, addNumbers)
	fmt.Println("apply operations", result1)

	multiplier := createMultiplier(2)
	fmt.Println("Dree dree", multiplier(4))

}

func addNumbers(a int, b int) int{
	return a + b
}

func applyOperation (x int, y int, opt func(int , int) int) int{
	return opt(x, y)
}

func createMultiplier (factor int) func(int) int {
	return func( x int) int{
		return factor * x
	}
}