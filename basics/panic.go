package basics

import "fmt"

func main() {

	// panic(interface{})
	process(-1)

}

func process(input int) {
	defer fmt.Println("This will always execute 1")
	defer fmt.Println("This will always execute 2")

	if input < 0 {
		fmt.Println("Before panic")
		panic("Input must be a non negative number")
		fmt.Println("After panic")
	}
	fmt.Println("inputt")
}