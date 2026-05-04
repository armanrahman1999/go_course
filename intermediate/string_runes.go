package intermdeiate

import "fmt"

func main() {

	message := "hello there \nGo!"
	message1 := "hello there \tGo!"
	message2 := "hello there \rGo!"
	rawMessage := `hello there \nboy`
	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	for i, char:= range message{
		fmt.Printf("This is the index %d and th char is %c \n", i, char)
	} 
}