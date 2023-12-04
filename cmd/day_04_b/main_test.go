package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed example.input
var exampleInput string

func Test_ExampleInput(t *testing.T) {
	expectedOutput := []int{1, 2, 4, 8, 14, 1}
	expectedSum := 30

	actualOutput, actualSum := calculate(exampleInput)
	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_RealInput(t *testing.T) {
	expectedSum := 6284877
	_, actualSum := calculate(realInput)
	assert.Equal(t, expectedSum, actualSum)
}
