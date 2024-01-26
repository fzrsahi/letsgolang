package main

import "fmt"

func main() {
	realName := "fazrul"
	if name := realName; name == "fazrul" {
		fmt.Println("Hello Fazrul")
	} else {
		fmt.Println("Hello Not Fazrul")

	}
}
