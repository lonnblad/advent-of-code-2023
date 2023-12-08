package day07

import (
	"sort"

	"github.com/lonnblad/advent-of-code-2023/cmd/day07/p2"
)

func SolvePartTwoForRealInput() int {
	_, sum := SolvePartTwoForInput(realInput)
	return sum
}

func SolvePartTwoForInput(input string) ([]int, int) {
	hands := p2.ParseHandsForPartOne(input)
	sort.Sort(p2.ByValue(hands))

	result := make([]int, len(hands))
	sum := 0

	for idx, hand := range hands {
		result[idx] = hand.Bid
		sum += hand.Bid * (idx + 1)
	}

	return result, sum
}
