package day05

import (
	"math"

	"github.com/lonnblad/advent-of-code-2023/days/day05/p2"
)

func SolvePartTwoForRealInput() int {
	_, sum := SolvePartTwoForInput(realInput)
	return sum
}

func SolvePartTwoForInput(input string) ([]p2.Interval, int) {
	seeds := p2.FindSeeds(input)
	soils := p2.MapToSoils(input, seeds)
	fertilizers := p2.MapToFertilizers(input, soils)
	waters := p2.MapToWaters(input, fertilizers)
	lights := p2.MapToLights(input, waters)
	temperatures := p2.MapToTemperatures(input, lights)
	humidities := p2.MapToHumidities(input, temperatures)
	locations := p2.MapToLocations(input, humidities)

	min := math.MaxInt
	for _, loc := range locations {
		if loc.SourceStart < min {
			min = loc.SourceStart
		}
	}

	return locations, min
}
