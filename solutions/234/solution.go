package leetcode

import (
	"leetcode-solutions/structures"
)

type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
	slice := []int{}
	for head != nil {
		slice = append(slice, head.Val)
		head = head.Next
	}
	for i, j := 0, len(slice)-1; i < j; {
		if slice[i] != slice[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isPalindrome1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	res := true

	p1 := head
	p2 := head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	preMiddle := p1
	preCurrent := p1.Next
	for preCurrent.Next != nil {
		current := preCurrent.Next
		preCurrent.Next = current.Next
		current.Next = preMiddle.Next
		preMiddle.Next = current
	}

	p1 = head
	p2 = preMiddle.Next

	for p1 != preMiddle {

		if p1.Val == p2.Val {
			p1 = p1.Next
			p2 = p2.Next

		} else {
			res = false
			break
		}
	}
	if p1 == preMiddle {
		if p2 != nil && p1.Val != p2.Val {
			return false
		}
	}
	return res
}

func L2ss(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}
