package main

// https://leetcode.com/problems/merge-two-sorted-lists/

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1.Next == nil && list2.Next == nil {
		return list1
	}

	if list1.Next == nil && list2.Next != nil {
		return list2
	}
	if list1.Next != nil && list2.Next == nil {
		return list1
	}

	curr1 := list1

	curr2 := list2

	mergeListNode := &ListNode{0, nil}

	headNode := mergeListNode

	for {

		if curr1 == nil || curr2 == nil {
			break
		}

		if curr1.Val < curr2.Val {
			headNode.Next = curr1
			curr1 = curr1.Next
		} else {
			headNode.Next = curr2
			curr2 = curr2.Next
		}
		headNode = headNode.Next
	}

	if curr1 != nil {
		headNode.Next = curr1
	} else {
		headNode.Next = curr2
	}

	return mergeListNode.Next

}

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	list := mergeTwoLists(list1, list2)

	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}

}
