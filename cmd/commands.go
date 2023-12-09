package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lonnblad/advent-of-code-2023/cmd/day01"
	"github.com/lonnblad/advent-of-code-2023/cmd/day02"
	"github.com/lonnblad/advent-of-code-2023/cmd/day03"
	"github.com/lonnblad/advent-of-code-2023/cmd/day04"
	"github.com/lonnblad/advent-of-code-2023/cmd/day05"
	"github.com/lonnblad/advent-of-code-2023/cmd/day06"
	"github.com/lonnblad/advent-of-code-2023/cmd/day07"
	"github.com/lonnblad/advent-of-code-2023/cmd/day08"
	"github.com/lonnblad/advent-of-code-2023/cmd/day09"
)

func init() {
	rootCmd.Flags().IntVarP(&day, "day", "d", day, "Which Day to run")
	rootCmd.Flags().IntVarP(&part, "part", "p", part, "Which Part to run")
}

var day, part int = 9, 2

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Advent of Code",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Run Day: %d, Part: %d\n", day, part)

		var answer int

		switch {
		case day == 1 && part == 1:
			answer = day01.SolvePartOneForRealInput()
		case day == 1 && part == 2:
			answer = day01.SolvePartTwoForRealInput()
		case day == 2 && part == 1:
			answer = day02.SolvePartOneForRealInput()
		case day == 2 && part == 2:
			answer = day02.SolvePartTwoForRealInput()
		case day == 3 && part == 1:
			answer = day03.SolvePartOneForRealInput()
		case day == 3 && part == 2:
			answer = day03.SolvePartTwoForRealInput()
		case day == 4 && part == 1:
			answer = day04.SolvePartOneForRealInput()
		case day == 4 && part == 2:
			answer = day04.SolvePartTwoForRealInput()
		case day == 5 && part == 1:
			answer = day05.SolvePartOneForRealInput()
		case day == 5 && part == 2:
			answer = day05.SolvePartTwoForRealInput()
		case day == 6 && part == 1:
			answer = day06.SolvePartOneForRealInput()
		case day == 6 && part == 2:
			answer = day06.SolvePartTwoForRealInput()
		case day == 7 && part == 1:
			answer = day07.SolvePartOneForRealInput()
		case day == 7 && part == 2:
			answer = day07.SolvePartTwoForRealInput()
		case day == 8 && part == 1:
			answer = day08.SolvePartOneForRealInput()
		case day == 8 && part == 2:
			answer = day08.SolvePartTwoForRealInput()
		case day == 9 && part == 1:
			answer = day09.SolvePartOneForRealInput()
		case day == 9 && part == 2:
			answer = day09.SolvePartTwoForRealInput()
		default:
			fmt.Println("Not Implemented")
			return
		}

		fmt.Println("Answer:", answer)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
