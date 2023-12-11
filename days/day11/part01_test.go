package day11_test

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lonnblad/advent-of-code-2023/days/day11"
)

func Test_ExampleInput(t *testing.T) {
	// expectedResult := []int{18, 28, 68}
	expectedSum := 374

	_, actualSum := day11.SolvePartOneForInput(exampleInput)
	// assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_RealInput(t *testing.T) {
	expectedSum := 9312968
	actualSum := day11.SolvePartOneForRealInput()
	assert.Equal(t, expectedSum, actualSum)
}
