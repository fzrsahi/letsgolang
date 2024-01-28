package main

import "fmt"

func main() {
	var laptopOne Laptop
	laptopOne.Name = "Lenovo Ideapad Gaming3"
	laptopOne.Price = 2000
	laptopOne.Brand = "Lenovo"
	fmt.Println(laptopOne)
}

type Laptop struct {
	Name, Brand string
	Price       int
}
