package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
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
	var prioritySum = 0
	var elfGroup []string
	for _, l := range lines {
		if l == "" {
			break
		}
		elfGroup = append(elfGroup, l)
		if len(elfGroup) == 3 {
			badge, err := findBadge(elfGroup[0], elfGroup[1], elfGroup[2])
			if err != nil {
				fmt.Printf("%e", err)
				break
			}
			prioritySum += getPriority(badge)
			elfGroup = []string{}
		}
		// Uncomment for solution 1
		//rucksack1, rucksack2 := splitRucksack(l)
		//commonCharacter, err := findCommomCharacter(rucksack1, rucksack2)
		//if err != nil {
		//	fmt.Printf("%e", err)
		//	break
		//}
		//prioritySum += getPriority(commonCharacter)
	}

	fmt.Printf("Sum of priorities was %d\n", prioritySum)
}

func getPriority(r rune) int {
	if unicode.IsUpper(r) {
		return int(r - 38)
	} else {
		return int(r - 96)
	}
}

func splitRucksack(line string) (string, string) {
	var length = len(line)
	return line[:(length / 2)], line[(length / 2):]
}

func findBadge(a, b, c string) (rune, error) {
	for _, v := range a {
		if strings.Contains(b, string(v)) && strings.Contains(c, string(v)) {
			return v, nil
		}
	}
	return ' ', fmt.Errorf("could not find a common character between %q and %q", a, b)
}

func findCommomCharacter(a, b string) (rune, error) {
	for _, v := range a {
		if strings.Contains(b, string(v)) {
			return v, nil
		}
	}

	return ' ', fmt.Errorf("could not find a common character between %q and %q", a, b)
}
