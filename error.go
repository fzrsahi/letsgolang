package main

import (
	"errors"
	"fmt"
)

func Distribution(value int, value2 int) (int, error) {
	if value2 == 0 {
		return 0, errors.New("Pembagian dengan nol")
	} else {
		return value / value2, nil
	}
}

func main() {

	result, err := Distribution(10, 0)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println("Error : ", err.Error())
	}
}
