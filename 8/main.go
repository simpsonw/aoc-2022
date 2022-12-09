package main

import (
	"fmt"
	"github.com/simpsonw/aoc-2022/utils"
	"strconv"
)

type Tree struct {
	Row       int
	Col       int
	Height    int
	isVisible bool
}

func (t Tree) String() string {
	return fmt.Sprintf("%d(%t)", t.Height, t.isVisible)
}

func main() {
	lines := utils.GetLines()
	forest := readTrees(lines)
	visibleTrees := 0
	for _, v := range forest {
		for _, j := range v {
			if isVisible(j, forest) {
				visibleTrees++
			}
		}
	}
	fmt.Printf("%d trees were visible\n", visibleTrees)
}

func isVisible(tree *Tree, forest [][]*Tree) bool {
	// The tree is on the edge of the forest, so we marked it visible when we read input
	if tree.isVisible {
		return true
	}
	tree.isVisible = visibleNorth(tree, forest) ||
		visibleSouth(tree, forest) ||
		visibleWest(tree, forest) ||
		visibleEast(tree, forest)
	return tree.isVisible
}

func visibleEast(tree *Tree, forest [][]*Tree) bool {
	var row = tree.Row
	var col = tree.Col + 1
	for col < len(forest[row]) {
		if forest[row][col].Height >= tree.Height {
			return false
		}
		col++
	}
	return true
}

func visibleWest(tree *Tree, forest [][]*Tree) bool {
	var row = tree.Row
	var col = tree.Col - 1
	for col >= 0 {
		if forest[row][col].Height >= tree.Height {
			return false
		}
		col--
	}
	return true
}

func visibleSouth(tree *Tree, forest [][]*Tree) bool {
	var row = tree.Row + 1
	var col = tree.Col
	for row < len(forest) {
		if forest[row][col].Height >= tree.Height {
			return false
		}
		row++
	}
	return true
}

func visibleNorth(tree *Tree, forest [][]*Tree) bool {
	var row = tree.Row - 1
	var col = tree.Col
	for row >= 0 {
		if forest[row][col].Height >= tree.Height {
			return false
		}
		row--
	}
	return true
}

func readTrees(lines []string) [][]*Tree {
	var forest [][]*Tree
	for row, v := range lines {
		if v == "" {
			break
		}
		forest = append(forest, []*Tree{})
		for col, j := range v {
			height, err := strconv.Atoi(string(j))
			if err != nil {
				panic(err)
			}
			isVisible := row == 0 || col == 0 || row == len(lines)-2 || col == len(v)-1
			tree := Tree{
				Row:       row,
				Col:       col,
				Height:    height,
				isVisible: isVisible,
			}
			forest[row] = append(forest[row], &tree)
		}
	}
	return forest
}
