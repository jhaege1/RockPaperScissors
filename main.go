package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var user string
	var computer string
	var score int
	var cpuScore int

	score = 0
	cpuScore = 0

	for {
		computer = getRandomChoice()

		fmt.Println("What's your choice?\n'r' for rock, 'p' for paper, 's' for scissors")
		fmt.Scanln(&user)

		if user == computer {
			fmt.Printf("It's a tie!\nThe score is %v - %v\n", score, cpuScore)
		} else if win(user, computer) {
			score = score + 1
			fmt.Printf("You won!\nThe score is %v - %v\n", score, cpuScore)
		} else {
			cpuScore = cpuScore + 1
			fmt.Printf("You lost!\nThe score is %v - %v\n", score, cpuScore)
		}
	}
}

func getRandomChoice() string {
	var computer string

	choices := []string{"r", "p", "s"}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(choices))
	computer = choices[index]

	return computer
}

func win(p string, o string) bool {
	var boolean bool

	if (p == "r" && o == "s") || (p == "s" && o == "p") || (p == "p" && o == "r") {
		boolean = true
	}

	return boolean
}
