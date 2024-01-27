package main

import "fmt"

func main() {
	sayHello()
	name := "fazrul"
	fmt.Println(sayHelloTo(name))
	fmt.Println(getFullName())
}

func sayHello() {
	fmt.Println("Hello, World!")
}

func sayHelloTo(name string) string {
	return "hello " + name
}

func getFullName() (string, string, string) {
	return "fazrul", "anugrah", "sahi"
}
