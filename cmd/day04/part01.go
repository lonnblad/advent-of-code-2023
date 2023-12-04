package day04

import (
	"strings"
)

func SolvePartOneForRealInput() int {
	_, sum := SolvePartOneForInput(realInput)
	return sum
}

func SolvePartOneForInput(input string) ([]int, int) {
	rows := strings.Split(input, "\n")
	result := make([]int, len(rows))
	sum := 0

	for idx, row := range rows {
		winning := parseWinningNumbers(row)
		youHave := parseYourNumbers(row)

		result[idx] = calculateScore(winning, youHave)
		sum += result[idx]
	}

	return result, sum
}

func calculateScore(winning map[int]bool, youHave []int) int {
	var score int

	for _, no := range youHave {
		if winning[no] {
			if score == 0 {
				score = 1
			} else {
				score = score * 2
			}
		}
	}

	return score
}
