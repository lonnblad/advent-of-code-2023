package p2

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"
)

func FindSeeds(input string) []Interval {
	seedInput := strings.SplitN(input, "\n", 2)[0]

	pattern := regexp.MustCompile(`(\d+) (\d+)`)
	matches := pattern.FindAllStringSubmatch(seedInput[7:], -1)

	result := make([]Interval, len(matches))
	for idx, match := range matches {
		result[idx].SourceStart, _ = strconv.Atoi(match[1])
		result[idx].RangeLength, _ = strconv.Atoi(match[2])
	}

	return result
}

func MapToSoils(input string, source []Interval) []Interval {
	return mapTo(input, "seed-to-soil map:\n", source)
}

func MapToFertilizers(input string, source []Interval) []Interval {
	return mapTo(input, "soil-to-fertilizer map:\n", source)
}

func MapToWaters(input string, source []Interval) []Interval {
	return mapTo(input, "fertilizer-to-water map:\n", source)
}

func MapToLights(input string, source []Interval) []Interval {
	return mapTo(input, "water-to-light map:\n", source)
}

func MapToTemperatures(input string, source []Interval) []Interval {
	return mapTo(input, "light-to-temperature map:\n", source)
}

func MapToHumidities(input string, source []Interval) []Interval {
	return mapTo(input, "temperature-to-humidity map:\n", source)
}

func MapToLocations(input string, source []Interval) []Interval {
	return mapTo(input, "humidity-to-location map:\n", source)
}

func mapTo(input, key string, sourceIntervals []Interval) []Interval {
	input = strings.Split(input, key)[1]
	numbers := strings.SplitN(input, "\n\n", 2)[0]

	pattern := regexp.MustCompile(`(\d+) (\d+) (\d+)`)
	rows := strings.Split(numbers, "\n")

	mappers := make([]mapInterval, len(rows))

	for idx, row := range rows {
		matches := pattern.FindStringSubmatch(row)[1:]

		mappers[idx].destStart, _ = strconv.Atoi(matches[0])
		mappers[idx].SourceStart, _ = strconv.Atoi(matches[1])
		mappers[idx].RangeLength, _ = strconv.Atoi(matches[2])
	}

	result := []Interval{}

	for _, source := range sourceIntervals {
		result = append(result, tryAndMap(mappers, source)...)
	}

	return result
}

type Interval struct {
	SourceStart int
	RangeLength int
}

type mapInterval struct {
	Interval
	destStart int
}

func tryAndMap(mappers []mapInterval, source Interval) (_ []Interval) {
	for _, mapper := range mappers {
		if source.SourceStart >= mapper.SourceStart && source.SourceStart+source.RangeLength <= mapper.SourceStart+mapper.RangeLength {
			// The whole interval fits in the map

			source.SourceStart = mapper.destStart - mapper.SourceStart + source.SourceStart
			return []Interval{source}
		}

		if source.SourceStart >= mapper.SourceStart && source.SourceStart < mapper.SourceStart+mapper.RangeLength {
			// Part of the interval fits in the map

			newRangeLength := mapper.SourceStart + mapper.RangeLength - source.SourceStart
			intervalOne := Interval{SourceStart: source.SourceStart, RangeLength: newRangeLength}
			newIntervalsOne := tryAndMap(mappers, intervalOne)

			intervalTwo := Interval{SourceStart: source.SourceStart + newRangeLength, RangeLength: source.RangeLength - newRangeLength}
			newIntervalsTwo := tryAndMap(mappers, intervalTwo)

			return append(newIntervalsOne, newIntervalsTwo...)
		}
	}

	return []Interval{source}
}
