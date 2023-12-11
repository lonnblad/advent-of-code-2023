package day11

func SolvePartOneForRealInput() int {
	_, sum := SolvePartOneForInput(realInput)
	return sum
}

func SolvePartOneForInput(input string) ([]int, int) {
	sum := solution(input, 1)
	return nil, sum
}
