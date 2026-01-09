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

func flatten(root *TreeNode) {
	list := preorder(root)
	for i := 1; i < len(list); i++ {
		prev, cur := list[i-1], list[i]
		prev.Left, prev.Right = nil, cur
	}
	return
}

func preorder(root *TreeNode) (ans []*TreeNode) {
	if root != nil {
		ans = append(ans, root)
		ans = append(ans, preorder(root.Left)...)
		ans = append(ans, preorder(root.Right)...)
	}
	return
}

func flatten1(root *TreeNode) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	currRight := root.Right
	root.Right = root.Left
	root.Left = nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = currRight
}

func flatten2(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Right)
	if root.Left == nil {
		return
	}
	flatten(root.Left)
	p := root.Left
	for p.Right != nil {
		p = p.Right
	}
	p.Right = root.Right
	root.Right = root.Left
	root.Left = nil
}
