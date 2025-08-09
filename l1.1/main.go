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
	human1 := Human{Name: "Vera", Age: 22}
	actionHuman1 := Action{Human: human1, Job: "Programmer", Hobby: "Play Guitar"}

	human1.PrintHumanName()
	human1.PrintHumanAge()

	actionHuman1.PrintActionHobby()
	actionHuman1.PrintActionJob()

	actionHuman1.PrintHumanName()
	actionHuman1.PrintHumanAge()
}
