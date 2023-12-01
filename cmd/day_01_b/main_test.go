package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Example(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	expectedOutput := []int{29, 83, 13, 24, 42, 14, 76}
	expectedSum := 281

	actualOutput, actualSum := calculate(input)
	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_RealInput(t *testing.T) {
	expectedSum := 54719
	_, actualSum := calculate(realInput)
	assert.Equal(t, expectedSum, actualSum)
}
