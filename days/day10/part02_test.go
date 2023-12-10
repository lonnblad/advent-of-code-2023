package day10_test

import (
	"testing"

	"github.com/lonnblad/advent-of-code-2023/days/day10"
	"github.com/stretchr/testify/assert"
)

func Test_SolvePartTwoForInput_Example3Input(t *testing.T) {
	// expectedResult := []int{-3, 0, 5}
	expectedSum := 4

	_, actualSum := day10.SolvePartTwoForInput(example3Input)
	// assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_SolvePartTwoForInput_Example4Input(t *testing.T) {
	// expectedResult := []int{-3, 0, 5}
	expectedSum := 4
	// 	example4Input = `..........
	// .S------7.
	// .|F----7|.
	// .||....||.
	// .||....||.
	// .|L-7F-J|.
	// .|..||..|.
	// .L--JL--J.
	// ..........`

	_, actualSum := day10.SolvePartTwoForInput(example4Input)
	// assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_SolvePartTwoForInput_Example5Input(t *testing.T) {
	// expectedResult := []int{-3, 0, 5}
	expectedSum := 8

	_, actualSum := day10.SolvePartTwoForInput(example5Input)
	// assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_SolvePartTwoForInput_Example6Input(t *testing.T) {
	// expectedResult := []int{-3, 0, 5}
	expectedSum := 10

	_, actualSum := day10.SolvePartTwoForInput(example6Input)
	// assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_SolvePartTwoForInput_RealInput(t *testing.T) {
	expectedSum := 297
	actualSum := day10.SolvePartTwoForRealInput()
	assert.Equal(t, expectedSum, actualSum)
}
