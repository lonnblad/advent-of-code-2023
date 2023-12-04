package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed real.input
var realInput string

func main() {
	_, sum := calculate(realInput)
	fmt.Println("Sum:", sum)
}

func calculate(input string) ([]int, int) {
	rows := strings.Split(input, "\n")
	result := []int{}
	sum := 0

	symbols := parseSymbol(rows)
	potentialPartNumbers := parsePotentialPartNumbers(rows)

	realParts := filterRealPartNumbers(potentialPartNumbers, symbols)
	for _, part := range realParts {
		result = append(result, part.number)
		sum += part.number
	}

	return result, sum
}

func parseSymbol(rows []string) map[int][]int {
	symbolCoords := map[int][]int{}

	for idx, row := range rows {
		for jdx, char := range row {
			switch char {
			case '*', '+', '=', '%', '#', '-', '@', '&', '/', '$':
				symbolCoords[idx] = append(symbolCoords[idx], jdx)
			}
		}
	}

	return symbolCoords
}

type partNumber struct {
	number          int
	row, start, end int
}

func parsePotentialPartNumbers(rows []string) []partNumber {
	parts := []partNumber{}

	pattern := regexp.MustCompile(`(\d+)`)

	for idx, row := range rows {
		var offset int

		for true {
			loc := pattern.FindIndex([]byte(row))
			if loc == nil {
				break
			}

			val, _ := strconv.Atoi(row[loc[0]:loc[1]])

			part := partNumber{number: val, row: idx, start: offset + loc[0], end: offset + loc[1] - 1}
			parts = append(parts, part)

			offset += loc[1]
			row = row[loc[1]:]
		}
	}

	return parts
}

func filterRealPartNumbers(potentialParts []partNumber, symbols map[int][]int) []partNumber {
	var realParts []partNumber

	for _, part := range potentialParts {
		prevRow := symbols[part.row-1]
		if isAdjecentToASymbol(part, prevRow) {
			realParts = append(realParts, part)
			continue
		}

		sameRow := symbols[part.row]
		if isAdjecentToASymbol(part, sameRow) {
			realParts = append(realParts, part)
			continue
		}

		nextRow := symbols[part.row+1]
		if isAdjecentToASymbol(part, nextRow) {
			realParts = append(realParts, part)
			continue
		}
	}

	return realParts
}

func isAdjecentToASymbol(part partNumber, symbols []int) bool {
	for _, symbolIdx := range symbols {
		if symbolIdx >= part.start-1 && symbolIdx <= part.end+1 {
			return true
		}
	}

	return false
}
