package p1

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"
)

func FindSeeds(input string) []int {
	seedInput := strings.SplitN(input, "\n", 2)[0]

	pattern := regexp.MustCompile(`(\d+)`)
	matches := pattern.FindAllStringSubmatch(seedInput[7:], -1)

	result := make([]int, len(matches))
	for idx, match := range matches {
		val, _ := strconv.Atoi(match[0])
		result[idx] = val
	}

	return result
}

func MapToSoils(input string, seeds []int) []int {
	return mapTo(input, "seed-to-soil map:\n", seeds)
}

func MapToFertilizers(input string, seeds []int) []int {
	return mapTo(input, "soil-to-fertilizer map:\n", seeds)
}

func MapToWaters(input string, seeds []int) []int {
	return mapTo(input, "fertilizer-to-water map:\n", seeds)
}

func MapToLights(input string, seeds []int) []int {
	return mapTo(input, "water-to-light map:\n", seeds)
}

func MapToTemperatures(input string, seeds []int) []int {
	return mapTo(input, "light-to-temperature map:\n", seeds)
}

func MapToHumidities(input string, seeds []int) []int {
	return mapTo(input, "temperature-to-humidity map:\n", seeds)
}

func MapToLocations(input string, seeds []int) []int {
	return mapTo(input, "humidity-to-location map:\n", seeds)
}

func mapTo(input, key string, sources []int) []int {
	input = strings.Split(input, key)[1]
	numbers := strings.SplitN(input, "\n\n", 2)[0]

	pattern := regexp.MustCompile(`(\d+) (\d+) (\d+)`)
	rows := strings.Split(numbers, "\n")

	mappers := make([]interval, len(rows))

	for idx, row := range rows {
		matches := pattern.FindStringSubmatch(row)[1:]

		mappers[idx].destStart, _ = strconv.Atoi(matches[0])
		mappers[idx].sourceStart, _ = strconv.Atoi(matches[1])
		mappers[idx].rangeLength, _ = strconv.Atoi(matches[2])
	}

	result := make([]int, len(sources))

	for idx, source := range sources {
		val, has := hasMapping(mappers, source)
		if has {
			result[idx] = val
			continue
		}

		result[idx] = source
	}

	return result
}

type interval struct {
	destStart   int
	sourceStart int
	rangeLength int
}

func hasMapping(mappers []interval, source int) (val int, found bool) {
	for _, interval := range mappers {
		if source >= interval.sourceStart && source < interval.sourceStart+interval.rangeLength {
			return interval.destStart + (source - interval.sourceStart), true
		}
	}

	return 0, false
}
