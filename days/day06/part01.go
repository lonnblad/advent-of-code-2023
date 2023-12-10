package day06

import (
	"regexp"
	"strconv"
	"strings"
)

func SolvePartOneForRealInput() int {
	_, sum := SolvePartOneForInput(realInput)
	return sum
}

func SolvePartOneForInput(input string) ([]int, int) {
	races := parseRacesForPartOne(input)

	result := make([]int, len(races))
	sum := 1

	for idx, race := range races {
		result[idx] = race.calcWaysOfWinning()
		sum *= result[idx]
	}

	return result, sum
}

func parseRacesForPartOne(input string) []race {
	parts := strings.Split(input, "\n")

	pattern := regexp.MustCompile(`(\d+)`)

	times := pattern.FindAllStringSubmatch(parts[0], -1)
	result := make([]race, len(times))

	for idx, time := range times {
		result[idx].time, _ = strconv.Atoi(time[0])
	}

	distances := pattern.FindAllStringSubmatch(parts[1], -1)
	for idx, distance := range distances {
		result[idx].distance, _ = strconv.Atoi(distance[0])
	}

	return result
}
