package day06

import (
	_ "embed"
)

//go:embed real.input
var realInput string

type race struct {
	time     int
	distance int
}

func (race race) calcWaysOfWinning() int {
	var win int

	for i := 0; i <= race.time; i++ {
		length := i * (race.time - i)

		if length > race.distance {
			win++
		}
	}

	return win
}
