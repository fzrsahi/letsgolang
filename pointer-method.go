package main

import "fmt"

type Person struct {
	Name, Age string
}

func main() {
	person1 := Person{"fazrul", "20"}
	person1.ChangeName()
	fmt.Println(person1)
}

func (person *Person) ChangeName() {
	person.Name = "Mr " + person.Name
}
