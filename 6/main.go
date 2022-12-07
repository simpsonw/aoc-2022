package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var patternLength int

func main() {
	patternLength, _ = strconv.Atoi(os.Args[1])
	stream := os.Args[2]

	for k, _ := range stream {
		if k >= len(stream)+patternLength {
			fmt.Printf("Could not find start of message marker")
			break
		} else if isStartOfPacketMarker(stream[k : k+patternLength]) {
			fmt.Printf("First marker is after character %d\n", k+patternLength)
			break
		}
	}
}

func isStartOfPacketMarker(candidate string) bool {
	var sb strings.Builder
	for _, v := range candidate {
		if strings.ContainsRune(sb.String(), v) {
			return false
		}
		sb.WriteRune(v)
	}
	return true
}
