package day05_test

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lonnblad/advent-of-code-2023/days/day05"
)

func Test_ExampleInput(t *testing.T) {
	expectedOutput := []int{82, 43, 86, 35}
	expectedSum := 35

	actualOutput, actualSum := day05.SolvePartOneForInput(exampleInput)
	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_RealInput(t *testing.T) {
	expectedSum := 535088217
	actualSum := day05.SolvePartOneForRealInput()
	assert.Equal(t, expectedSum, actualSum)
}
