package day10

func SolvePartTwoForRealInput() int {
	_, sum := SolvePartTwoForInput(realInput)
	return sum
}

func SolvePartTwoForInput(input string) ([]int, int) {
	maze := pipeMazeP2{}
	maze.start, maze.coords = parseCoords(input)

	// result := make([]int, len(coords))

	sum := maze.walkThrough()

	return nil, sum
}

type pipeMazeP2 struct {
	start     coord
	coords    map[coord]string
	traversed map[coord]bool
	left      map[coord]bool
	right     map[coord]bool
}

func (maze pipeMazeP2) walkThrough() int {
	var (
		leftUnique  = map[coord]bool{}
		rightUnique = map[coord]bool{}
		traversed   = map[coord]bool{}
	)

	traversed[maze.start] = true

	findStartDirection := maze.findStartDirection()
	nextCoord, direction, left, right := p2WalkFrom(maze.start, findStartDirection, maze.coords)
	steps := 1

	for _, c := range left {
		leftUnique[c] = true
	}
	for _, c := range right {
		rightUnique[c] = true
	}

	traversed[nextCoord] = true

	for nextCoord != maze.start {
		nextCoord, direction, left, right = p2WalkFrom(nextCoord, direction, maze.coords)
		steps += 1

		for _, c := range left {
			leftUnique[c] = true
		}
		for _, c := range right {
			rightUnique[c] = true
		}

		traversed[nextCoord] = true
	}

	var (
		prevGrouped int = -1
		grouped     int
	)

	for prevGrouped != grouped {
		prevGrouped = grouped
		grouped = 0

		for c := range maze.coords {
			if traversed[c] || leftUnique[c] || rightUnique[c] {
				continue
			}

			allNeighbours := []coord{
				{y: c.y - 1, x: c.x - 1},
				{y: c.y - 1, x: c.x},
				{y: c.y - 1, x: c.x + 1},
				{y: c.y, x: c.x - 1},
				{y: c.y, x: c.x + 1},
				{y: c.y + 1, x: c.x - 1},
				{y: c.y + 1, x: c.x},
				{y: c.y + 1, x: c.x + 1},
			}

			for _, n := range allNeighbours {
				if traversed[n] {
					continue
				}

				if leftUnique[n] {
					leftUnique[c] = true
					grouped++
					break
				}

				if rightUnique[n] {
					rightUnique[c] = true
					grouped++
					break
				}
			}
		}
	}

	var (
		height int
		with   int
	)

	for coord := range maze.coords {
		if coord.y > height {
			height = coord.y
		}
		if coord.x > with {
			with = coord.x
		}
	}

	var leftIsOuter bool

	for coord := range leftUnique {
		if traversed[coord] {
			delete(leftUnique, coord)
		}

		if !traversed[coord] {
			if coord.x == 0 || coord.x == with || coord.y == 0 || coord.y == height {
				leftIsOuter = true
				break
			}
		}
	}

	if !leftIsOuter {
		return len(leftUnique)
	}

	var rightIsOuter bool

	for coord := range rightUnique {
		if traversed[coord] {
			delete(rightUnique, coord)
		}

		if !traversed[coord] {
			if coord.x == 0 || coord.x == with || coord.y == 0 || coord.y == height {
				rightIsOuter = true
				break
			}
		}
	}

	if !rightIsOuter {
		return len(rightUnique)
	}

	panic("p2WalkThrough")
}

func (maze pipeMazeP2) findStartDirection() direction {
	var coord coord

	// north
	coord.y, coord.x = maze.start.y-1, maze.start.x
	pos := maze.coords[coord]
	switch pos {
	case "|", "7", "F":
		return north
	}

	// east
	coord.y, coord.x = maze.start.y, maze.start.x+1
	pos = maze.coords[coord]
	switch pos {
	case "-", "7", "J":
		return east
	}

	// south
	coord.y, coord.x = maze.start.y+1, maze.start.x
	pos = maze.coords[coord]
	switch pos {
	case "|", "J", "L":
		return south
	}

	// west
	coord.y, coord.x = maze.start.y, maze.start.x-1
	pos = maze.coords[coord]
	switch pos {
	case "-", "F", "L":
		return west
	}

	panic("findStartDirection")
}

func p2WalkFrom(current coord, dir direction, coords map[coord]string) (next coord, newDir direction, left, right []coord) {
	switch dir {
	case north:
		nextCoord := coord{y: current.y - 1, x: current.x}
		pos := coords[nextCoord]
		switch pos {
		case "|":
			left := []coord{{y: nextCoord.y, x: nextCoord.x - 1}}
			right := []coord{{y: nextCoord.y, x: nextCoord.x + 1}}

			return nextCoord, north, left, right
		case "7":
			left := []coord{}
			right := []coord{
				{y: nextCoord.y, x: nextCoord.x + 1},
				{y: nextCoord.y - 1, x: nextCoord.x + 1},
				{y: nextCoord.y - 1, x: nextCoord.x},
			}

			return nextCoord, west, left, right
		case "F":
			left := []coord{
				{y: nextCoord.y, x: nextCoord.x - 1},
				{y: nextCoord.y - 1, x: nextCoord.x - 1},
				{y: nextCoord.y - 1, x: nextCoord.x},
			}
			right := []coord{}

			return nextCoord, east, left, right
		case "S":
			return nextCoord, east, []coord{}, []coord{}
		}
	case east:
		nextCoord := coord{y: current.y, x: current.x + 1}
		pos := coords[nextCoord]
		switch pos {
		case "-":
			left := []coord{{y: nextCoord.y - 1, x: nextCoord.x}}
			right := []coord{{y: nextCoord.y + 1, x: nextCoord.x}}

			return nextCoord, east, left, right
		case "7":
			left := []coord{
				{y: nextCoord.y - 1, x: nextCoord.x},
				{y: nextCoord.y - 1, x: nextCoord.x + 1},
				{y: nextCoord.y, x: nextCoord.x + 1},
			}
			right := []coord{}

			return nextCoord, south, left, right
		case "J":
			left := []coord{}
			right := []coord{
				{y: nextCoord.y + 1, x: nextCoord.x},
				{y: nextCoord.y + 1, x: nextCoord.x + 1},
				{y: nextCoord.y, x: nextCoord.x + 1},
			}

			return nextCoord, north, left, right
		case "S":
			return nextCoord, north, []coord{}, []coord{}
		}
	case south:
		nextCoord := coord{y: current.y + 1, x: current.x}
		pos := coords[nextCoord]
		switch pos {
		case "|":
			left := []coord{{y: nextCoord.y, x: nextCoord.x + 1}}
			right := []coord{{y: nextCoord.y, x: nextCoord.x - 1}}

			return nextCoord, south, left, right
		case "J":
			left := []coord{
				{y: nextCoord.y, x: nextCoord.x + 1},
				{y: nextCoord.y + 1, x: nextCoord.x + 1},
				{y: nextCoord.y + 1, x: nextCoord.x},
			}
			right := []coord{}

			return nextCoord, west, left, right
		case "L":
			left := []coord{}
			right := []coord{
				{y: nextCoord.y, x: nextCoord.x - 1},
				{y: nextCoord.y + 1, x: nextCoord.x - 1},
				{y: nextCoord.y + 1, x: nextCoord.x},
			}

			return nextCoord, east, left, right
		case "S":
			return nextCoord, east, []coord{}, []coord{}
		}
	case west:
		nextCoord := coord{y: current.y, x: current.x - 1}
		pos := coords[nextCoord]
		switch pos {
		case "-":
			left := []coord{{y: nextCoord.y + 1, x: nextCoord.x}}
			right := []coord{{y: nextCoord.y - 1, x: nextCoord.x}}

			return nextCoord, west, left, right
		case "F":
			left := []coord{}
			right := []coord{
				{y: nextCoord.y - 1, x: nextCoord.x},
				{y: nextCoord.y - 1, x: nextCoord.x - 1},
				{y: nextCoord.y, x: nextCoord.x - 1},
			}

			return nextCoord, south, left, right
		case "L":
			left := []coord{
				{y: nextCoord.y + 1, x: nextCoord.x},
				{y: nextCoord.y + 1, x: nextCoord.x - 1},
				{y: nextCoord.y, x: nextCoord.x - 1},
			}
			right := []coord{}

			return nextCoord, north, left, right
		case "S":
			return nextCoord, north, []coord{}, []coord{}
		}
	}

	panic("p2WalkFrom")
}
