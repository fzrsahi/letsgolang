package main

import "fmt"

func main() {

	numbers := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	for i := 0; i < len(numbers); i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
