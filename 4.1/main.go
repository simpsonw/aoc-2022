package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ElfRange struct {
	start int
	end   int
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
	var numOverlappingRanges = 0
	for _, l := range lines {
		if l == "" {
			break
		}
		var ranges, err = parseRangeFromLine(l)
		if err != nil {
			log.Fatalf("%e", err)
		}
		if rangesOverlap(ranges[0], ranges[1]) {
			numOverlappingRanges++
		}
	}
	fmt.Printf("The number of overlapping ranges was %d\n", numOverlappingRanges)
}

func rangesOverlap(a, b ElfRange) bool {
	return (a.start <= b.start && a.end >= b.end) || (b.start <= a.start && b.end >= a.end)
}

func parseRangeFromLine(l string) (ranges []ElfRange, err error) {
	var commaValues = strings.Split(l, ",")
	for _, v := range commaValues {
		var er ElfRange
		er, err = getElfRangeFromString(v)
		if err != nil {
			return
		}
		ranges = append(ranges, er)
	}
	return
}

func getElfRangeFromString(s string) (er ElfRange, err error) {
	parsedRange := strings.Split(s, "-")
	if len(parsedRange) != 2 {
		err = fmt.Errorf("Could not parse range from %q", s)
		return
	}
	start, err := strconv.Atoi(parsedRange[0])
	if err != nil {
		return
	}
	er.start = start
	end, err := strconv.Atoi(parsedRange[1])
	if err != nil {
		return
	}
	er.end = end
	return
}
