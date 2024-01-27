package main

import "fmt"

func loging() {
	fmt.Println("Logging Bosss.........")
	message := recover()
	fmt.Println("Error Panic:", message)
}

func executeApp(err bool) {
	defer loging()
	if err {
		panic("Error")
	}
}
func main() {
	executeApp(true)
	fmt.Println("Lanjut...")
}
