package main

import "fmt"

func main() {
	names := []string{
		"fazrul", "anugrah", "sahi",
	}

	for i := 0; i < len(names); i++ {
		//fmt.Println(names[i])
	}

	for i, name := range names {
		fmt.Println("index:", i, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
