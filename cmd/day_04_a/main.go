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

func parseWinningNumbers(row string) map[int]bool {
	numbers := strings.Split(row, ":")[1]
	winning := strings.Split(numbers, "|")[0]

	pattern := regexp.MustCompile(`(\d+)`)
	matches := pattern.FindAllStringSubmatch(winning, -1)

	result := make(map[int]bool, len(matches))
	for _, match := range matches {
		val, _ := strconv.Atoi(match[0])
		result[val] = true
	}

	return result
}

func parseYourNumbers(row string) []int {
	numbers := strings.Split(row, ":")[1]
	yourNumbers := strings.Split(numbers, "|")[1]

	pattern := regexp.MustCompile(`(\d+)`)
	matches := pattern.FindAllStringSubmatch(yourNumbers, -1)

	result := make([]int, len(matches))
	for idx, match := range matches {
		val, _ := strconv.Atoi(match[0])
		result[idx] = val
	}

	return result
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
