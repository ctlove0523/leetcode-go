package questions

import "leetcode-go"

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	rows := len(triangle)
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		dp[i][len(triangle[i])-1] = dp[i-1][len(triangle[i])-2] + triangle[i][len(triangle[i])-1]
		for j := 1; j < len(triangle[i])-1; j++ {
			dp[i][j] = leetcode_go.Min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
	}

	return leetcode_go.MinValueOfSlice(dp[rows-1])

}
