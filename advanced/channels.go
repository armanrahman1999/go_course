package advanced

import (
	"fmt"
	"time"
)

func main() {

	// variable := make( chan string)
	greetings := make(chan string)
	greetHello := "Hello"

	// to make it receive value twice we have to do this same thing again
	go func() {
		time.Sleep(2 * time.Second)
		greetings <- greetHello
		greetings <- "World"

		for _, e := range "abcde"{
			greetings <- "Alphabet:" + string(e) 
		}
	}()
	
	go func(){
	reciever := <-greetings
	fmt.Println(" This is the channel reciever",reciever )
	reciever = <- greetings

	fmt.Println(" This is the channel reciever",reciever )
	for range 5{
		rcvr := <- greetings
		fmt.Println(rcvr)
	}
	}()
	
	time.Sleep(3 * time.Second)
	fmt.Println("End of program")

}