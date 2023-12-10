package day09_test

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lonnblad/advent-of-code-2023/days/day09"
)

func Test_Example1Input(t *testing.T) {
	expectedResult := []int{18, 28, 68}
	expectedSum := 114

	actualResult, actualSum := day09.SolvePartOneForInput(exampleInput)
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_RealInput(t *testing.T) {
	expectedSum := 1757008019
	actualSum := day09.SolvePartOneForRealInput()
	assert.Equal(t, expectedSum, actualSum)
}
