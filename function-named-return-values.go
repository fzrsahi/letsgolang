package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println(getFullStats("fazrul", 2003))
	firstName, middleName, lastName := getMyName()
	fmt.Println(firstName, middleName, lastName)
	fmt.Println(getNameAndAge())

}

func getFullStats(name string, birth int) (string, string) {
	age := time.Now().Year() - birth
	return "hello nama saya :" + name + " ", "umur saya : " + strconv.Itoa(age)
}

func getMyName() (firstName, middleName, lastName string) {
	firstName = "fazrul"
	middleName = "anugrah"
	lastName = "sahi"

	return firstName, middleName, lastName
}

func getNameAndAge() (name string, age int) {
	name = "fazrul"
	age = 20

	return name, age
}
