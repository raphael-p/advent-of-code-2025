package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct{ x, y int }

func getPositions(file string) []position {
	content, _ := os.ReadFile(file)
	positions := []position{}
	for line := range strings.SplitSeq(string(content), "\n") {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		positions = append(positions, position{x, y})
	}
	return positions
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calcArea(p1, p2 position) int {
	dx := abs(p1.x-p2.x) + 1
	dy := abs(p1.y-p2.y) + 1
	return dx * dy
}

func getCorners(p1, p2 position) (position, position, position, position) {
	var greatestX, greatestY, smallestX, smallestY int
	if p1.x >= p2.x {
		greatestX = p1.x
		smallestX = p2.x
	} else {
		greatestX = p2.x
		smallestX = p1.x
	}
	if p1.y >= p2.y {
		greatestY = p1.y
		smallestY = p2.y
	} else {
		greatestY = p2.y
		smallestY = p1.y
	}
	// top-left, top-right, bottom-left, bottom-right
	return position{smallestX, smallestY}, position{greatestX, smallestY}, position{smallestX, greatestY}, position{greatestX, greatestY}
}

func validateVirtualPoint(v position, positions []position) (bool, bool, bool, bool) {
	var isValidTL, isValidTR, isValidBL, isValidBR bool
	for _, p := range positions {
		if p.y <= v.y && p.x <= v.x {
			isValidTL = true
		}
		if p.y <= v.y && p.x >= v.x {
			isValidTR = true
		}
		if p.y >= v.y && p.x <= v.x {
			isValidBL = true
		}
		if p.y >= v.y && p.x >= v.x {
			isValidBR = true
		}
	}
	return isValidTL, isValidTR, isValidBL, isValidBR
}

func isVirtual(p, p1, p2 position) bool {
	if p.x == p1.x && p.y == p1.y {
		return false
	}
	if p.x == p2.x && p.y == p2.y {
		return false
	}

	return true
}

func isSquarePossible(i, j int, positions []position) bool {
	p1, p2 := positions[i], positions[j]
	tl, tr, bl, br := getCorners(p1, p2)

	if isVirtual(tl, p1, p2) {
		isValidTL, _, _, _ := validateVirtualPoint(tl, positions)
		if !isValidTL {
			return false
		}
	}

	if isVirtual(tr, p1, p2) {
		_, isValidTR, _, _ := validateVirtualPoint(tr, positions)
		if !isValidTR {
			return false
		}
	}

	if isVirtual(bl, p1, p2) {
		_, _, isValidBL, _ := validateVirtualPoint(bl, positions)
		if !isValidBL {
			return false
		}
	}

	if isVirtual(br, p1, p2) {
		_, _, _, isValidBR := validateVirtualPoint(br, positions)
		if !isValidBR {
			return false
		}
	}

	return true
}

func main() {
	// SETUP
	positions := getPositions("test.txt")
	n := len(positions)

	// PART ONE
	maxArea := 0
	for i := range n - 1 {
		for j := i + 1; j < n; j++ {
			area := calcArea(positions[i], positions[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}
	fmt.Println(maxArea)

	// PART TWO
	maxArea = 0
	for i := range n - 1 {
		for j := i + 1; j < n; j++ {
			fmt.Printf("p1: %v, p2: %v\n", positions[i], positions[j])
			if !isSquarePossible(i, j, positions) {
				fmt.Print(" --> not possible\n")
				continue
			}
			fmt.Print(" --> possible\n")
			area := calcArea(positions[i], positions[j])

			// // DEBUG
			// if positions[i].x == 11 && positions[i].y == 7 {
			// 	fmt.Printf("area: %d", area)
			// 	fmt.Printf(" p1: %v, p2: %v\n", positions[i], positions[j])
			// }

			if area > maxArea {
				maxArea = area
			}
		}
	}
	fmt.Println(maxArea) // this doesn't work, I give up

}
