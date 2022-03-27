package main

import "fmt"

type Human struct {
	Age int
}

func (h *Human) HumanFunc() {
	fmt.Println("I'm human")
}

type Action struct {
	Human
}

func main() {
	var action Action

	action.HumanFunc()
}
