package main

import (
	"fmt"

	guestnumber "github.com/dwarvesf/go23/ex2/game"
)

func main() {
	low := 1
	high := 100

	secretNumber := guestnumber.GenerateRandomNumber(low, high)

	fmt.Println("Welcome to Guess the Number!")
	fmt.Printf("Guess a number between %d and %d:\n", low, high)

	for {
		guess := guestnumber.GetUserGuess()

		result := guestnumber.CheckGuess(guess, secretNumber)

		if result == guestnumber.CorrectGuess {
			fmt.Println("Congratulations! You guessed the correct number!")
			break
		} else if result == guestnumber.HighGuess {
			guestnumber.Var(low, high, guestnumber.RandomStrategy(), random)
			fmt.Println("Too high! Try again.")
		} else if result == guestnumber.LowGuess {
			fmt.Println("Too low! Try again.")
		}
	}
}

func random(low, high int) int {
	return 1
}
