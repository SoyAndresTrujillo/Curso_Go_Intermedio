package main

import "fmt"

type Employee struct {
	name   string
	id     int
	active bool
}

func newEmployee(name string, id int, active bool) *Employee {
	return &Employee{
		name:   name,
		id:     id,
		active: active,
	}
}

func main() {
	// Create a new employee
	//emp := new(Employee)
	//emp := Employee{}
	//emp := Employee{
	//	name:   "John",
	//	id:     1,
	//	active: true,
	//}
	emp := newEmployee("John", 1, true)
	fmt.Println(emp)
	// Create a new employee with the builder
}
