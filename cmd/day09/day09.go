package day09

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"
)

//go:embed real.input
var realInput string

func parseRows(input string) [][]int {
	rows := strings.Split(input, "\n")
	result := make([][]int, len(rows))

	pattern := regexp.MustCompile(`(\-*\d+)`)

	for idx, row := range rows {
		readings := pattern.FindAllStringSubmatch(row, -1)

		for _, reading := range readings {
			val, _ := strconv.Atoi(reading[0])
			result[idx] = append(result[idx], val)
		}
	}

	return result
}

func findNextRow(row []int) (bool, []int) {
	nextRow := make([]int, len(row)-1)
	allZeros := true

	for idx := 0; idx < len(row)-1; idx++ {
		nextRow[idx] = row[idx+1] - row[idx]
		allZeros = allZeros && nextRow[idx] == 0
	}

	return allZeros, nextRow
}
