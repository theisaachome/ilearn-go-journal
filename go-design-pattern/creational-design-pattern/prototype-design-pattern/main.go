package main

import "fmt"

// Step-1 Define the structure of the object
type Employee struct {
	Name        string
	EmployeeID  string
	Designation string
	Department  string
	Skills      []string
}

// Step-2 Create a function to clone the object

func (e *Employee) Clone() *Employee {
	// copy the original struct into a new struct
	cloned := *e
	// deep copy the slice (skills)
	// make a new slice of the same length
	// the original slice is not modified
	// the original and cloned employee have independent slice of skills
	cloned.Skills = make([]string, len(e.Skills))
	copy(cloned.Skills, e.Skills)
	return &cloned
}

type EmployeeRepository struct {
	// templates is a map of Employee objects
	// example: templates["HR"] = &Employee{...}
	templates map[string]*Employee
}

func NewEmployeeRepository() *EmployeeRepository {
	return &EmployeeRepository{
		templates: make(map[string]*Employee),
	}
}
func (r *EmployeeRepository) Register(key string, employee *Employee) {
	r.templates[key] = employee
}

func (r *EmployeeRepository) Create(key string) *Employee {
	// get the template from the map
	template, exists := r.templates[key]
	if !exists {
		panic("employee template not found" + key)
	}
	return template.Clone()
}

func main() {
	registry := NewEmployeeRepository()

	// create a new Employee struct
	javaDeveloperTemplate := &Employee{Name: "Template", EmployeeID: "Temp0123", Designation: "Java Developer", Department: "Engineering", Skills: []string{"Java", "Spring Boot"}}
	// now that we have the employee template, we can register it with the registry
	registry.Register("JAVA_DEVELOPER", javaDeveloperTemplate)

	// hire a new employee and new to register
	// therefore we clone the template

	employee001 := registry.Create("JAVA_DEVELOPER")
	employee001.Name = "John"
	employee001.EmployeeID = "EMP001"
	employee001.Skills = append(employee001.Skills, "Kotlin", "Container", "Docker")

	// second Employee 2
	employee002 := registry.Create("JAVA_DEVELOPER")
	employee002.Name = "Jane"
	employee002.EmployeeID = "EMP002"

	fmt.Printf("%+v\n", employee001)
	fmt.Printf("%+v\n", employee002)
}

/*
Understand the Go cloning mechanism
In Go, struct assignment copies the fields by value, but reference-like fields need special attention
If you modify an existing element through one slice, the change can be visible through the other.

original := Employee{
    Name:   "Isaac",
    Skills: []string{"Java", "Spring Boot"},
}

cloned := original

Output//
*/
