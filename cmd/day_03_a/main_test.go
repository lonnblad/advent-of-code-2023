package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed example.input
var exampleInput string

func Test_ExampleInput(t *testing.T) {
	expectedOutput := []int{467, 35, 633, 617, 592, 755, 664, 598}
	expectedSum := 4361

	actualOutput, actualSum := calculate(exampleInput)
	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedSum, actualSum)
}

func Test_RealInput(t *testing.T) {
	expectedSum := 512794
	_, actualSum := calculate(realInput)
	assert.Equal(t, expectedSum, actualSum)
}
