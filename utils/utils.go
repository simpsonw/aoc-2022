package utils

import (
	"log"
	"os"
	"strings"
)

func GetLines() []string {
	if len(os.Args) < 1 {
		log.Fatalf("Usage: %s <filename> \n", os.Args[0])
	}
	filename := os.Args[1]
	var content, err = os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	return lines
}
