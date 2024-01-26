package main

import "fmt"

func main() {
	number := 1

	for number <= 10 {
		fmt.Println("Perulangan ke", number)
		number++
	}

	for i := 0; i < 10; i++ {
		fmt.Println("Counter Ke", i)
	}

}
