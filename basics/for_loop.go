package basics

import "fmt"
func main() {
 
	for i:= 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}
	numbers := []int{1, 2, 3, 4, 5}

	for index, value:= range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	for i:= 0; i<=10; i++{
		if i%2 == 0 {
			continue
		}else if i == 9{
			break
		}else{
			fmt.Println("Odd numbers are : ", i )
		}
	}
	fmt.Println("-------------seperation-------------")
	for i:= range 10{
		fmt.Println(i)
	}
}