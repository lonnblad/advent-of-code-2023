package day06

import (
	"strconv"
	"strings"
)

func SolvePartTwoForRealInput() int {
	return SolvePartTwoForInput(realInput)
}

func SolvePartTwoForInput(input string) int {
	return parseRaceForPartTwo(input).
		calcWaysOfWinning()
}

func parseRaceForPartTwo(input string) (race race) {
	parts := strings.Split(input, "\n")

	time := strings.ReplaceAll(parts[0][5:], " ", "")
	distance := strings.ReplaceAll(parts[1][9:], " ", "")

	race.time, _ = strconv.Atoi(time)
	race.distance, _ = strconv.Atoi(distance)

	return race
}
