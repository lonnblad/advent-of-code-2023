package day01_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lonnblad/advent-of-code-2023/cmd/day01"
)

func Test_SolvePartTwoForInput_Example(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	expectedOutput := []int{29, 83, 13, 24, 42, 14, 76}
	expectedSum := 281

	actualOutput, actualSum := day01.SolvePartTwoForInput(input)
	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_SolvePartTwoForInput_RealInput(t *testing.T) {
	expectedSum := 54719
	actualSum := day01.SolvePartTwoForRealInput()
	assert.Equal(t, expectedSum, actualSum)
}
