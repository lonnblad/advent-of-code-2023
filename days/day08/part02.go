package day08

import "github.com/lonnblad/advent-of-code-2023/days/day08/p2"

func SolvePartTwoForRealInput() int {
	sum := SolvePartTwoForInput(realInput)
	return sum
}

func SolvePartTwoForInput(input string) int {
	navs := p2.ParseNavigation(input)
	startPositions, directions := p2.ParseDirections(input)

	result := make([]int, len(startPositions))
	var sum int
	for idx, startPos := range startPositions {
		result[idx] = p2.Navigate(navs, startPos, directions)
	}

	sum = p2.LCM(result[0], result[1], result[2:]...)

	return sum
}
