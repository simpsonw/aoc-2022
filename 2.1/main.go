package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

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
		totalScore += getScore(round[0], round[1]) + getShapeValue(round[1])
	}
	fmt.Printf("Your total score was %d\n", totalScore)
}

func getScore(opponentMove, playerMove string) int {
	var score = 0
	var tieMoves = map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}
	if tieMoves[playerMove] == opponentMove {
		score = 3
	}
	var winningMoves = map[string]string{
		"X": "C", // Rock beats scissors
		"Y": "A", // Paper beats rock
		"Z": "B", // Scissors beats paper
	}
	if winningMoves[playerMove] == opponentMove {
		score = 6
	}

	return score
}

func getShapeValue(s string) int {
	switch s {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	}
	return -1
}
