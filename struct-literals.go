package main

import "fmt"

func main() {
	personOne := Person{
		Name:      "Fazrul Anugrah Sahi",
		Address:   "Jalan Manado",
		Age:       20,
		CanCoding: true,
	}
	personTwo := Person{"John", "USA", 20, false}
	fmt.Println("=======================================")
	personOne.makeApplication("Website")
	fmt.Println("=======================================")
	personTwo.makeApplication("Mobile")

}

type Person struct {
	Name, Address string
	Age           int
	CanCoding     bool
}

func (person Person) makeApplication(division string) {
	fmt.Println("Hello, my name is " + person.Name)
	if !person.CanCoding {
		fmt.Println("Sorry I don't Understand How to make Application")
		return
	}

	fmt.Println("Currently Im Make a " + division + " Application!")
}
