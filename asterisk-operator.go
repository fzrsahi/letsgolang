package main

import "fmt"

func main() {
	days := [...]string{
		"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu",
	}

	newDays := &days //pointer
	newDays[0] = "senin baru"
	fmt.Println(newDays, "newdays1")

	*newDays = [...]string{
		"senin baru", "selasa baru", "rabu baru", "kamis baru", "jumat baru", "sabtu baru", "minggu baru",
	}

	fmt.Println(days, "days============") // Senin tidak berubah menjadi senin baru
	fmt.Println(newDays, "newdays1")
}
