package day08

import (
	"github.com/lonnblad/advent-of-code-2023/cmd/day08/p1"
)

func SolvePartOneForRealInput() int {
	sum := SolvePartOneForInput(realInput)
	return sum
}

func SolvePartOneForInput(input string) int {
	navs := p1.ParseNavigation(input)
	directions := p1.ParseDirections(input)

	return p1.Navigate(navs, directions)
}
