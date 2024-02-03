package main

import "fmt"

type PersonName struct {
	FirstName, LastName string
}

func main() {
	name := new(PersonName)
	name2 := name
	name2.FirstName = "John"
	name3 := name2

	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name3)
}
