package advanced

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 5)

	consumer(ch)
	// time.Sleep(6 * time.Second)
	producer(ch)
}
func consumer( ch chan <- int){
	go func(ch chan <- int){
		for i:= range 5{
			time.Sleep(time.Second)
			fmt.Println("Value added: ", i)
			ch <- i
		}
		close(ch)
	}(ch)
}

func producer(ch chan int){
	for val:= range ch{
		fmt.Println("The consumed values are: ", val)
	}
}
