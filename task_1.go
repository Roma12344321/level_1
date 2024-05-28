package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) SayHello() {
	fmt.Printf("Name is %s and %d age", h.Name, h.Age)
}

type Action struct {
	Human
	Activity string
}

func (a *Action) PerformAction() {
	fmt.Printf("%s is performing: %s", a.Human.Name, a.Activity)
}

func task_1() {
	action := Action{
		Human: Human{
			Name: "John",
			Age:  30,
		},
		Activity: "Running",
	}
	action.SayHello()
	fmt.Println()
	action.PerformAction()
}
