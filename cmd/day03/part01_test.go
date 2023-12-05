package day03_test

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lonnblad/advent-of-code-2023/cmd/day03"
)

func Test_ExampleInput(t *testing.T) {
	expectedOutput := []int{467, 35, 633, 617, 592, 755, 664, 598}
	expectedSum := 4361

	actualOutput, actualSum := day03.SolvePartOneForInput(exampleInput)
	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_RealInput(t *testing.T) {
	expectedSum := 512794
	actualSum := day03.SolvePartOneForRealInput()
	assert.Equal(t, expectedSum, actualSum)
}
