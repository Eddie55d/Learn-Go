package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	tail, head int
	len        int
	data       *[3]int
}

func (s *Stack) Push(val int) error {
	if s.len == len(s.data) {
		return errors.New("the stack is full")
	}

	if s.len == 0 {
		s.data = &[3]int{}
		s.data[0] = val
		s.len++
		s.head = s.data[0]
		s.tail = s.head
		return nil
	}

	if s.len > 0 {
		data := *s.data
		for i, v := range data {
			if v == 0 {
				s.data[i] = val
				s.tail = s.data[i]
				s.len++
				return nil
			}
		}
	}

	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return 0, errors.New("stack is empty")
	}

	for i, v := range s.data {
		if v == s.tail && i >= 1 {
			popNode := s.data[i]
			s.data[i] = 0
			s.tail = s.data[i-1]
			s.len--
			return popNode, nil
		}
		if v == s.tail && i == 0 {
			popNode := s.data[i]
			s.data[i] = 0
			s.head = 0
			s.tail = 0
			s.len--
			s.data = nil
			return popNode, nil
		}
	}

	return 0, nil
}

func (s *Stack) Size() int {
	return s.len
}

func (s *Stack) Empty() bool {
	if s.head == 0 && s.tail == 0 {
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
