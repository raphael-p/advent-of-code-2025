package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func readTachyonManifold() [][]string {
	content, _ := os.ReadFile("input.txt")
	strMatrix := [][]string{}
	for line := range strings.SplitSeq(string(content), "\n") {
		chars := make([]string, 0, len(line))
		for _, r := range line {
			chars = append(chars, string(r))
		}
		strMatrix = append(strMatrix, chars)
	}
	return strMatrix
}

func getNextIndices(row []string, index int) (bool, []int) {
	hasSplit := false
	nextIndices := []int{}
	if row[index] == "^" {
		hasSplit = true
		if index > 0 {
			nextIndices = append(nextIndices, index-1)
		}
		if index+1 < len(row) {
			nextIndices = append(nextIndices, index+1)
		}
	} else {
		nextIndices = append(nextIndices, index)
	}
	return hasSplit, nextIndices
}

func findStartIndex(row []string) int {
	return slices.Index(row, "S")
}

func traverseRow(row []string, beamIndices map[int]bool) (map[int]bool, int) {
	nextBeamIndices := map[int]bool{}
	splitterHitcount := 0
	for beamIdx := range beamIndices {
		hasSplit, nextIndices := getNextIndices(row, beamIdx)
		if hasSplit {
			splitterHitcount += 1
		}
		for _, nextIndex := range nextIndices {
			nextBeamIndices[nextIndex] = true
		}
	}
	return nextBeamIndices, splitterHitcount
}

func traverseRowQuantum(row []string, beamIndices map[int]int) (map[int]int, int) {
	nextBeamIndices := map[int]int{}
	newTimelineCount := 0
	for beamIdx, timelineCount := range beamIndices {
		hasSplit, nextIndices := getNextIndices(row, beamIdx)
		if hasSplit {
			newTimelineCount += timelineCount
		}
		for _, nextIndex := range nextIndices {
			nextBeamIndices[nextIndex] = nextBeamIndices[nextIndex] + timelineCount
		}
	}
	return nextBeamIndices, newTimelineCount
}

func main() {
	// SETUP
	manifold := readTachyonManifold()
	startingColumn := findStartIndex(manifold[0])

	// PART ONE
	beamIndices := map[int]bool{startingColumn: true}
	splitCount := 0
	for _, row := range manifold[1:] {
		newBeamIndices, rowSplitCount := traverseRow(row, beamIndices)
		beamIndices = newBeamIndices
		splitCount += rowSplitCount
	}
	fmt.Println(splitCount)

	// PART TWO
	beamIndicesQuantum := map[int]int{startingColumn: 1}
	timelineCount := 1
	for _, row := range manifold[1:] {
		newBeamIndices, rowSplitCount := traverseRowQuantum(row, beamIndicesQuantum)
		beamIndicesQuantum = newBeamIndices
		timelineCount += rowSplitCount
	}
	fmt.Println(timelineCount)
}
