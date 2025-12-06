package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func load() string {
	content, _ := os.ReadFile("input.txt")
	return string(content)
}

type humanData struct {
	colCount  int
	rowCount  int
	numbers   [][]int64
	operators []string
}

func readForHumans() humanData {
	data := humanData{}
	for line := range strings.SplitSeq(load(), "\n") {
		fields := strings.Fields(line)

		if fields[0] == "*" || fields[0] == "+" {
			data.operators = fields
			break
		}

		numArr := []int64{}
		for _, str := range fields {
			num, _ := strconv.ParseInt(str, 10, 0)
			numArr = append(numArr, num)
		}
		data.numbers = append(data.numbers, numArr)
	}

	data.rowCount = len(data.numbers)
	data.colCount = len(data.numbers[0])
	return data
}

func readForCephalopod() ([][]string, int, int) {
	strMatrix := [][]string{}
	for line := range strings.SplitSeq(load(), "\n") {
		chars := make([]string, 0, len(line))
		for _, r := range line {
			chars = append(chars, string(r))
		}
		strMatrix = append(strMatrix, chars)
	}
	return strMatrix, len(strMatrix), len(strMatrix[0])
}

func operationOnCol(column []int64, operator string) int64 {
	var columnTotal int64
	if operator == "*" {
		columnTotal = 1
	}
	for _, val := range column {
		if operator == "+" {
			columnTotal += val
		} else {
			columnTotal *= val
		}
	}
	return columnTotal
}

func main() {
	// PART ONE
	data := readForHumans()
	var sum int64
	for j := range data.colCount {
		var col []int64
		for i := range data.rowCount {
			col = append(col, data.numbers[i][j])
		}
		sum += operationOnCol(col, data.operators[j])
	}
	fmt.Println(sum)

	// PART TWO
	sum = 0
	strMatrix, rowCount, colCount := readForCephalopod()
	var operator string
	var col []int64
	for j := range colCount {
		strVal := ""
		for i := range rowCount {
			val := strMatrix[i][j]
			if val == "*" || val == "+" {
				operator = val
			} else if val != " " {
				strVal += val
			}
		}

		isAllSpace := strVal == ""
		if !isAllSpace {
			num, _ := strconv.ParseInt(strVal, 10, 0)
			col = append(col, num)
		}

		if isAllSpace || j == colCount-1 {
			sum += operationOnCol(col, operator)

			// // reset
			operator = ""
			strVal = ""
			col = col[:0]
		}
	}
	fmt.Println(sum)
}
