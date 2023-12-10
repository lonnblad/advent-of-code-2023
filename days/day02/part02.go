package day02

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func SolvePartTwoForRealInput() int {
	_, sum := SolvePartTwoForInput(realInput)
	return sum
}

func SolvePartTwoForInput(input string) ([]int, int) {
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
