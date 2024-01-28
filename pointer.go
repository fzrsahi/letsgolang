package main

import "fmt"

func main() {
	days := [...]string{
		"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu",
	}

	newDays := &days //pointer
	newDays[0] = "senin baru"

	fmt.Println(days) // Senin tidak berubah menjadi senin baru
	fmt.Println(newDays)
}
