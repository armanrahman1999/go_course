package basics

import "fmt"
func main() {

	var numbers[5]int
	fmt.Println(numbers)

	numbers[0] = 1
	fmt.Print(numbers)

	fruits := [4]string{"Apple", "Banana", "Orange", "Mango"}

	fmt.Println(fruits)

	for i:=0; i<= len(fruits) -1 ; i++{
		fmt.Println("The fruits are :",fruits[i] )
	}
	//How do u define an array
	// numbers := [4]int{1,2,3,4}	
	// to initializ e this just write var numbers[4]int

	for index, value := range fruits{
		fmt.Println("The index is", index , "and the value is",value )
	}
	a, _:= someFunction()
	fmt.Print(a)
	//we use undersore when we dont want to use that variable similary in for in range we can just add _ in index if we dont need it


	array1 := [3]int{1,2,3}
	var array2 *[3]int
	array2 = &array1
	array2[0] =100
	fmt.Println("grogu")
	fmt.Println(array1)

}

func someFunction() (int, int) {
	return 1, 2
}