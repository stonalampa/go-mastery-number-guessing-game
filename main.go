package main

import (
	"number-guessing-game/gamelogic"
)

func main() {
	for {
		// * Start the game
		numberOfAttempts := gamelogic.StartGame()

		// * Play the game
		gamelogic.PlayGame(numberOfAttempts)

		// * Play again or quit
		gamelogic.PlayAgain()
	}
}
