package main

import (
	"fmt"
	"os"
	"strings"
)

func readGrid() [][]bool {
	content, _ := os.ReadFile("input.txt")

	var grid [][]bool = [][]bool{}
	for row := range strings.SplitSeq(string(content), "\n") {
		rowArr := []bool{}
		for _, obj := range row {
			rowArr = append(rowArr, obj == '@')
		}
		grid = append(grid, rowArr)

	}
	return grid
}

func isRollHere(grid [][]bool, i, j int) int {
	if i < 0 || i >= len(grid) {
		return 0
	}
	if j < 0 || j >= len(grid[0]) {
		return 0
	}

	if grid[i][j] {
		return 1
	}
	return 0
}

func isRollAccessible(grid [][]bool, i, j int) bool {
	if isRollHere(grid, i, j) == 0 {
		return false
	}

	count := 0
	count += isRollHere(grid, i-1, j-1)
	count += isRollHere(grid, i-1, j)
	count += isRollHere(grid, i-1, j+1)
	count += isRollHere(grid, i, j-1)
	count += isRollHere(grid, i, j+1)
	count += isRollHere(grid, i+1, j+1)
	count += isRollHere(grid, i+1, j)
	count += isRollHere(grid, i+1, j-1)

	return count < 4
}

func main() {
	// SETUP
	grid := readGrid()
	sum := 0

	// PART ONE
	for i, r := range grid {
		for j := range r {
			increment := isRollAccessible(grid, i, j)
			if increment {
				sum += 1
			}
		}
	}
	fmt.Println(sum)

	// PART TWO
	sum = 0
	for gridChanged := true; gridChanged; {
		gridChanged = false
		for i, r := range grid {
			for j := range r {
				increment := isRollAccessible(grid, i, j)
				if increment {
					gridChanged = true
					grid[i][j] = false
					sum += 1
				}
			}
		}
	}
	fmt.Println(sum)
}
