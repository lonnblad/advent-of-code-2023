package day03

import (
	_ "embed"
	"regexp"
	"strconv"
)

//go:embed real.input
var realInput string

type partNumber struct {
	number          int
	row, start, end int
}

func parsePotentialPartNumbers(rows []string) []partNumber {
	parts := []partNumber{}

	pattern := regexp.MustCompile(`(\d+)`)

	for idx, row := range rows {
		var offset int

		for true {
			loc := pattern.FindIndex([]byte(row))
			if loc == nil {
				break
			}

			val, _ := strconv.Atoi(row[loc[0]:loc[1]])

			part := partNumber{number: val, row: idx, start: offset + loc[0], end: offset + loc[1] - 1}
			parts = append(parts, part)

			offset += loc[1]
			row = row[loc[1]:]
		}
	}

	return parts
}
