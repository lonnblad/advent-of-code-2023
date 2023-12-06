package day06_test

import (
	"testing"

	"github.com/lonnblad/advent-of-code-2023/cmd/day06"
	"github.com/stretchr/testify/assert"
)

func Test_SolvePartTwoForInput_ExampleInput(t *testing.T) {
	expectedSum := 71503
	actualSum := day06.SolvePartTwoForInput(exampleInput)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_SolvePartTwoForInput_RealInput(t *testing.T) {
	expectedSum := 41382569
	actualSum := day06.SolvePartTwoForRealInput()
	assert.Equal(t, expectedSum, actualSum)
}
