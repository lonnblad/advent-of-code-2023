package day05

import (
	"math"

	"github.com/lonnblad/advent-of-code-2023/cmd/day05/p1"
)

func SolvePartOneForRealInput() int {
	_, sum := SolvePartOneForInput(realInput)
	return sum
}

func SolvePartOneForInput(input string) ([]int, int) {
	seeds := p1.FindSeeds(input)
	soils := p1.MapToSoils(input, seeds)
	fertilizers := p1.MapToFertilizers(input, soils)
	waters := p1.MapToWaters(input, fertilizers)
	lights := p1.MapToLights(input, waters)
	temperatures := p1.MapToTemperatures(input, lights)
	humidities := p1.MapToHumidities(input, temperatures)
	locations := mapToLocations(input, humidities)

	min := math.MaxInt
	for _, loc := range locations {
		if loc < min {
			min = loc
		}
	}

	return locations, min
}
