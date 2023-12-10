package day02_test

import (
	_ "embed"
	"testing"

	"github.com/lonnblad/advent-of-code-2023/days/day02"
	"github.com/stretchr/testify/assert"
)

func Test_SolvePartTwoForInput_ExampleInput(t *testing.T) {
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	expectedOutput := []int{48, 12, 1560, 630, 36}
	expectedSum := 2286

	actualOutput, actualSum := day02.SolvePartTwoForInput(input)
	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_SolvePartTwoForInput_RealInput(t *testing.T) {
	expectedSum := 64097
	actualSum := day02.SolvePartTwoForRealInput()
	assert.Equal(t, expectedSum, actualSum)
}
