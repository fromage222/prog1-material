package main

import (
	"fmt"
	"math/rand"
)

func ReadNumber() int {
	fmt.Print("Rate eine Zahl: ")
	var n int
	fmt.Scan(&n)
	return n
}

func GuessingGame() {
	my_number := rand.Intn(101) - 50
	for n := 0; n < 3; n++ {
		guess := ReadNumber()
		if NumberGood(guess, my_number) {
			fmt.Println("Richtig geraten :-)")
			return
		} else {
			if guess < my_number {
				fmt.Println("Die gesuchte Zahl ist größer")
			} else {
				fmt.Println("Die gesuchte Zahl ist kleiner")
			}
			fmt.Println("Falsch geraten (Muhaha)")
		}
	}
	fmt.Println("Zu viele falsche Zahlen :-(")
	fmt.Println("Die richtige Zahl war: ", my_number)
}

func NumberGood(g, e int) bool {
	return g == e
}

func main() {
	GuessingGame()
}
