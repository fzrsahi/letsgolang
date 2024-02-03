package main

import "fmt"

type Person struct {
	Name, Address, Age string
}

func main() {
	person1 := &Person{"fazrul", "Jl.Manado", "20"}
	changeValue(person1)
	fmt.Println(person1)
}

func changeValue(person *Person) {
	person.Name = "John"
}
