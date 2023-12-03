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
	sum := 0

	symbols := parseSymbolCoords(rows)
	potentialGearRatioParts := parsePotentialPartNumbers(rows)

	realGears := findRealGears(potentialGearRatioParts, symbols)
	for _, part := range realGears {
		sum += part
	}

	return realGears, sum
}

func parseSymbolCoords(rows []string) map[int][]int {
	symbolCoords := map[int][]int{}

	for idx, row := range rows {
		for jdx, char := range row {
			switch char {
			case '*':
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

func findRealGears(potentialGearRatioParts []partNumber, symbols map[int][]int) []int {
	var result []int

	for symbolRowNo, symbols := range symbols {
		for _, symbolIdx := range symbols {
			var gearRatioPartOne int

			for _, gearRatioPart := range potentialGearRatioParts {
				if !(gearRatioPart.row >= symbolRowNo-1 && gearRatioPart.row <= symbolRowNo+1) {
					continue
				}

				if !(symbolIdx >= gearRatioPart.start-1 && symbolIdx <= gearRatioPart.end+1) {
					continue
				}

				if gearRatioPartOne == 0 {
					gearRatioPartOne = gearRatioPart.number
					continue
				}

				result = append(result, gearRatioPartOne*gearRatioPart.number)

				break
			}
		}
	}

	return result
}
