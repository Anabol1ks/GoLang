package main

import "fmt"

func Rps(p1, p2 string) (win string) {
	switch {
	case p1 == "scissors" && p2 == "paper":
		win = "Player 1 won!"
	case p2 == "scissors" && p1 == "paper":
		win = "Player 2 won!"
	case p1 == "scissors" && p2 == "rock":
		win = "Player 2 won!"
	case p2 == "scissors" && p1 == "rock":
		win = "Player 1 won!"
	case p2 == "paper" && p1 == "rock":
		win = "Player 2 won!"
	case p1 == "paper" && p2 == "rock":
		win = "Player 1 won!"
	default:
		win = "Draw!"
	}
	return
}

func main() {
	fmt.Println(Rps("rock", "scissors"))
}
