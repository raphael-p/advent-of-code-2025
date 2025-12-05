package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type freshRange struct {
	start, end uint64
}

type ingredients struct {
	freshRanges []freshRange
	ids         []uint64
}

func readDB() ingredients {
	content, _ := os.ReadFile("input.txt")

	var ingredients ingredients = ingredients{}
	isParsingIds := false
	for row := range strings.SplitSeq(string(content), "\n") {
		if row == "" {
			isParsingIds = true
			continue
		}

		if isParsingIds {
			id, _ := strconv.ParseUint(row, 10, 0)
			ingredients.ids = append(ingredients.ids, id)
		} else {
			split := strings.Split(row, "-")
			start, _ := strconv.ParseUint(split[0], 10, 0)
			end, _ := strconv.ParseUint(split[1], 10, 0)
			ingredients.freshRanges = append(ingredients.freshRanges, freshRange{start, end})
		}
	}

	return ingredients
}

func isIngredientFresh(id uint64, freshRanges []freshRange) bool {
	for _, r := range freshRanges {
		if id >= r.start && id <= r.end {
			return true
		}
	}
	return false
}

func main() {
	// SETUP
	ingredients := readDB()
	var sum uint64 = 0

	// PART ONE
	for _, id := range ingredients.ids {
		if isIngredientFresh(id, ingredients.freshRanges) {
			sum += 1
		}
	}
	fmt.Println(sum)

	// PART TWO
	slices.SortStableFunc(ingredients.freshRanges, func(a, b freshRange) int {
		if a.start > b.start {
			return 1
		} else if a.start < b.start {
			return -1
		} else {
			return 0
		}
	})

	normalisedRanges := []freshRange{}
	for _, r := range ingredients.freshRanges {
		rangeCount := len(normalisedRanges)

		// first iteration
		if rangeCount == 0 {
			normalisedRanges = append(normalisedRanges, r)
			continue
		}

		lastNormalised := normalisedRanges[rangeCount-1]

		// no overlap
		if r.start > lastNormalised.end {
			normalisedRanges = append(normalisedRanges, r)
			continue
		}

		// fully enclosed
		if r.end <= lastNormalised.end {
			continue
		}

		// consolidate
		normalisedRanges[rangeCount-1].end = r.end
	}

	sum = 0
	for _, r := range normalisedRanges {
		sum += r.end - r.start + 1
	}
	fmt.Println(sum)
}
