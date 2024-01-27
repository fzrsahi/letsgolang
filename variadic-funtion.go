package main

import "fmt"

func main() {
	days := []string{"senin", "selasa", "rabu"}
	getDays(7, "senin", "selasa", "rabu", "kami")
	getDays(7, days...)
}

func getDays(totalDays int, days ...string) {
	for _, day := range days {
		fmt.Println(day)
	}
}
