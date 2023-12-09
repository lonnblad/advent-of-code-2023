package day09_test

import (
	"testing"

	"github.com/lonnblad/advent-of-code-2023/cmd/day09"
	"github.com/stretchr/testify/assert"
)

func Test_SolvePartTwoForInput_Example3Input(t *testing.T) {
	expectedResult := []int{-3, 0, 5}
	expectedSum := 2

	actualResult, actualSum := day09.SolvePartTwoForInput(exampleInput)
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_SolvePartTwoForInput_RealInput(t *testing.T) {
	expectedSum := 995
	actualSum := day09.SolvePartTwoForRealInput()
	assert.Equal(t, expectedSum, actualSum)
}
