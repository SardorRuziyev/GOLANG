package main

import "fmt"

type Employee struct {
	ID           int
	Name         string
	DepartmentID int
	Projects     []string
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

func (e *Employee) AssignToProject(projectName Project) {
	e.Projects = append(e.Projects, projectName)
}
func (d *Department) AddEmployee(employee Employee) {
	d.Employees = append(d.Employees, employee)

}
func (d *Department) GetEmployees() []Employee {
	return d.Employees

}

func main() {
	employee1 := Employee{
		ID:           1,
		Name:         "Eshmirza",
		DepartmentID: 101,
		Projects:     []string{},
	}
	employee2 := Employee{
		ID:           2,
		Name:         "Toshmirza",
		DepartmentID: 102,
		Projects:     []string{},
	}
	// employee3 := Employee{
	// 	ID:           3,
	// 	Name:         "Toshpolat",
	// 	DepartmentID: 101,
	// 	Projects:     []string{},
	// }
	// employee4 := Employee{
	// 	ID:           4,
	// 	Name:         "EshPolat",
	// 	DepartmentID: 102,
	// 	Projects:     []string{},
	// }

	department1 := Department{ID: 101, Name: "Human Resource", Employees: []Employee{}}
	//department2 := Department{ID: 102, Name: "ENginering", Employees: []Employee{}}

	project1 := Project{ID: 201, Name: "AI Robotics", DepartmentID: 101, Employees: []Employee{}}
	project2 := Project{ID: 202, Name: "Research on Robots", DepartmentID: 102, Employees: []Employee{}}

	employee1.AssignToProject(project1)
	employee2.AssignToProject(project2)

	project1.AddEmployee(employee1)
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
