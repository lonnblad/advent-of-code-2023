package day11

import (
	_ "embed"
	"slices"
	"strings"
)

//go:embed real.input
var realInput string

func solution(input string, expansion int) int {
	// === Parse the universe === //
	rows := strings.Split(input, "\n")

	galaxies := [][2]int{}

	rowsWithGalaxies := make(map[int]bool, len(rows))
	columnsWithGalaxies := make(map[int]bool, len(rows[0]))

	for y, row := range rows {
		rowsWithGalaxies[y] = false
		positions := strings.Split(row, "")

		for x, pos := range positions {
			if _, exists := columnsWithGalaxies[x]; !exists {
				columnsWithGalaxies[x] = false
			}

			if pos == "#" {
				galaxies = append(galaxies, [2]int{y, x})
				rowsWithGalaxies[y] = true
				columnsWithGalaxies[x] = true
			}
		}
	}

	// === Expand the universe === //
	newGalaxies := slices.Clone(galaxies)

	for idx, hasGalaxy := range rowsWithGalaxies {
		if !hasGalaxy {
			for jdx, galaxy := range galaxies {
				if idx < galaxy[0] {
					g := newGalaxies[jdx]
					g[0] = g[0] + expansion
					newGalaxies[jdx] = g
				}
			}
		}
	}

	for idx, hasGalaxy := range columnsWithGalaxies {
		if !hasGalaxy {
			for jdx, galaxy := range galaxies {
				if idx < galaxy[1] {
					g := newGalaxies[jdx]

					g[1] = g[1] + expansion
					newGalaxies[jdx] = g
				}
			}
		}
	}

	// === Calculate and Sum Distances === //
	var sum int

	for n := 0; n < len(newGalaxies)-1; n++ {
		for k := n + 1; k < len(newGalaxies); k++ {
			dy := newGalaxies[n][0] - newGalaxies[k][0]
			dx := newGalaxies[n][1] - newGalaxies[k][1]

			if dy < 0 {
				dy *= -1
			}
			if dx < 0 {
				dx *= -1
			}

			sum += dy + dx
		}
	}

	return sum
}
