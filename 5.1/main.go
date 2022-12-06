package main

import (
	"fmt"
	"github.com/simpsonw/aoc-2022/utils"
	"regexp"
	"strings"
)

var columns [][]string

func main() {
	lines := utils.GetLines()
	var startedMoveInstructions = false
	for _, l := range lines {
		if l == "" {
			continue
		}
		if !startedMoveInstructions && isNumberLine(l) {
			startedMoveInstructions = true
			continue
		}
		if startedMoveInstructions {
			//moveCrates(l)
			moveCrates9001(l)
		} else {
			scanCrates(l)
		}
	}

	fmt.Println("Crates on the tops of stacks:")

	for _, v := range columns {
		fmt.Printf("%s,", v[len(v)-1])
	}

	fmt.Println("")
}

func printColumns() {
	for k, v := range columns {
		fmt.Printf("Column %d:", k+1)
		for _, v1 := range v {
			fmt.Printf(" %s", v1)
		}
		fmt.Println("")
	}
}

func moveCrates9001(l string) {
	var qty, to, from int
	_, err := fmt.Sscanf(l, "move %d from %d to %d", &qty, &from, &to)
	if err != nil {
		panic(err)
	}
	from--
	to--
	var fromIndex = len(columns[from]) - qty
	columns[to] = append(columns[to], columns[from][fromIndex:]...)
	columns[from] = columns[from][:fromIndex]
}

func moveCrates(l string) {
	var qty, to, from int
	_, err := fmt.Sscanf(l, "move %d from %d to %d", &qty, &from, &to)
	if err != nil {
		panic(err)
	}
	from--
	to--
	var fromIndex = len(columns[from]) - qty
	for i := len(columns[from]) - 1; i >= fromIndex; i-- {
		columns[to] = append(columns[to], columns[from][i])
	}
	columns[from] = columns[from][:fromIndex]
}

func isNumberLine(l string) bool {
	var r = regexp.MustCompile(`(\s[0-9]\s)(\s)?`)
	var matches = r.FindAllString(l, -1)
	for k, v := range matches {
		matches[k] = strings.TrimSpace(v)
	}
	return len(matches) > 0
}

func scanCrates(l string) {
	var r = regexp.MustCompile(`([^0-9]{3})(\s)?`)
	var matches = r.FindAllString(l, -1)
	if len(columns) == 0 {
		columns = make([][]string, len(matches))
	}
	for k, v := range matches {
		var crate = strings.TrimSpace(v)
		if crate != "" {
			columns[k] = append([]string{crate}, columns[k]...)
		}
	}
}
