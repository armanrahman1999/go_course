package basics

import "fmt"

func main() {

	something, result := sum("rad dee",1, 2, 3)
	// kaiju := sum("dree dree")
	fmt.Println(something, result)
	// fmt.Println(kaiju)
	//Variadic params needs to be used in the end

}

func sum(name string, nums ...int) (string, int) {
	total := 0
	for _, v := range nums {
		total += v

	}
	return name, total
}