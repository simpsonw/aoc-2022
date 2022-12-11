package main

import (
	"fmt"
	"github.com/simpsonw/aoc-2022/utils"
	"strconv"
	"strings"
)

var X = 1
var cycle = 1
var nextMeasuringCycle = 20
var sum int
var currentRow int
var crt [][]rune

func main() {
	lines := utils.GetLines()

	for _, l := range lines {
		if l == "" {
			break
		}
		split := strings.Split(l, " ")
		//Either noop or first cycle of addx
		drawPixel(cycle, X)
		cycle++
		measureSignalStrength(cycle)
		if len(split) == 2 {
			//addx
			value, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}

			drawPixel(cycle, X)
			X += value
			cycle++
			measureSignalStrength(cycle)
		}
	}
	fmt.Println()
	//fmt.Printf("Sum of signal strengths: %d\n", sum)
}
func drawPixel(cycle, X int) {
	col := cycle % 40
	col--
	if col == 0 && cycle > 1 {
		fmt.Println()
	}
	if col == X || col-1 == X || col+1 == X {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}
}

func measureSignalStrength(cycle int) {
	if cycle == nextMeasuringCycle {
		strength := cycle * X
		sum += strength
		nextMeasuringCycle += 40
		//fmt.Printf("\tSignal strength: %d (cycle: %d X: %d)\n", strength, cycle, X)
	}
}
