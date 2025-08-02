package structures

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var NULL = -1 << 63

func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func GetTargetNode(root *TreeNode, target int) *TreeNode {
	if root == nil || root.Val == target {
		return root
	}

	res := GetTargetNode(root.Left, target)
	if res != nil {
		return res
	}
	return GetTargetNode(root.Right, target)
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	msg := fmt.Sprintf("%d 不存在于 %v 中", val, nums)
	panic(msg)
}

func PreIn2Tree(pre, in []int) *TreeNode {
	if len(pre) != len(in) {
		panic("preIn2Tree 中两个切片的长度不相等")
	}

	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: pre[0],
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)

	res.Left = PreIn2Tree(pre[1:idx+1], in[:idx])
	res.Right = PreIn2Tree(pre[idx+1:], in[idx+1:])

	return res
}

func InPost2Tree(in, post []int) *TreeNode {
	if len(post) != len(in) {
		panic("inPost2Tree 中两个切片的长度不相等")
	}

	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: post[len(post)-1],
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)

	res.Left = InPost2Tree(in[:idx], post[:idx])
	res.Right = InPost2Tree(in[idx+1:], post[idx:len(post)-1])

	return res
}

func Tree2Preorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := []int{root.Val}
	res = append(res, Tree2Preorder(root.Left)...)
	res = append(res, Tree2Preorder(root.Right)...)

	return res
}

func Tree2Inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := Tree2Inorder(root.Left)
	res = append(res, root.Val)
	res = append(res, Tree2Inorder(root.Right)...)

	return res
}

func Tree2Postorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := Tree2Postorder(root.Left)
	res = append(res, Tree2Postorder(root.Right)...)
	res = append(res, root.Val)

	return res
}

func (tn *TreeNode) Equal(a *TreeNode) bool {
	if tn == nil && a == nil {
		return true
	}

	if tn == nil || a == nil || tn.Val != a.Val {
		return false
	}

	return tn.Left.Equal(a.Left) && tn.Right.Equal(a.Right)
}

func Tree2ints(tn *TreeNode) []int {
	res := make([]int, 0, 1024)

	queue := []*TreeNode{tn}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			nd := queue[i]
			if nd == nil {
				res = append(res, NULL)
			} else {
				res = append(res, nd.Val)
				queue = append(queue, nd.Left, nd.Right)
			}
		}
		queue = queue[size:]
	}

	i := len(res)
	for i > 0 && res[i-1] == NULL {
		i--
	}

	return res[:i]
}

func T2s(head *TreeNode, array *[]int) {
	fmt.Printf("运行到这里了 head = %v array = %v\n", head, array)

	*array = append(*array, head.Val)
	if head.Left != nil {
		T2s(head.Left, array)
	}
	if head.Right != nil {
		T2s(head.Right, array)
	}
}

func Strings2TreeNode(strs []string) *TreeNode {
	n := len(strs)
	if n == 0 {
		return nil
	}
	x, _ := strconv.Atoi(strs[0])
	root := &TreeNode{Val: x}
	queue := make([]*TreeNode, 1, n<<1)
	queue[0] = root
	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]
		if i < n && strs[i] != "null" {
			x, _ = strconv.Atoi(strs[i])
			node.Left = &TreeNode{Val: x}
			queue = append(queue, node.Left)
		}
		i++
		if i < n && strs[i] != "null" {
			x, _ = strconv.Atoi(strs[i])
			node.Right = &TreeNode{Val: x}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func Tree2LevelOrderStrings(root *TreeNode) []string {
	var ans []string
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	var level int
	for level = 0; len(queue) > 0; level++ {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node == nil {
				ans = append(ans, "null")
			} else {
				ans = append(ans, strconv.Itoa(node.Val))
				if node.Left != nil || node.Right != nil {
					queue = append(queue, node.Left, node.Right)
				}
			}
		}
		queue = queue[size:]
	}
	level--
	return ans
}

func Tree2PreOrderStrings(root *TreeNode) []string {
	var ans []string
	if root == nil {
		return ans
	}
	stack := []*TreeNode{root}
	node := root
	for len(stack) > 0 {
		if node == nil {
			ans = append(ans, "null")
		}
		for node != nil {
			ans = append(ans, strconv.Itoa(node.Val))
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return ans
}
