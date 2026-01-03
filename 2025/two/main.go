package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type idRange struct {
	start int64
	end   int64
}

func getRanges() []idRange {
	content, _ := os.ReadFile("input.txt")

	ranges := []idRange{}
	for _, a := range strings.Split(string(content), ",") {
		parts := strings.Split(a, "-")
		start, _ := strconv.ParseInt(parts[0], 10, 0)
		end, _ := strconv.ParseInt(parts[1], 10, 0)
		ranges = append(ranges, idRange{start, end})
	}
	return ranges
}

func isInvalid_PartOne(num int64) bool {
	str := strconv.FormatInt(num, 10)
	strLen := len(str)
	if len(str)%2 == 1 {
		return false
	}
	mid := strLen / 2
	return str[:mid] == str[mid:]
}

func isInvalid_PartTwo(num int64) bool {
	str := strconv.FormatInt(num, 10)
	strLen := len(str)
	maxPatternLen := strLen / 2

	for patternLen := 1; patternLen <= maxPatternLen; patternLen++ {
		if strLen%patternLen != 0 {
			continue
		}

		numSections := strLen / patternLen
		var prevSection string
		for i := range numSections {
			if i == 0 {
				prevSection = str[0 : i+patternLen]
				continue
			}

			curSection := str[i*patternLen : i*patternLen+patternLen]

			// no pattern found at this length
			if curSection != prevSection {
				break
			}

			prevSection = curSection

			// pattern was found
			if i == numSections-1 {
				return true
			}
		}
	}
	return false
}

func main() {
	// SETUP
	var sum int64 = 0
	ranges := getRanges()

	// PART 1
	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			if isInvalid_PartOne(i) {
				sum += i
			}
		}
	}
	fmt.Println(sum)

	// PART 2
	sum = 0
	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			if isInvalid_PartTwo(i) {
				sum += i
			}
		}
	}
	fmt.Println(sum)
}
