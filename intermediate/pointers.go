package intermediate

import "fmt"

func main() {

	var ptr *int 
	var a int = 10
	ptr = &a // referencing

	fmt.Println(a)
	fmt.Println(ptr)
	fmt.Println(*ptr) //dereferencing a pont
	// 0 values of pointer is nil
	modValue(ptr)
	fmt.Println("func chnages", a)


}

func modValue( ptr *int) {
 *ptr++
}