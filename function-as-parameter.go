package main

import (
	"fmt"
	"math/rand"
)

func main() {
	guestNumber(1, guestNumberAlgorithm)
	fmt.Println("===============================")
	guestNumber(3, guestNumberAlgorithm)
	fmt.Println("===============================")
	guestNumber(5, guestNumberAlgorithm)
	fmt.Println("===============================")
}

type Guess func(int) bool

func guestNumber(number int, guess Guess) {
	fmt.Println("Jawaban Anda Adalah : ", number)
	if guess(number) {
		fmt.Println("Anda Benar")
	} else {
		fmt.Println("Anda Salah")
	}
}

func guestNumberAlgorithm(guess int) bool {
	if randomNumber := rand.Intn(5); guess != randomNumber {
		fmt.Println("Yang Benar Adalah :", randomNumber)
		return false
	}
	return true
}
