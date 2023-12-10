package day04

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"
)

//go:embed real.input
var realInput string

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
