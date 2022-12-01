package main

import (
	"fmt"
	"log"
	"os"
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
	var currentCalories, maxCalories, elfWithMaxCalories int
	currentElf := 1
	for _, l := range lines {
		if l == "" {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
				elfWithMaxCalories = currentElf
			}
			currentElf++
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
	fmt.Printf("Elf number %d had the most calories (%d)\n", elfWithMaxCalories, maxCalories)
}
