package main

import (
	"fmt"
)

type Human struct {
	Name string
	Age  uint8
}

type Action struct {
	Human
	Job   string
	Hobby string
}

func (h Human) PrintHumanName() {
	fmt.Printf("Name: %v\n", h.Name)
}

func (h Human) PrintHumanAge() {
	fmt.Printf("Age: %v\n", h.Age)
}

func (a Action) PrintActionJob() {
	fmt.Printf("Job: %v\n", a.Job)
}

func (a Action) PrintActionHobby() {
	fmt.Printf("Hobby: %v\n", a.Hobby)
}

func main() {
	// Create an instance of Human
	human1 := Human{Name: "Vera", Age: 22}

	// Create an instance of Action that embeds human1
	actionHuman1 := Action{Human: human1, Job: "Programmer", Hobby: "Play Guitar"}

	// Call methods of Human directly
	human1.PrintHumanName()
	human1.PrintHumanAge()

	// Call methods of Action
	actionHuman1.PrintActionHobby()
	actionHuman1.PrintActionJob()

	// Call inherited methods of Human through Action
	actionHuman1.PrintHumanName()
	actionHuman1.PrintHumanAge()
}
