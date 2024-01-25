package main

import "fmt"

func main() {
	var laptopNames = [...]string{
		"asus", "lenovo", "huawei", "LG", "acer", "ROG", "macbook",
	}

	slice := laptopNames[2:5]
	slice2 := laptopNames[1:5]
	slice3 := laptopNames[0:3]
	slice4 := laptopNames[1:]

	var slice5 []string = laptopNames[:1]

	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)

	fmt.Println("Hello World")

}
