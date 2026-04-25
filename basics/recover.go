package basics

import "fmt"
func main() {

	process()
	fmt.Print("After panic")

}

func process(){
	defer func(){
		if r:= recover(); r!= nil{
			fmt.Println("Recovered", r)
		}
	}()

	fmt.Println("before Panic")
	panic("lets panic")
}