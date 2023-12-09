package day08_test

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lonnblad/advent-of-code-2023/cmd/day08"
)

func Test_Example1Input(t *testing.T) {
	expectedSum := 2

	actualSum := day08.SolvePartOneForInput(example1Input)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_Example2Input(t *testing.T) {
	expectedSum := 6

	actualSum := day08.SolvePartOneForInput(example2Input)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_RealInput(t *testing.T) {
	expectedSum := 12643
	actualSum := day08.SolvePartOneForRealInput()
	assert.Equal(t, expectedSum, actualSum)
}
