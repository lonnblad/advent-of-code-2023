package day11

func SolvePartTwoForRealInput() int {
	_, sum := SolvePartTwoForInput(realInput, 1e6)
	return sum
}

func SolvePartTwoForInput(input string, expand int) ([]int, int) {
	sum := solution(input, expand-1)

	return nil, sum
}
