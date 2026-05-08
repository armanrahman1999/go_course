package main

import (
	"fmt"
	"time"
)

// goroutine is a lightweight thread of execution managed by the Go runtime. It allows you to run multiple					 functions concurrently without blocking the main thread. When you call a function with the "go" keyword, it runs in a separate goroutine, allowing the main function to continue executing without waiting for the goroutine to finish.
func main() {
	var err error
	fmt.Println("Starting main function")
	go sayHello()
	fmt.Println("This is main function")
	time.Sleep(2 * time.Second)
	go printNumbers()
	go printLetters()
	go func(){
		err = doWork()
			if err != nil{
		fmt.Println("Error:", err)
	}else{
		fmt.Println("The bitch is back")
	}
	}()


	time.Sleep(10 * time.Second)

}
func sayHello() {
	time.Sleep(4 * time.Second)
	fmt.Println("Hello")
}

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for ch := 'A'; ch <= 'E'; ch++ {
		fmt.Printf("%c\n", ch)
	}
}

func doWork() error{
	time.Sleep(1 * time.Second)
	return fmt.Errorf("Shit went siedways")
}