package main

import (
	"fmt"
	"math/rand"
)

func main() {
	guessNumber(1, func(guess int) bool {
		if randomNumber := rand.Intn(5); guess != randomNumber {
			fmt.Println("Yang Benar Adalah :", randomNumber)
			return false
		}
		return true
	})
}

type GuessType func(int) bool

func guessNumber(number int, guess GuessType) {
	fmt.Println("Jawaban Anda Adalah : ", number)
	if guess(number) {
		fmt.Println("Anda Benar")
	} else {
		fmt.Println("Anda Salah")
	}
}
