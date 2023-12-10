package day02

import (
	_ "embed"
	"regexp"
	"strconv"
)

//go:embed real.input
var realInput string

func parseGameNo(game string) int {
	re := regexp.MustCompile(`Game (\d+)`)

	match := re.FindStringSubmatch(game)[1]
	gameNo, _ := strconv.Atoi(match)

	return gameNo
}
