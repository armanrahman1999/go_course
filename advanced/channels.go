package advanced

import "fmt"

func main() {

	//variable:= make(chan type)
	// we can not make a channel work inside any function 
	//It needs goroutine
	greeting := make(chan string)
	greetString := "Hello"
	go func(){
		greeting <- greetString

	}()
	reciever := <-greeting
	fmt.Println(reciever)
}