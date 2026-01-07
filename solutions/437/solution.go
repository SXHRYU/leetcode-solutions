package leetcode

import (
	"leetcode-solutions/structures"
)

type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pathSum(root *TreeNode, targetSum int) int {
	prefixSum := make(map[int]int)
	prefixSum[0] = 1
	return dfs(root, prefixSum, 0, targetSum)
}

func dfs(root *TreeNode, prefixSum map[int]int, cur, sum int) int {
	if root == nil {
		return 0
	}
	cur += root.Val
	cnt := 0
	if v, ok := prefixSum[cur-sum]; ok {
		cnt = v
	}
	prefixSum[cur]++
	cnt += dfs(root.Left, prefixSum, cur, sum)
	cnt += dfs(root.Right, prefixSum, cur, sum)
	prefixSum[cur]--
	return cnt
}

func pathSumIII(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := findPath437(root, sum)
	res += pathSumIII(root.Left, sum)
	res += pathSumIII(root.Right, sum)
	return res
}

func findPath437(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Val == sum {
		res++
	}
	res += findPath437(root.Left, sum-root.Val)
	res += findPath437(root.Right, sum-root.Val)
	return res
}
