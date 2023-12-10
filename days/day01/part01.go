package day01

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"
)

func SolvePartOneForRealInput() int {
	_, sum := SolvePartOneForInput(realInput)
	return sum
}

func SolvePartOneForInput(input string) ([]int, int) {
	parts := strings.Split(input, "\n")
	result := make([]int, len(parts))
	sum := 0

	for idx, part := range parts {
		re := regexp.MustCompile(`^[^\d]*(\d{0,1}).*(\d)[^\d]*$`)
		match := re.FindStringSubmatch(part)

		val := match[1] + match[2]
		if match[1] == "" {
			val += match[2]
		}

		number, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		result[idx] = number
		sum += number
	}

	return result, sum
}
