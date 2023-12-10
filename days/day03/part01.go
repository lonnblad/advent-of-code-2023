package day03

import (
	"strings"
)

func SolvePartOneForRealInput() int {
	_, sum := SolvePartOneForInput(realInput)
	return sum
}

func SolvePartOneForInput(input string) ([]int, int) {
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
