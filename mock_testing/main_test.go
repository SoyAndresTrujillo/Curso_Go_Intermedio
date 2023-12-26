package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              int
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: 12345678,
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "CEO",
					}, nil
				}
				GetPersonById = func(dni int) (Person, error) {
					return Person{
						DNI:  12345678,
						Name: "John Doe",
						Age:  30,
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Age:  30,
					DNI:  12345678,
					Name: "John Doe",
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
		},
	}

	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonById

	for _, test := range table {
		test.mockFunc()
		result, _ := GetFullTimeEmployeeById(test.id, test.dni)
		if result != test.expectedEmployee {
			t.Errorf("Expected employee: %v, but got: %v", test.expectedEmployee, result)
		}
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonById = originalGetPersonByDNI
}
