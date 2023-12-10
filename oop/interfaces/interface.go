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

type TemporaryEmployee struct {
	Employee
	//Person
	vacation bool
}

type PrintInfo interface {
	getMessage() string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full time employee"
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary employee"
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "John"
	ftEmployee.age = 30
	ftEmployee.id = 1
	ftEmployee.active = true
	fmt.Printf("%v\n", ftEmployee)

	tEmployee := TemporaryEmployee{}
	tEmployee.name = "Mike"
	tEmployee.age = 25
	tEmployee.id = 2
	tEmployee.vacation = false
	fmt.Printf("%v\n", tEmployee)

	getMessage(ftEmployee)
	getMessage(tEmployee)
}
