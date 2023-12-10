package day10_test

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lonnblad/advent-of-code-2023/days/day10"
)

func Test_Example1Input(t *testing.T) {
	// expectedResult := []int{18, 28, 68}
	expectedSum := 4

	_, actualSum := day10.SolvePartOneForInput(example1Input)
	// assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_Example2Input(t *testing.T) {
	// expectedResult := []int{18, 28, 68}
	expectedSum := 8

	_, actualSum := day10.SolvePartOneForInput(example2Input)
	// assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_RealInput(t *testing.T) {
	expectedSum := 6942
	actualSum := day10.SolvePartOneForRealInput()
	assert.Equal(t, expectedSum, actualSum)
}
