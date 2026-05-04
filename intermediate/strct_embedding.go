package intermediate

import "fmt"

type Person struct {
	name string
	age  int
}
type Employee struct {
	empDetails Person
	empId      string
}

func (p Person) details() {
    fmt.Printf("Hi my name is %s and my age is %d\n", p.name, p.age)
}
func (emp Employee) details() {
    fmt.Printf("Hi my name is %s and my age is %d, employee id: %s\n", emp.empDetails.name, emp.empDetails.age, emp.empId)
}
func main() {
	p := Person{name: "Raz", age: 20}
	emp := Employee{empDetails: Person{name: "Raz", age: 20}, empId: "1223345"}
	p.details()
	emp.details()
}