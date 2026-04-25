package basics

import "fmt"

func main() {

	// var sliceNmae[]ElementType
	slice := make([]int, 5)
	fmt.Println(slice)
	//lets make a copy of the slice

	pointArray := [5]int{1,2,3,4,5}

	slice1 := pointArray[1: 4]
	slice1 = append(slice1, 5,6,7)
	fmt.Println(slice1)

	copySlice := make([]int, len(slice1))
	copy(copySlice, slice1)
	fmt.Println("copy slice", copySlice)
	

}