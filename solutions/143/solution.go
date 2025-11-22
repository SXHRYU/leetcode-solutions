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

func reorderList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

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
		preMiddle.Next = p2.Next
		p2.Next = p1.Next
		p1.Next = p2
		p1 = p2.Next
		p2 = preMiddle.Next
	}
	return head
}

func reorderList1(head *ListNode) *ListNode {
	array := listToArray(head)
	length := len(array)
	if length == 0 {
		return head
	}
	cur := head
	last := head
	for i := 0; i < len(array)/2; i++ {
		tmp := &ListNode{Val: array[length-1-i], Next: cur.Next}
		cur.Next = tmp
		cur = tmp.Next
		last = tmp
	}
	if length%2 == 0 {
		last.Next = nil
	} else {
		cur.Next = nil
	}
	return head
}

func listToArray(head *ListNode) []int {
	array := []int{}
	if head == nil {
		return array
	}
	cur := head
	for cur != nil {
		array = append(array, cur.Val)
		cur = cur.Next
	}
	return array
}
