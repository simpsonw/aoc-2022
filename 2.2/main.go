package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Outcome struct {
	win  string
	lose string
}

// A = Rock
// B = Paper
// C = Scissors

var outcomes = map[string]Outcome{
	"A": { // Opponent plays rock
		win:  "B", // Paper beats rock
		lose: "C", // Rock beats scissors
	},
	"B": {
		win:  "C", // Scissors beats paper
		lose: "A", // Paper beats rock
	},
	"C": {
		win:  "A", // Rock beats scissors
		lose: "B", // Scissors beats paper
	},
}

func main() {
	if len(os.Args) < 1 {
		log.Fatalf("Usage: %s <filename> \n", os.Args[0])
	}
	filename := os.Args[1]
	var content, err = os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	var totalScore int
	for _, l := range lines {
		if l == "" {
			break
		}
		var round = strings.Split(l, " ")
		totalScore += getScore(round[0], round[1])
	}
	fmt.Printf("Your total score was %d\n", totalScore)
}

func getScore(opponentMove, desiredResult string) int {
	var score = 0
	var playerMove = opponentMove

	if desiredResult == "X" {
		playerMove = outcomes[opponentMove].lose
	} else if desiredResult == "Z" {
		playerMove = outcomes[opponentMove].win
		score += 6
	} else if desiredResult == "Y" {
		score += 3
	}

	return score + getShapeValue(playerMove)
}

func getShapeValue(s string) int {
	switch s {
	case "A":
		return 1
	case "B":
		return 2
	case "C":
		return 3
	}
	return 0
}
