package basics

import (
	"fmt"
	"os"
)

func main() {

	// We do this when we run into a critical error
	defer fmt.Println("Rad deeee")
	fmt.Println("Start")

	os.Exit(1)

	fmt.Println("End")


}