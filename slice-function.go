package main

import "fmt"

func main() {
	data := [...]string{
		"Senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu",
	}

	daySlice1 := data[:]
	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(data)

	daySlice2 := append(daySlice1, "libur baru")
	daySlice2[0] = "ups"
	fmt.Println(daySlice2)
	fmt.Println(data)
	fmt.Println(daySlice2)

}
