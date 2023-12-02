package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExampleInput(t *testing.T) {
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	expectedOutput := []bool{true, true, false, false, true}
	expectedSum := 8

	actualOutput, actualSum := calculate(input)
	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_RealInput(t *testing.T) {
	expectedSum := 2265
	_, actualSum := calculate(realInput)
	assert.Equal(t, expectedSum, actualSum)
}
