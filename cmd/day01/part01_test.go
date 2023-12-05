package day01_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lonnblad/advent-of-code-2023/cmd/day01"
)

func Test_SolvePartOneForInput_ExampleInput(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	expectedOutput := []int{12, 38, 15, 77}
	expectedSum := 142

	actualOutput, actualSum := day01.SolvePartOneForInput(input)
	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_SolvePartOneForInput_RealInput(t *testing.T) {
	expectedSum := 55971
	actualSum := day01.SolvePartOneForRealInput()
	assert.Equal(t, expectedSum, actualSum)
}
