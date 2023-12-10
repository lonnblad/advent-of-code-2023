package day10

import (
	"strings"
)

func SolvePartOneForRealInput() int {
	_, sum := SolvePartOneForInput(realInput)
	return sum
}

func SolvePartOneForInput(input string) ([]int, int) {
	startCoord, coords := parseCoords(input)

	result := make([]int, len(coords))

	sum := walkThrough(startCoord, coords)
	sum = int(sum / 2.0)

	return result, sum
}

type coord struct {
	y, x int
}

func parseCoords(input string) (coord, map[coord]string) {
	var startCoord coord
	coords := map[coord]string{}

	rows := strings.Split(input, "\n")
	for y, row := range rows {
		positions := strings.Split(row, "")
		for x, pos := range positions {
			c := coord{y: y, x: x}
			coords[c] = pos

			if pos == "S" {
				startCoord = c
			}
		}
	}

	return startCoord, coords
}

func walkThrough(startCoord coord, coords map[coord]string) int {
	findStartDirection := findStartDirection(startCoord, coords)
	nextCoord, direction := walkFrom(startCoord, findStartDirection, coords)
	steps := 1

	for nextCoord != startCoord {
		nextCoord, direction = walkFrom(nextCoord, direction, coords)
		steps += 1
	}

	return steps
}

type direction string

const (
	north direction = "north"
	east  direction = "east"
	south direction = "south"
	west  direction = "west"
)

func findStartDirection(current coord, coords map[coord]string) direction {
	var coord coord

	// north
	coord.y, coord.x = current.y-1, current.x
	pos := coords[coord]
	switch pos {
	case "|", "7", "F":
		return north
	}

	// east
	coord.y, coord.x = current.y, current.x+1
	pos = coords[coord]
	switch pos {
	case "-", "7", "J":
		return east
	}

	// south
	coord.y, coord.x = current.y+1, current.x
	pos = coords[coord]
	switch pos {
	case "|", "J", "L":
		return south
	}

	// west
	coord.y, coord.x = current.y, current.x-1
	pos = coords[coord]
	switch pos {
	case "-", "F", "L":
		return west
	}

	panic("apa")
}

func walkFrom(current coord, dir direction, coords map[coord]string) (next coord, newDir direction) {
	switch dir {
	case north:
		coord := coord{y: current.y - 1, x: current.x}
		pos := coords[coord]
		switch pos {
		case "|":
			return coord, north
		case "7":
			return coord, west
		case "F":
			return coord, east
		case "S":
			return coord, east
		}
	case east:
		coord := coord{y: current.y, x: current.x + 1}
		pos := coords[coord]
		switch pos {
		case "-":
			return coord, east
		case "7":
			return coord, south
		case "J":
			return coord, north
		case "S":
			return coord, north
		}
	case south:
		coord := coord{y: current.y + 1, x: current.x}
		pos := coords[coord]
		switch pos {
		case "|":
			return coord, south
		case "J":
			return coord, west
		case "L":
			return coord, east
		case "S":
			return coord, east
		}
	case west:
		coord := coord{y: current.y, x: current.x - 1}
		pos := coords[coord]
		switch pos {
		case "-":
			return coord, west
		case "F":
			return coord, south
		case "L":
			return coord, north
		case "S":
			return coord, north
		}
	}

	panic("apa")
}
