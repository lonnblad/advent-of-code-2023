package day07

import (
	"sort"

	"github.com/lonnblad/advent-of-code-2023/cmd/day07/p1"
)

func SolvePartOneForRealInput() int {
	_, sum := SolvePartOneForInput(realInput)
	return sum
}

func SolvePartOneForInput(input string) ([]int, int) {
	hands := p1.ParseHandsForPartOne(input)
	sort.Sort(p1.ByValue(hands))

	result := make([]int, len(hands))
	sum := 0

	for idx, hand := range hands {
		result[idx] = hand.Bid * (idx + 1)
		sum += result[idx]
	}

	return result, sum
}
