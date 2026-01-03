package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type position struct{ x, y, z int }

func getPositions(file string) []position {
	content, _ := os.ReadFile(file)
	positions := []position{}
	for line := range strings.SplitSeq(string(content), "\n") {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		positions = append(positions, position{x, y, z})
	}
	return positions
}

func calcDistance(p1, p2 position) int {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	dz := p1.z - p2.z
	return dx*dx + dy*dy + dz*dz
}

type edges struct {
	indexA   int
	indexB   int
	distance int
}

func createEdges(positions []position) []edges {
	connections := []edges{}
	n := len(positions)
	for i := range n - 1 {
		for y := i + 1; y < n; y++ {
			d := calcDistance(positions[i], positions[y])
			connections = append(connections, edges{i, y, d})
		}
	}
	slices.SortStableFunc(connections, func(a, b edges) int { return a.distance - b.distance })
	return connections
}

func searchExistingCircuits(circuits [][]int, newConnection edges) (int, int) {
	circuitIndexA := -1
	circuitIndexB := -1
	for circuitIndex, c := range circuits {
		if slices.Contains(c, newConnection.indexA) {
			circuitIndexA = circuitIndex
		}
		if slices.Contains(c, newConnection.indexB) {
			circuitIndexB = circuitIndex
		}
		if circuitIndexA != -1 && circuitIndexB != -1 {
			break
		}
	}
	return circuitIndexA, circuitIndexB
}

func addConnection(circuits [][]int, newConnection edges) [][]int {
	circuitIndexA, circuitIndexB := searchExistingCircuits(circuits, newConnection)

	// case: neither are in a circuit -> add a new circuit
	if circuitIndexA == -1 && circuitIndexB == -1 {
		return append(circuits, []int{newConnection.indexA, newConnection.indexB})
	}

	// case: they are already in the same circuit -> noop
	if circuitIndexA == circuitIndexB {
		return circuits
	}

	// case: both already in different circuits -> combine the circuits
	if circuitIndexA != -1 && circuitIndexB != -1 {
		circuits[circuitIndexA] = append(circuits[circuitIndexA], circuits[circuitIndexB]...)
		return append(circuits[:circuitIndexB], circuits[circuitIndexB+1:]...)
	}

	// cases: one is in a circuit but not the other -> add the other to the firsts' circuit
	if circuitIndexB == -1 {
		circuits[circuitIndexA] = append(circuits[circuitIndexA], newConnection.indexB)
	} else if circuitIndexA == -1 {
		circuits[circuitIndexB] = append(circuits[circuitIndexB], newConnection.indexA)
	}
	return circuits
}

func main() {
	// SETUP
	positions := getPositions("input.txt")
	connections := createEdges(positions)
	maxConnections := 1000

	// PART ONE
	circuits := [][]int{}
	for _, conn := range connections[:maxConnections] {
		circuits = addConnection(circuits, conn)
	}
	slices.SortFunc(circuits, func(a, b []int) int { return len(b) - len(a) })
	fmt.Println(len(circuits[0]) * len(circuits[1]) * len(circuits[2]))

	// PART TWO
	for _, conn := range connections[maxConnections:] {
		circuits = addConnection(circuits, conn)
		if len(circuits) == 1 && len(circuits[0]) == len(positions) {
			fmt.Println(positions[conn.indexA].x * positions[conn.indexB].x)
			break
		}
	}
}
