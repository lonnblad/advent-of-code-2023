package day04_test

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lonnblad/advent-of-code-2023/days/day04"
)

func Test_ExampleInput(t *testing.T) {
	expectedOutput := []int{8, 2, 2, 1, 0, 0}
	expectedSum := 13

	actualOutput, actualSum := day04.SolvePartOneForInput(exampleInput)
	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_RealInput(t *testing.T) {
	expectedSum := 26443
	actualSum := day04.SolvePartOneForRealInput()
	assert.Equal(t, expectedSum, actualSum)
}
