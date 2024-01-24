package main

import "fmt"

func main() {
	type NoKTP string
	randomNik := "123213"
	const nik NoKTP = "7571"
	contohUseWithNoKtp := NoKTP(randomNik)

	fmt.Println(randomNik)
	fmt.Println(nik)
	fmt.Println(contohUseWithNoKtp)
}
