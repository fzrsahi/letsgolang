package main

import "fmt"

func main() {
	hello := sayHelo
	fmt.Println(hello("fazrul"))
}

func sayHelo(name string) string {
	return "hello " + name
}
