package main

import "fmt"

func main() {
	array := [...]string{
		"senin", "selasa", "rabu",
	}
	slice := []string{
		"senin", "selasa", "rabu",
	}

	fmt.Println(array)
	fmt.Println(slice)
}
