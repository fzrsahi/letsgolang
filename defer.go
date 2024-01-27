package main

import "fmt"

func logging() {
	fmt.Println("Ada Logging")
}

func execute() {
	defer logging()
	fmt.Println("Execute1")
	fmt.Println("Execute2")
	fmt.Println("Execute3")
	fmt.Println("Execute4")
	fmt.Println("Execute5")
	fmt.Println("Execute6")
}

func main() {
	execute()
}
