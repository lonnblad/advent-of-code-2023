package day08_test

import (
	"testing"

	"github.com/lonnblad/advent-of-code-2023/days/day08"
	"github.com/stretchr/testify/assert"
)

func Test_SolvePartTwoForInput_Example3Input(t *testing.T) {
	expectedSum := 6

	actualSum := day08.SolvePartTwoForInput(example3Input)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_SolvePartTwoForInput_RealInput(t *testing.T) {
	expectedSum := 13133452426987
	actualSum := day08.SolvePartTwoForRealInput()
	assert.Equal(t, expectedSum, actualSum)
}
