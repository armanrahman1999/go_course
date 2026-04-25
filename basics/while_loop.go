package basics

import "fmt"
func main() {

	

	sum := 0

	for {
		fmt.Println("Sum :", sum)
		sum += 10
		if sum >= 50 {
			break
		}
	}
}