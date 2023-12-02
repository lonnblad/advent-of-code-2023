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

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

func calculate(input string) ([]bool, int) {
	games := strings.Split(input, "\n")
	result := make([]bool, len(games))
	sum := 0

	for idx, game := range games {
		gameNo := parseGameNo(game)

		if !isColorBelowLimit(game, "red", maxRed) {
			result[idx] = false
			continue
		}

		if !isColorBelowLimit(game, "green", maxGreen) {
			result[idx] = false
			continue
		}

		if !isColorBelowLimit(game, "blue", maxBlue) {
			result[idx] = false
			continue
		}

		result[idx] = true
		sum += gameNo
	}

	return result, sum
}

func parseGameNo(game string) int {
	re := regexp.MustCompile(`Game (\d+)`)

	match := re.FindStringSubmatch(game)[1]
	gameNo, _ := strconv.Atoi(match)

	return gameNo
}

func isColorBelowLimit(game, color string, limit int) bool {
	re := regexp.MustCompile(fmt.Sprintf(`(\d+) %s`, color))

	matches := re.FindAllStringSubmatch(game, -1)
	for _, match := range matches {
		val, _ := strconv.Atoi(match[1])

		if val > limit {
			return false
		}
	}

	return true
}
