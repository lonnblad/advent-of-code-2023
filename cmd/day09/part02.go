package day09

func SolvePartTwoForRealInput() int {
	_, sum := SolvePartTwoForInput(realInput)
	return sum
}

func SolvePartTwoForInput(input string) ([]int, int) {
	rows := parseRows(input)

	result := make([]int, len(rows))
	sum := 0

	for idx, row := range rows {
		result[idx] = findPrevHistoricReading(row)
		sum += result[idx]
	}

	return result, sum
}

func findPrevHistoricReading(row []int) int {
	allZeros, nextRow := findNextRow(row)
	if allZeros {
		return row[0]
	}

	newVal := findPrevHistoricReading(nextRow)
	return row[0] - newVal
}
