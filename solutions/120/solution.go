package leetcode

import (
	"math"
)

func minimumTotal(triangle [][]int) int {
	if triangle == nil {
		return 0
	}
	for row := len(triangle) - 2; row >= 0; row-- {
		for col := 0; col < len(triangle[row]); col++ {
			triangle[row][col] += min(triangle[row+1][col], triangle[row+1][col+1])
		}
	}
	return triangle[0][0]
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func minimumTotal1(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	dp, minNum, index := make([]int, len(triangle[len(triangle)-1])), math.MaxInt64, 0
	for ; index < len(triangle[0]); index++ {
		dp[index] = triangle[0][index]
	}
	for i := 1; i < len(triangle); i++ {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			if j == 0 {

				dp[j] += triangle[i][0]
			} else if j == len(triangle[i])-1 {

				dp[j] += dp[j-1] + triangle[i][j]
			} else {

				dp[j] = min(dp[j-1]+triangle[i][j], dp[j]+triangle[i][j])
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		if dp[i] < minNum {
			minNum = dp[i]
		}
	}
	return minNum
}
