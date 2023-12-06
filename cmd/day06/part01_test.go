package day06_test

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lonnblad/advent-of-code-2023/cmd/day06"
)

func Test_ExampleInput(t *testing.T) {
	expectedOutput := []int{4, 8, 9}
	expectedSum := 288

	actualOutput, actualSum := day06.SolvePartOneForInput(exampleInput)
	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_RealInput(t *testing.T) {
	expectedSum := 840336
	actualSum := day06.SolvePartOneForRealInput()
	assert.Equal(t, expectedSum, actualSum)
}
