package model

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Age       int
}

// Factory, this is just a function
func NewInstructor(name string, lastName string) Instructor {
	return Instructor{FirstName: name, LastName: lastName}
}