package main

import (
	"fmt"

	"github.com/EvegeniyNekrasov/golang-study/struct/model"
)

func main() {
	instructor := model.Instructor{Id: 1, FirstName: "Jeka", LastName: "Nekrasov", Age: 20}
	kyle := model.NewInstructor("Paco", "Paquito")
	
	goCource := model.Course{Id: 2, Name: "Go fundamentals", Instructor:  instructor}

	swiftWS := model.NewWorkshop("Swift", instructor)
	fmt.Println("%v", swiftWS)
	print(kyle.Print())

	var courses [2]model.Signable
	courses[0] = goCource
	courses[1] = swiftWS
}