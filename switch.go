package main

import "fmt"

func main() {
	name := "fazruls"

	switch name {
	case "fazrul":
		fmt.Println("hello", name)
	default:
		fmt.Println("Hello Not Fazrul")
	}

	switch fixName := name; fixName == "fazrul" {
	case true:
		fmt.Println("hello", name)
	default:
		fmt.Println("Hello Not Fazrul")
	}

}
