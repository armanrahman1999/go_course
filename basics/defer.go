package basics

import "fmt"

func main() {
	process()
}

func process() {
	// Thiis works in lifo
	defer fmt.Println("This is dfered")
	defer fmt.Println("dree dree")
	fmt.Println(" This is normal execution")
}