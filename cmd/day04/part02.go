package day04

import (
	"regexp"
	"strconv"
	"strings"
)

func SolvePartTwoForRealInput() int {
	_, sum := SolvePartTwoForInput(realInput)
	return sum
}

func SolvePartTwoForInput(input string) ([]int, int) {
	rows := strings.Split(input, "\n")
	result := make([]int, len(rows))
	sum := 0

	copies := map[int]int{}

	for _, row := range rows {
		gameNo := parseGameNo(row)

		winning := parseWinningNumbers(row)
		youHave := parseYourNumbers(row)

		var score int

		for _, no := range youHave {
			if winning[no] {
				score += 1
				copies[gameNo+score] = copies[gameNo+score] + 1 + 1*copies[gameNo]
			}
		}
	}

	for idx := range result {
		result[idx] = copies[idx+1] + 1
		sum += result[idx]
	}

	return result, sum
}

func parseGameNo(row string) int {
	gameNo := strings.Split(row, ":")[0]

	pattern := regexp.MustCompile(`(\d+)`)
	match := pattern.FindStringSubmatch(gameNo)

	val, _ := strconv.Atoi(match[0])

	return val
}
