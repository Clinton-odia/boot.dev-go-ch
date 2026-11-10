package main

func createMatrix(rows, cols int) [][]int {
	result := [][]int{}
	for i := range rows {
		var row []int
		for j := range cols {
			row = append(row, (i * j))
		}
		result = append(result, row)
	}
	return result
}
