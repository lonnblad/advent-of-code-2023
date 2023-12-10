package day07_test

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lonnblad/advent-of-code-2023/days/day07"
)

func Test_ExampleInput(t *testing.T) {
	expectedOutput := []int{
		765 * 1,
		220 * 2,
		28 * 3,
		684 * 4,
		483 * 5,
	}
	expectedSum := 6440

	actualOutput, actualSum := day07.SolvePartOneForInput(exampleInput)
	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_RealInput(t *testing.T) {
	expectedSum := 248836197
	actualSum := day07.SolvePartOneForRealInput()
	assert.Equal(t, expectedSum, actualSum)
}
