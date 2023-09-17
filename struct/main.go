package main

import (
	"fmt"

	"github.com/EvegeniyNekrasov/golang-study/struct/model"
)

func main() {
	instructor := model.Instructor{Id: 1, FirstName: "Jeka", LastName: "Nekrasov", Age: 20}
	kyle := model.NewInstructor("Paco", "Paquito")
	
	swiftWS := model.NewWorkshop("Swift", instructor)
	fmt.Println("%v", swiftWS)
	print(kyle.Print())
}