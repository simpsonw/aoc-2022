package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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
	var calories []int
	var currentCalories int
	for _, l := range lines {
		if l == "" {
			calories = append(calories, currentCalories)
			currentCalories = 0
			continue
		}
		itemCalories, err := strconv.Atoi(l)
		if err != nil {
			fmt.Printf("Error parsing %s: %s\n", l, err)
			break
		}
		currentCalories += itemCalories
	}
	var totalCalories int
	sort.Ints(calories)
	for i := len(calories) - 3; i < len(calories); i++ {
		totalCalories += calories[i]
	}
	fmt.Printf("Total calories of the top three elves: %d\n", totalCalories)
}
