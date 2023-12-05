package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lonnblad/advent-of-code-2023/cmd/day01"
	"github.com/lonnblad/advent-of-code-2023/cmd/day02"
)

func init() {
	rootCmd.Flags().IntVarP(&day, "day", "d", day, "Which Day to run")
	rootCmd.Flags().IntVarP(&part, "part", "p", part, "Which Part to run")
}

var day, part int = 2, 2

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
			answer = day02.SolvePartOneForRealInput()
		case day == 3 && part == 2:
			answer = day02.SolvePartTwoForRealInput()
		case day == 4 && part == 1:
			answer = day02.SolvePartOneForRealInput()
		case day == 4 && part == 2:
			answer = day02.SolvePartTwoForRealInput()
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
