package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
	Person
}

type FullTimeEmployee struct {
	Employee
	//Person
	active bool
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "John"
	ftEmployee.age = 30
	ftEmployee.id = 1
	ftEmployee.active = true
	fmt.Printf("%v\n", ftEmployee)
}
