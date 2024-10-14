package main

import "fmt"

type Human struct {
	Name  string
	Age   int
	Hobby string
}

func (h *Human) Speak() {
	fmt.Printf("Hi, my name is %s, and I am %d years old.\n", h.Name, h.Age)
}
func (h *Human) MyHobby() {
	fmt.Printf("My hobby is  %s \n", h.Hobby)
}

type Action struct {
	Human
}

func main() {

	action := Action{
		Human: Human{
			Name:  "Andrey",
			Age:   25,
			Hobby: "motorcycles",
		},
	}

	action.Speak()
	action.MyHobby()
}
