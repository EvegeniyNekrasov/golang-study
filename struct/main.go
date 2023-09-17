package main

import (
	"fmt"

	"github.com/EvegeniyNekrasov/golang-study/struct/model"
)

func main() {
	instructor := model.Instructor{Id: 1, FirstName: "Jeka", LastName: "Nekrasov", Age: 20}
	kyle := model.NewInstructor("Paco", "Paquito")


	goCource := model.Course{Id: 2, Name: "Go fundamentals", Instructor: instructor}
	fmt.Println("%v", goCource)
	print(kyle.Print())
}