package main

import "fmt"

type Employee struct {
	ID           int
	Name         string
	DepartmentID int
	//Projects     []Project
}
type Department struct {
	ID        int
	Name      string
	Employees []Employee
}

type Project struct {
	ID           int
	Name         string
	DepartmentID int
	Employees    []Employee
}

func (e *Employee) AssignToProject(projectName *Project) {
	projectName.Employees = append(projectName.Employees, *e)
}

func (d *Department) AddEmployee(employee Employee) {
	d.Employees = append(d.Employees, employee)
}

func (p *Project) AddEmployee(employee Employee) {
	p.Employees = append(p.Employees, employee)
}

func (d *Department) GetEmployees() []Employee {
	return d.Employees

}

func main() {
	employee1 := Employee{
		ID:           1,
		Name:         "Eshmirza",
		DepartmentID: 101,
	}
	employee2 := Employee{
		ID:           2,
		Name:         "Toshmirza",
		DepartmentID: 102,
	}
	employee3 := Employee{
		ID:           3,
		Name:         "Toshpolat",
		DepartmentID: 101,
	}
	employee4 := Employee{
		ID:           4,
		Name:         "EshPolat",
		DepartmentID: 102,
	}

	department1 := Department{ID: 101, Name: "Human Resource", Employees: []Employee{employee1, employee2, employee3, employee4}}
	//department2 := Department{ID: 102, Name: "ENginering", Employees: []Employee{}}

	project1 := Project{ID: 201, Name: "AI Robotics", DepartmentID: 101, Employees: []Employee{employee1, employee2}}
	project2 := Project{ID: 202, Name: "Research on Robots", DepartmentID: 102, Employees: []Employee{employee2, employee3, employee4}}

	employee1.AssignToProject(&project1)
	employee2.AssignToProject(&project2)

	department1.AddEmployee(employee1)
	project2.AddEmployee(employee2)

	fmt.Println("Employees in Engineering Department:")
	for _, emp := range department1.GetEmployees() {
		fmt.Printf("ID: %d, Name: %s, Projects: %v\n", emp.ID, emp.Name)
	}

	fmt.Println("\nProject Employees:")
	for _, emp := range project1.Employees {
		fmt.Printf("ID: %d, Name: %s\n", emp.ID, emp.Name)
	}

}
