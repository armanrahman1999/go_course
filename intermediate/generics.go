package intermediate

import "fmt"

func swap[T any](a, b T) (T, T) {
	return b, a
}

type Stack[T any] struct{
	elements []T
}
func (s *Stack[T]) Push(element T){
	s.elements = append(s.elements, element)
}
func (s *Stack[T]) PrintStack(){
	for _, val := range s.elements{
		fmt.Println(val)
	} 
}
func main() {
	x, y := 1, 2
	x, y = swap(x, y)
	fmt.Println(x, y)
	// Since this is any anything will work

	x1, y1 := "Raz", "Dee"
	x1, y1 = swap(x1, y1)
	fmt.Println(x1, y1)

	// s := &Stack[interface{}]{}
	// s.Push(5)
	// s.Push("Hello")
	// s.Push('c')
	// s.pop()
	// s.pop()
	// s.PrintStack()

	intStack := Stack[int]{}
	intStack.Push(5)
	intStack.Push(51)
	intStack.Push(4)
	intStack.Push(4)
	intStack.PrintStack()

}

func (s *Stack[T]) pop() (element T, flag bool){
	length := len(s.elements) - 1
	element = s.elements[length]
	s.elements = s.elements[: length]
	return element, true
}