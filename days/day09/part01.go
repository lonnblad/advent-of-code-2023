package day09

func SolvePartOneForRealInput() int {
	_, sum := SolvePartOneForInput(realInput)
	return sum
}

func SolvePartOneForInput(input string) ([]int, int) {
	rows := parseRows(input)

	result := make([]int, len(rows))
	sum := 0

	for idx, row := range rows {
		result[idx] = findPrediction(row)
		sum += result[idx]
	}

	return result, sum
}

func findPrediction(row []int) int {
	allZeros, nextRow := findNextRow(row)
	if allZeros {
		return row[len(row)-1]
	}

	newVal := findPrediction(nextRow)
	return row[len(row)-1] + newVal
}
