package main

import "fmt"

func main() {
	var personMap map[string]string = map[string]string{}
	personMap["name"] = "fazrul"
	personMap["age"] = "20"
	fmt.Println(personMap)

	personMap2 := map[string]string{}
	personMap2["name"] = "fazrul"
	personMap2["age"] = "20"
	fmt.Println(personMap2)

	fmt.Println("==========================")
	personMap3 := make(map[string]string)
	personMap3["name"] = "fazrul"
	personMap3["age"] = "20"
	personMap3["address"] = "Jl. Manado"
	fmt.Println(personMap3)
	delete(personMap3, "address")
	fmt.Println(personMap3)
	fmt.Println("==========================")

	lenMap := len(personMap)
	keyMap := personMap["name"]

	fmt.Println(lenMap)
	fmt.Println(keyMap)

}
