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

		solution, exist := solutions[day*10+part]
		if !exist {
			fmt.Println("Not Implemented")
			return
		}

		fmt.Println("Answer:", solution())
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var solutions = map[int]func() int{}

func init() {
	solutions[11] = day01.SolvePartOneForRealInput
	solutions[12] = day01.SolvePartTwoForRealInput
	solutions[21] = day02.SolvePartOneForRealInput
	solutions[22] = day02.SolvePartTwoForRealInput
	solutions[31] = day03.SolvePartOneForRealInput
	solutions[32] = day03.SolvePartTwoForRealInput
	solutions[41] = day04.SolvePartOneForRealInput
	solutions[42] = day04.SolvePartTwoForRealInput
	solutions[51] = day05.SolvePartOneForRealInput
	solutions[52] = day05.SolvePartTwoForRealInput
	solutions[61] = day06.SolvePartOneForRealInput
	solutions[62] = day06.SolvePartTwoForRealInput
	solutions[71] = day07.SolvePartOneForRealInput
	solutions[72] = day07.SolvePartTwoForRealInput
	solutions[81] = day08.SolvePartOneForRealInput
	solutions[82] = day08.SolvePartTwoForRealInput
	solutions[91] = day09.SolvePartOneForRealInput
	solutions[92] = day09.SolvePartTwoForRealInput
}
