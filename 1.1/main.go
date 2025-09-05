package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) GetName() string {
	return h.Name
}

func (h *Human) GetAge() int {
	return h.Age
}

type Action struct {
	Human
}

func main() {
	human := &Human{
		Name: "Ivan",
		Age:  25,
	}
	fmt.Println(human.GetName(), human.Age)

	action := &Action{
		Human: Human{
			Name: "Ivan",
			Age:  25,
		},
	}
	fmt.Println(action.GetName(), action.GetAge())
}
