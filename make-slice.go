package main

import "fmt"

func main() {
	slice := make([]string, 2, 5)
	slice[0] = "senin"
	slice[1] = "selasa"
	newSlice := append(slice, "rabu")

	newSlice[0] = "minggu"

	fmt.Println(slice)
	fmt.Println(newSlice)
}
