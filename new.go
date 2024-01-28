package main

import "fmt"

type Person struct {
	Name, Address string
	Age           int
}

func main() {
	person1 := new(Person)
	person2 := person1
	person3 := person1

	person2.Name = "fazrul"

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)
}
