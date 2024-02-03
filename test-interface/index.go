package test_interface

import "fmt"

func Gas() {
	fmt.Println(Diet(&Person{"fazrul", "20", 130}))
	orangPertama := &Person{
		Name:   "John",
		Age:    "21",
		Calory: 200,
	}

	fmt.Println(orangPertama.Run())
}
