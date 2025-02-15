package gamelogic

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func StartGame() int {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("**************************************************")
		fmt.Println("Number guessing game")
		fmt.Println("**************************************************")
		fmt.Println("RULES:")
		fmt.Println("1. Choose a difficulty level: easy, medium, hard")
		fmt.Println("2. Guess the number between 1 and 100")
		fmt.Println("3. You have 10 attempts to guess the number")
		fmt.Println("4. After each guess, you will be told if you are too high or too low")
		fmt.Println("5. If you guess the number, you win!")
		fmt.Println("6. If you want to quit the game, type 'quit'")
		fmt.Println("**************************************************")
		fmt.Println("Enter your choice:")
		fmt.Println("**************************************************")

		scanner.Scan()
		guess := scanner.Text()

		if guess == "quit" {
			handleQuit()
		}

		if guess == "easy" || guess == "medium" || guess == "hard" {
			fmt.Println("You have chosen the difficulty level:", guess)
			if guess == "easy" {
				return 10
			} else if guess == "medium" {
				return 5
			} else if guess == "hard" {
				return 3
			}
			fmt.Println("**************************************************")
		} else {
			fmt.Println("Invalid input")
		}
	}
}

func PlayGame(numberOfAttempts int) {
	randomNumber := rand.Intn(100) + 1
	scanner := bufio.NewScanner(os.Stdin)
	attempts := 0
	startTime := time.Now()

	for {
		if attempts == numberOfAttempts {
			fmt.Println("You have no more attempts left.")
			fmt.Println("The number was:", randomNumber)
			break
		}

		fmt.Println("Enter your guess:")
		scanner.Scan()
		guess := scanner.Text()

		if guess == "quit" {
			handleQuit()
		}

		// * convert guess - string to int
		guessInt, err := strconv.Atoi(guess)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		attempts++

		if guessInt == randomNumber {
			endTime := time.Now()
			fmt.Println("Congratulations! You guessed the number!")
			fmt.Println("Number of attempts:", attempts)
			fmt.Println("Time taken:", endTime.Sub(startTime))
			break
		} else if guessInt < randomNumber {
			fmt.Println("Too low! Try again.")
		} else if guessInt > randomNumber {
			fmt.Println("Too high! Try again.")
		}

		fmt.Println("Number of attempts remaining:", numberOfAttempts-attempts)
		fmt.Println("**************************************************")
	}
}

func PlayAgain() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Do you want to play again? (y/n)")
	scanner.Scan()
	playAgain := scanner.Text()

	if playAgain == "n" {
		handleQuit()
	}
}

func handleQuit() {
	fmt.Println("Thank you for playing the Number Guessing Game!")
	os.Exit(0)
}
