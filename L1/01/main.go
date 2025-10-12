package main

import "fmt"

type Human struct {
	person   string
	bodyPart string
}

type Action struct {
	*Human
}

func (h Human) getPerson() string {
	return h.person
}

func (h Human) getBodyPart() string {
	return h.bodyPart
}

func (a Action) askToRaiseBodyPart() string {
	person := a.getPerson()
	bodyPart := a.getBodyPart()
	return person + ", please, raise your " + bodyPart
}

func (a Action) askToSacrifice() string {
	return " to sacrifice yourself for the sake of the Milky Way Galaxy !"
}

func main() {
	a := Action{
		Human: &Human{
			person:   "Human being Bob",
			bodyPart: "left hand",
		},
	}
	fmt.Println(a.askToRaiseBodyPart() + a.askToSacrifice())
}
