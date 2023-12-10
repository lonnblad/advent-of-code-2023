package day04_test

import (
	"testing"

	"github.com/lonnblad/advent-of-code-2023/days/day04"
	"github.com/stretchr/testify/assert"
)

func Test_SolvePartTwoForInput_ExampleInput(t *testing.T) {
	expectedOutput := []int{1, 2, 4, 8, 14, 1}
	expectedSum := 30

	actualOutput, actualSum := day04.SolvePartTwoForInput(exampleInput)
	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_SolvePartTwoForInput_RealInput(t *testing.T) {
	expectedSum := 6284877
	actualSum := day04.SolvePartTwoForRealInput()
	assert.Equal(t, expectedSum, actualSum)
}
