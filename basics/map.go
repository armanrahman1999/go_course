package basics

import "fmt"

func main() {

	myMap := make(map[string]int)
	fmt.Println(myMap)
	myMap["Key1"] = 9
	myMap["code"] = 18
	// delete(myMap, "Key1")
	// clear(myMap)
	value, unknown := myMap["Key1"]
		fmt.Println(myMap, value, unknown)

	map2 := map[string]int{"a": 2, "b": 7}
	for _, b:= range map2{
		fmt.Println(b, "biatch")
	}
}