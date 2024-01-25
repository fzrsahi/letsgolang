package main

import "fmt"

func main() {
	days := [...]string{
		"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu", "hari baru",
	}

	fromSlice := days[:]
	toSlice := make([]string, len(days), cap(days))
	copy(toSlice, fromSlice)
	fmt.Println(toSlice)
	fmt.Println([8]string(fromSlice))

}
