package webserver

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var emps = []Employee{
	{ID: 1, Name: "BAyu", Age: 29, Division: "IT"},
}

func GetEmployees() []Employee {
	return emps
}
