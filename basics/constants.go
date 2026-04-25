package basics

import "fmt"

const PI = 3.14
const GRAVITY = 9.81


func main(){

	const (	sunday = 1
		monday = 2
		tuesday int  = 3)
	
	var smallFloat float64 = 1.0e-323
	
	fmt.Println("PI:", PI, sunday)
	fmt.Println("GRAVITY:", GRAVITY, smallFloat)	
}