package main

import (
	"errors"
	"fmt"
)

type Node struct {
	name int
	next *Node
}

type Stack struct {
	tail, head *Node
	len        int
	data       *[3]*Node
}

func (s *Stack) Push(val int) error {
	if s.len == len(s.data) {
		return errors.New("the stack is full")
	}

	node := &Node{name: val}

	if s.len == 0 {
		s.data = &[3]*Node{}
		s.data[0] = node
		s.len++
		s.head = s.data[0]
		s.tail = s.head
		return nil
	}

	if s.len > 0 {
		data := *s.data
		for i, v := range data {
			if v == nil {
				node.next = data[i-1]
				s.data[i] = node
				s.tail = s.data[i]
				s.len++
				return nil
			}
		}
	}

	return nil
}

func (s *Stack) Pop() (*Node, error) {
	if s.Empty() {
		return nil, errors.New("stack is empty")
	}

	for i, v := range s.data {
		if v == s.tail && i >= 1 {
			popNode := s.data[i]
			s.data[i] = nil
			s.tail = s.data[i-1]
			s.len--
			return popNode, nil
		}
		if v == s.tail && i == 0 {
			popNode := s.data[i]
			s.data[i] = nil
			s.head = nil
			s.tail = nil
			s.len--
			s.data = nil
			return popNode, nil
		}
	}

	return nil, nil
}

func (s *Stack) Size() int {
	return s.len
}

func (s *Stack) Empty() bool {
	if s.head == nil && s.tail == nil {
		return true
	}
	return false
}

func main() {

	s := new(Stack)

	fmt.Printf("Stack is empty: %t\n", s.Empty())
	fmt.Println(s.Push(1))
	fmt.Println(s.data)
	fmt.Println(s.Push(2))
	fmt.Println(s.data)
	fmt.Println(s.Push(3))
	fmt.Println(s.data)
	fmt.Println(s.Push(3))
	fmt.Println(s.data)
	fmt.Println(s.tail)
	fmt.Println(s.head)
	fmt.Printf("Stack is empty: %t\n", s.Empty())
	fmt.Println(s.Size())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.tail)
	fmt.Println(s.head)
	fmt.Println(s.Size())
	fmt.Println(s.Pop())
}
