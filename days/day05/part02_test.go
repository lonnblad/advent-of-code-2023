package day05_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lonnblad/advent-of-code-2023/days/day05"
	"github.com/lonnblad/advent-of-code-2023/days/day05/p2"
)

func Test_SolvePartTwoForInput_ExampleInput(t *testing.T) {
	expectedOutput := []p2.Interval{
		{SourceStart: 82, RangeLength: 3},
		{SourceStart: 46, RangeLength: 11},
		{SourceStart: 86, RangeLength: 4},
		{SourceStart: 94, RangeLength: 3},
		{SourceStart: 56, RangeLength: 4},
		{SourceStart: 97, RangeLength: 2},
	}
	expectedSum := 46

	actualOutput, actualSum := day05.SolvePartTwoForInput(exampleInput)
	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_SolvePartTwoForInput_RealInput(t *testing.T) {
	expectedSum := 51399228
	actualSum := day05.SolvePartTwoForRealInput()
	assert.Equal(t, expectedSum, actualSum)
}
