package day03

import (
	"strings"
)

func SolvePartTwoForRealInput() int {
	_, sum := SolvePartTwoForInput(realInput)
	return sum
}

func SolvePartTwoForInput(input string) ([]int, int) {
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
