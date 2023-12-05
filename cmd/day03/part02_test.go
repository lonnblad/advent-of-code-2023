package day03_test

import (
	_ "embed"
	"testing"

	"github.com/lonnblad/advent-of-code-2023/cmd/day03"
	"github.com/stretchr/testify/assert"
)

//go:embed example.input
var exampleInput string

func Test_SolvePartTwoForInput_ExampleInput(t *testing.T) {
	expectedOutput := []int{16345, 451490}
	expectedSum := 467835

	actualOutput, actualSum := day03.SolvePartTwoForInput(exampleInput)
	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_SolvePartTwoForInput_RealInput(t *testing.T) {
	expectedSum := 67779080
	actualSum := day03.SolvePartTwoForRealInput()
	assert.Equal(t, expectedSum, actualSum)
}
