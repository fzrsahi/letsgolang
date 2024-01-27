package main

import "fmt"

func log() {
	fmt.Println("Logging Bosss.........")
}

func runApp(err bool) {
	defer log()
	if err {
		panic("Error")
	}
}
func main() {
	runApp(true)
}
