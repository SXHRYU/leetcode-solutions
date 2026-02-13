package leetcode

func findTargetSumWays(nums []int, S int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	if S+total < 0 || S > total || (S+total)%2 == 1 {
		return 0
	}
	target := (S + total) / 2
	dp := make([]int, target+1)
	dp[0] = 1
	for _, n := range nums {
		for i := target; i >= n; i-- {
			dp[i] += dp[i-n]
		}
	}
	return dp[target]
}

func findTargetSumWays1(nums []int, S int) int {

	sums := make([]int, len(nums))
	sums[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i > -1; i-- {
		sums[i] = sums[i+1] + nums[i]
	}
	res := 0
	dfsFindTargetSumWays(nums, 0, 0, S, &res, sums)
	return res
}

func dfsFindTargetSumWays(
	nums []int,
	index int,
	curSum int,
	S int,
	res *int,
	sums []int,
) {
	if index == len(nums) {
		if curSum == S {
			*(res) = *(res) + 1
		}
		return
	}

	if S-curSum > sums[index] {
		return
	}
	dfsFindTargetSumWays(nums, index+1, curSum+nums[index], S, res, sums)
	dfsFindTargetSumWays(nums, index+1, curSum-nums[index], S, res, sums)
}
