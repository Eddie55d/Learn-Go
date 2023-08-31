package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type singleList struct {
	head *ListNode
}

func (s *singleList) Add(num int) {
	node := &ListNode{
		Val: num,
	}
	if s.head == nil {
		s.head = node
	} else {
		currentNode := s.head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = node
	}
	return
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head
	var next *ListNode = nil
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	head = prev
	return head
}

func main() {
	list := new(singleList)

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)

	fmt.Println("Before reverse:")
	func(s singleList) {

		for s.head != nil {
			fmt.Printf("%v\n", s.head.Val)
			s.head = s.head.Next
		}
	}(*list)

	fmt.Println("After reverse:")
	func(s singleList) {
		for s.head != nil {
			node := reverseList(s.head)
			for node != nil {

				fmt.Printf("%v\n", node.Val)
				node = node.Next
			}
			s.head = s.head.Next
		}
	}(*list)
}
