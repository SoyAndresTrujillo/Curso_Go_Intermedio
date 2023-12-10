package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) setID(id int) {
	e.id = id
}

func (e *Employee) setName(name string) {
	e.name = name
}

func (e *Employee) getID() int {
	return e.id
}

func (e *Employee) getName() string {
	return e.name
}

func main() {
	e := Employee{1, "John"}

	e.setID(2)
	e.setName("Jane")
	fmt.Println("ID Employee:", e.getID())
	fmt.Println("Name Employee:", e.getName())
}
