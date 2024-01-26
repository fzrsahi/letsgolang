package main

import "fmt"

func main() {
	name := "fzrsahi"
	if name != "fazrul" {
		fmt.Println("Hello Not fazrul")
	} else if name == "fazrul" {
		fmt.Println("Hello fazrul")
	} else {
		fmt.Println("Hello World")
	}

}
