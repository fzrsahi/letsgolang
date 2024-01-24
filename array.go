package main

import "fmt"

func main() {
	var array = [10]int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	array2 := [...]int{
		0, 1, 2, 3, 4,
	}

	fmt.Println(array)
	fmt.Println(len(array2))
}
