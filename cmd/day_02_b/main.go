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
	games := strings.Split(input, "\n")
	result := make([]int, len(games))
	sum := 0

	for idx, game := range games {
		red := findMinForColor(game, "red")
		green := findMinForColor(game, "green")
		blue := findMinForColor(game, "blue")

		result[idx] = red * green * blue
		sum += result[idx]
	}

	return result, sum
}

func parseGameNo(game string) int {
	re := regexp.MustCompile(`Game (\d+)`)

	match := re.FindStringSubmatch(game)[1]
	gameNo, _ := strconv.Atoi(match)

	return gameNo
}

func findMinForColor(game, color string) int {
	re := regexp.MustCompile(fmt.Sprintf(`(\d+) %s`, color))
	matches := re.FindAllStringSubmatch(game, -1)

	var min int

	for _, match := range matches {
		val, _ := strconv.Atoi(match[1])

		if val > min {
			min = val
		}
	}

	return min
}
