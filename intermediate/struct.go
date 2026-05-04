package intermediate

import "fmt"

	type Person struct {
		firstName string 
		lastName string 
		age  int
	}
func main() {



	p := Person{
		firstName: "Raz",
		lastName: "Dee",
		age:  20,
	}
	p2 := Person{
		firstName: "Dree",
		lastName: "Dree",
		age:  30,
	}
	fmt.Println(p.firstName, p.lastName, p.age)
	fmt.Println(p.fullName())
	p2.incrementAge()
	fmt.Println(p2.age)
}

func (p Person) fullName() string{
	return p.firstName + " " + p.lastName
}

func (a *Person) incrementAge() {
	a.age ++
}