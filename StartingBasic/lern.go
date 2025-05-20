package main

import "fmt"

type person struct {
	name   string
	age    int
	salary float32
}

func (p person) display() string {
	return fmt.Sprintf("Name: %s, Age: %d, Salary: %.2f", p.name, p.age, p.salary)
}
// redeclaration

type display interface {
	display() string
}

func main() {
	fmt.Println(("hello world"))
	name := "mai kam chor hu"
	var age int = 20
	fmt.Println(name, age)

	// person1:=person{
	// 	name:   "shubham",
	// 	age:    20,
	// 	salary: 1000.50,
	// }
	var p person = person{
		name:   "shubham",
		age:    20,
		salary: 1000.50,
	}
	q := new(person)
	q.name = "shubh"
	q.age = 20
	q.salary = 1000.50

	fmt.Println(q.display())
	// p := person{"shubham", 20, 1000.50}

	fmt.Println(p.display())
	fmt.Println("printing  the display method")
	displayInfo()

	main2()

	
}
