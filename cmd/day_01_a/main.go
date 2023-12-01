package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
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
