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
	fmt.Println("Tidak Dijalankan!")
}
func main() {
	runApp(true)
}
