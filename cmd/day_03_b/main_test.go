package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed example.input
var exampleInput string

func Test_ExampleInput(t *testing.T) {
	expectedOutput := []int{16345, 451490}
	expectedSum := 467835

	actualOutput, actualSum := calculate(exampleInput)
	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_RealInput(t *testing.T) {
	expectedSum := 67779080
	_, actualSum := calculate(realInput)
	assert.Equal(t, expectedSum, actualSum)
}
