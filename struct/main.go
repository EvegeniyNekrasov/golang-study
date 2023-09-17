package main

import (
	"fmt"
	"time"

	"github.com/EvegeniyNekrasov/golang-study/struct/model"
)

func printMessage(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(800 * time.Millisecond)
	}
}

func main() {
	instructor := model.Instructor{Id: 1, FirstName: "Jeka", LastName: "Nekrasov", Age: 20}
	// kyle := model.NewInstructor("Paco", "Paquito")
	
	goCource := model.Course{Id: 2, Name: "Go fundamentals", Instructor:  instructor}
	swiftWS := model.NewWorkshop("Swift", instructor)

	var courses [2]model.Signable
	courses[0] = goCource
	courses[1] = swiftWS

	go printMessage("******** hola...")
	go printMessage("--------- hola...")
	printMessage("####### hola2...")

}