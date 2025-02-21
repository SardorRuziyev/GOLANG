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
	Name      string // Changed from int to string
	Employees []Employee
}

type Project struct {
	ID           int
	Name         string
	DepartmentID int
	Employees    []Employee
}

// AssignToProject assigns an employee to a project by adding the project name to their Projects list
func (e *Employee) AssignToProject(projectName string) { // Changed parameter to projectName
	e.Projects = append(e.Projects, projectName)
}

// AddEmployee adds an employee to the department's Employees list
func (d *Department) AddEmployee(employee Employee) {
	d.Employees = append(d.Employees, employee)
}

// GetEmployees returns the list of employees in the department
func (d *Department) GetEmployees() []Employee {
	return d.Employees
}

// AddEmployee adds an employee to the project's Employees list
func (p *Project) AddEmployee(employee Employee) {
	p.Employees = append(p.Employees, employee)
}

func main() {
	// Create employees
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
	employee3 := Employee{
		ID:           3,
		Name:         "Toshpolat",
		DepartmentID: 101,
		Projects:     []string{},
	}
	employee4 := Employee{
		ID:           4,
		Name:         "EshPolat",
		DepartmentID: 102,
		Projects:     []string{},
	}

	// Create departments
	department1 := Department{ID: 101, Name: "Human Resource", Employees: []Employee{}}
	department2 := Department{ID: 102, Name: "Engineering", Employees: []Employee{}}

	// Create projects
	project1 := Project{ID: 201, Name: "AI Robotics", DepartmentID: 101, Employees: []Employee{}}
	project2 := Project{ID: 202, Name: "Research on Robots", DepartmentID: 102, Employees: []Employee{}}

	// Add employees to departments
	department1.AddEmployee(employee1)
	department1.AddEmployee(employee3)
	department2.AddEmployee(employee2)
	department2.AddEmployee(employee4)

	// Assign employees to projects
	employee1.AssignToProject(project1.Name)
	employee2.AssignToProject(project2.Name)

	// Add employees to projects
	project1.AddEmployee(employee1)
	project2.AddEmployee(employee2)

	// Display employees in departments
	fmt.Println("Employees in Human Resource Department:")
	for _, emp := range department1.GetEmployees() {
		fmt.Printf("ID: %d, Name: %s, Projects: %v\n", emp.ID, emp.Name, emp.Projects)
	}

	fmt.Println("\nEmployees in Engineering Department:")
	for _, emp := range department2.GetEmployees() {
		fmt.Printf("ID: %d, Name: %s, Projects: %v\n", emp.ID, emp.Name, emp.Projects)
	}

	// Display employees working on projects
	fmt.Println("\nEmployees working on AI Robotics:")
	for _, emp := range project1.Employees {
		fmt.Printf("ID: %d, Name: %s\n", emp.ID, emp.Name)
	}

	fmt.Println("\nEmployees working on Research on Robots:")
	for _, emp := range project2.Employees {
		fmt.Printf("ID: %d, Name: %s\n", emp.ID, emp.Name)
	}
}
