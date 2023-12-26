package main

import "time"

type Person struct {
	DNI  int
	Name string
	Age  int
}

type Employee struct {
	Id       int
	Position string
}

type FullTimeEmployee struct {
	Employee
	Person
}

var GetPersonById = func(dni int) (Person, error) {
	time.Sleep(5 * time.Second)
	// Filter or Query to BD
	return Person{}, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	// Filter or Query to BD
	return Employee{}, nil
}

func GetFullTimeEmployeeById(id int, dni int) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

	employee, err := GetEmployeeById(id)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Employee = employee

	person, err := GetPersonById(dni)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Person = person

	return ftEmployee, nil
}
