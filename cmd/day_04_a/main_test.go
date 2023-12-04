package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed example.input
var exampleInput string

func Test_ExampleInput(t *testing.T) {
	expectedOutput := []int{8, 2, 2, 1, 0, 0}
	expectedSum := 13

	actualOutput, actualSum := calculate(exampleInput)
	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_RealInput(t *testing.T) {
	expectedSum := 26443
	_, actualSum := calculate(realInput)
	assert.Equal(t, expectedSum, actualSum)
}
