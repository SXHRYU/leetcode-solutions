package leetcode

func rob198(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0], dp[1] = nums[0], max(nums[1], nums[0])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[n-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob198_1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	curMax, preMax := 0, 0
	for i := 0; i < n; i++ {
		tmp := curMax
		curMax = max(curMax, nums[i]+preMax)
		preMax = tmp
	}
	return curMax
}

func rob(nums []int) int {

	a, b := 0, 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			a = max(a+nums[i], b)
		} else {
			b = max(a, b+nums[i])
		}
	}
	return max(a, b)
}
