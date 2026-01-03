package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type bank struct {
	batteries []int64
}

func getBanks() []bank {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	banks := []bank{}
	for scanner.Scan() {
		bank := bank{}
		for _, char := range scanner.Text() {
			dig, _ := strconv.ParseInt(string(char), 10, 0)
			bank.batteries = append(bank.batteries, dig)
		}
		banks = append(banks, bank)
	}
	return banks
}

func max(arr []int64) (string, int) {
	var biggest int64
	var biggestIdx int
	for idx, num := range arr {
		if num > biggest {
			biggest = num
			biggestIdx = idx
		}
	}
	biggestStr := strconv.FormatInt(biggest, 10)
	return biggestStr, biggestIdx
}

func findBankJoltage_PartOne(b bank) int64 {
	bLen := len(b.batteries)
	max1, idx := max(b.batteries[:bLen-1])
	max2, _ := max(b.batteries[idx+1:])
	res, _ := strconv.ParseInt(max1+max2, 10, 0)
	return res
}

func findBankJoltage_PartTwo(b bank) int64 {
	bLen := len(b.batteries)
	resStr := ""
	idxFrom := 0

	for batPos := 11; batPos >= 0; batPos-- {
		max, idx := max(b.batteries[idxFrom : bLen-batPos])
		resStr += max
		idxFrom += idx + 1
	}

	res, _ := strconv.ParseInt(resStr, 10, 0)
	return res
}

func main() {
	// SETUP
	banks := getBanks()
	var sum int64 = 0

	// PART 1
	for _, bank := range banks {
		sum += findBankJoltage_PartOne(bank)
	}
	fmt.Println(sum)

	// PART 2
	sum = 0
	for _, bank := range banks {
		sum += findBankJoltage_PartTwo(bank)
	}
	fmt.Println(sum)
}
