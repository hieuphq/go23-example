package guestnumber

import (
	"fmt"
	"math/rand"
	"time"
)

// Constants for guess comparison
const (
	CorrectGuess = iota
	HighGuess
	LowGuess
)

// GenerateRandomNumber generates a random number within the specified range
func GenerateRandomNumber(low, high int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(high-low+1) + low
}

// GetUserGuess takes and validates user input for guessing the number
func GetUserGuess() int {
	var guess int
	fmt.Print("Enter your guess: ")
	_, err := fmt.Scanln(&guess)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return GetUserGuess()
	}
	return guess
}

// CheckGuess compares the user's guess with the secret number and returns the result
func CheckGuess(guess, secretNumber int) int {
	if guess == secretNumber {
		return CorrectGuess
	} else if guess > secretNumber {
		return HighGuess
	} else {
		return LowGuess
	}
}

func RandomStrategy() func(int, int) int {
	return func(i1, i2 int) int {
		return i1 + i2
	}
}

func ChangeRule(secretNumber int, low, high int, randomFn func(int, int) int) int {
	newSecret := secretNumber
	for {
		newSecret = randomFn(low, high)
		if newSecret != secretNumber {
			return newSecret
		}

	}
}

func Var(low, high int, randomStrategies ...func(int, int) int) int {
	rs := 0
	for i := 0; i < len(randomStrategies); i++ {
		s := randomStrategies[i]
		rs = rs + s(low, high)
	}

	return rs
}
