package main

import (
	"fmt"
	"math/rand"
	"time"
)

func positiveDiffrence(guess, number int) int {
	if guess > number {
		return guess - number
	} else {
		return number - guess
	}
}

func number_guessing_game() {
	fmt.Scanln()
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(100)
	fmt.Println("Welcome to the number guessing game with go 🏁 !")
	fmt.Println("You have to Guess a number between 0 and 100")
	fmt.Println("You have 5 tries 🎏")
	for i := 0; i < 5; i++ {
		fmt.Printf("#%d Enter Your Guess : ", i+1)
		var guess int
		fmt.Scanf("%d \n", &guess)
		switch positiveDiffrence(guess, answer) {
		case 0:
			fmt.Println("You guessed it!")
			fmt.Println("You won! 🎊🎊🎊 🎉🎉🎉🎉🎉🎉")
			return
		case 1:
			fmt.Println("SO CLOSE!")
		case 2,3:
			fmt.Println("Almost there!")
		case 4,5:
			fmt.Println("Nearby!")
		default:
			if guess-answer > 5 {
				fmt.Println("Way too high!")
			} else {
				fmt.Println("Way too low!")
			}
		}

	}
	fmt.Println("You lost!")
	fmt.Printf("The answer was %d\n", answer)
}
