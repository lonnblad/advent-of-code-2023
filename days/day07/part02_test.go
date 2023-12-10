package day07_test

import (
	"testing"

	"github.com/lonnblad/advent-of-code-2023/days/day07"
	"github.com/stretchr/testify/assert"
)

func Test_SolvePartTwoForInput_ExampleInput(t *testing.T) {
	expectedOutput := []int{
		765,
		28,
		684,
		483,
		220,
	}
	expectedSum := 5905

	actualOutput, actualSum := day07.SolvePartTwoForInput(exampleInput)
	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_SolvePartTwoForInput_RealInput(t *testing.T) {
	expectedSum := 251195607
	actualSum := day07.SolvePartTwoForRealInput()
	assert.Equal(t, expectedSum, actualSum)
}
