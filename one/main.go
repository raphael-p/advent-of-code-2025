package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func getRotations() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := []string{}
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func parseRotation(rotation []byte) int {
	sign := 1
	if rotation[0] == 'L' {
		sign = -1
	}
	mag, _ := strconv.ParseInt(string(rotation[1:]), 10, 0)
	return sign * int(mag)
}

func modulus(a, mod int) int {
	return (a%mod + mod) % mod
}

func main() {
	// SETUP
	zeroCounter := 0
	position := 50
	MOD := 100
	rotations := getRotations()

	// PART 1
	for _, rotation := range rotations {
		op := parseRotation([]byte(rotation))
		position = modulus(position+op, MOD)
		if position == 0 {
			zeroCounter = zeroCounter + 1
		}
	}
	fmt.Println(zeroCounter)

	// PART 2
	zeroCounter = 0
	position = 50
	for _, rotation := range rotations {
		op := parseRotation([]byte(rotation))

		// changes due to magnitude
		magnitude := int(math.Abs(float64(op))) / MOD
		zeroCounter += magnitude

		// changes due to position (catch 1st round)
		oldPosition := position
		position = modulus(position+op, MOD)

		swing := (op % MOD) + oldPosition
		// case 1: already at 0
		if oldPosition == 0 {
			// case 2: goes over
		} else if swing >= MOD {
			zeroCounter += 1
			// case 3: goes under
		} else if swing <= 0 {
			zeroCounter += 1
		}
	}
	fmt.Println(zeroCounter)
}
