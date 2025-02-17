package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}
type Company struct {
	Employees []Employee
}
func (s *Company) TotalSalary() float64 {
	total := 0.0

	for _, employee := range s.Employees {
		total += employee.Salary
	}
	return total

}
func main() {
	company := Company{
		Employees: []Employee{
			{ID: 1, Name: "Sam", Salary: 100.5},
			{ID: 2, Name: "Cam", Salary: 200.4},
			{ID: 3, Name: "Dan", Salary: 400.65},
		},
	}

	total := company.TotalSalary()
	fmt.Println(total)

}
