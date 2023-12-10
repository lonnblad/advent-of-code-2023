package p2

import (
	"regexp"
	"strings"
)

type Direction struct {
	Left  string
	Right string
}

func ParseDirections(input string) ([]string, map[string]Direction) {
	pattern := regexp.MustCompile(`([0-9A-Z]{3}) = \(([0-9A-Z]{3}), ([0-9A-Z]{3})\)`)

	matches := pattern.FindAllStringSubmatch(input, -1)

	startPositions := []string{}
	directions := make(map[string]Direction, len(matches))

	for _, match := range matches {
		directions[match[1]] = Direction{Left: match[2], Right: match[3]}

		if strings.HasSuffix(match[1], "A") {
			startPositions = append(startPositions, match[1])
		}
	}

	return startPositions, directions
}

func ParseNavigation(input string) []string {
	pattern := regexp.MustCompile(`^([RL]+)`)

	matches := pattern.FindStringSubmatch(input)
	navs := strings.Split(matches[1], "")

	return navs
}

func Navigate(navigations []string, startPosition string, directions map[string]Direction) int {
	var steps int

	var idx int
	currentPos := startPosition

	for {
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

		if currentPos[2:] == "Z" {
			return steps
		}
	}
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
