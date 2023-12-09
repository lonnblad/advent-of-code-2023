package p1

import (
	"fmt"
	"regexp"
	"strings"
)

type Direction struct {
	Left  string
	Right string
}

func ParseDirections(input string) map[string]Direction {
	pattern := regexp.MustCompile(`([A-Z]+) = \(([A-Z]+), ([A-Z]+)\)`)

	matches := pattern.FindAllStringSubmatch(input, -1)
	result := make(map[string]Direction, len(matches))

	for _, match := range matches {
		result[match[1]] = Direction{Left: match[2], Right: match[3]}
	}

	return result
}

func ParseNavigation(input string) []string {
	pattern := regexp.MustCompile(`^([RL]+)`)

	matches := pattern.FindStringSubmatch(input)
	navs := strings.Split(matches[1], "")

	return navs
}

func Navigate(navigations []string, directions map[string]Direction) int {
	var steps int

	var idx int
	currentPos := "AAA"

	for {
		fmt.Println(currentPos)

		switch navigations[idx] {
		case "L":
			currentPos = directions[currentPos].Left
		case "R":
			currentPos = directions[currentPos].Right
		}

		steps++
		idx++
		if idx == len(navigations) {
			idx = 0
		}

		if currentPos == "ZZZ" {
			return steps
		}
	}
}
