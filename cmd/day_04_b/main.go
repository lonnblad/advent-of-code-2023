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
