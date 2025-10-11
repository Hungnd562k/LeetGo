package problems

func getRow(rowIndex int) []int {
	result := make([][]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		result[i] = make([]int, i+1)
		result[i][0], result[i][i] = 1, 1
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}
	return result[len(result)-1]
}
