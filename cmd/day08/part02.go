package day08

import (
	"regexp"
	"strings"
)

func SolvePartTwoForRealInput() int {
	sum := SolvePartTwoForInput(realInput)
	return sum
}

func SolvePartTwoForInput(input string) int {
	navs := ParseNavigation(input)
	startPositions, directions := ParseDirections(input)

	result := make([]int, len(startPositions))
	var sum int
	for idx, startPos := range startPositions {
		result[idx] = Navigate(navs, startPos, directions)
	}

	sum = LCM(result[0], result[1], result[2:]...)
	// fmt.Println(startPositions)
	// fmt.Println(directions)

	return sum
}

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
		// pos := match[1]
		// left := match[2]
		// right := match[3]
		// if pos == "" || left == "" || right == "" {
		// 	fmt.Println(pos)
		// 	fmt.Println(left)
		// 	fmt.Println(right)
		// 	os.Exit(1)
		// }

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
		// fmt.Println(currentPos)

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

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
