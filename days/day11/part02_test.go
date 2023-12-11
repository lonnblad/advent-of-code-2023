package day11_test

import (
	"testing"

	"github.com/lonnblad/advent-of-code-2023/days/day11"
	"github.com/stretchr/testify/assert"
)

func Test_ExampleInput_10(t *testing.T) {
	// expectedResult := []int{18, 28, 68}
	expectedSum := 1030

	_, actualSum := day11.SolvePartTwoForInput(exampleInput, 1e1)
	// assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_ExampleInput_100(t *testing.T) {
	// expectedResult := []int{18, 28, 68}
	expectedSum := 8410

	_, actualSum := day11.SolvePartTwoForInput(exampleInput, 1e2)
	// assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_SolvePartTwoForInput_RealInput(t *testing.T) {
	expectedSum := 597714117556
	actualSum := day11.SolvePartTwoForRealInput()
	assert.Equal(t, expectedSum, actualSum)
}
