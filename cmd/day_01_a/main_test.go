package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExampleInput(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	expectedOutput := []int{12, 38, 15, 77}
	expectedSum := 142

	actualOutput, actualSum := calculate(input)
	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_RealInput(t *testing.T) {
	expectedSum := 55971
	_, actualSum := calculate(realInput)
	assert.Equal(t, expectedSum, actualSum)
}
