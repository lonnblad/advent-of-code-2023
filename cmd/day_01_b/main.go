package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed real.input
var realInput string

func main() {
	_, sum := calculate(realInput)
	fmt.Println("Sum:", sum)
}

func calculate(input string) ([]int, int) {
	parts := strings.Split(input, "\n")
	result := make([]int, len(parts))
	sum := 0

	for idx, part := range parts {
		first := parseFirst(part)
		second := parseSecond(part)

		number := first + second
		result[idx] = number
		sum += number
	}

	return result, sum
}

func parseFirst(str string) int {
	re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|\d)`)
	matches := re.FindStringSubmatch(str)
	return mapper[matches[0]] * 10
}

func parseSecond(str string) int {
	for idx := len(str) - 1; idx >= 0; idx-- {
		re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|\d)`)
		matches := re.FindStringSubmatch(str[idx:])
		if len(matches) > 0 {
			return mapper[matches[0]]
		}
	}

	panic("couldn't find second value")
}

var mapper = map[string]int{
	"one": 1, "1": 1,
	"two": 2, "2": 2,
	"three": 3, "3": 3,
	"four": 4, "4": 4,
	"five": 5, "5": 5,
	"six": 6, "6": 6,
	"seven": 7, "7": 7,
	"eight": 8, "8": 8,
	"nine": 9, "9": 9,
}
