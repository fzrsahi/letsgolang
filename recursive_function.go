package main

import "fmt"

func main() {
	fmt.Println(recursiveLoop(5))
	fmt.Println(recursive(5))
}

func recursiveLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func recursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * recursive(value-1)
	}
}
