package main

import (
	"errors"
	"fmt"
)

type node struct {
	name string
	next *node
}

type singleList struct {
	len  int
	head *node
}

var (
	ErrListIsEmpty = errors.New("singly linked list is empty")

	ErrNotFound = errors.New("node not found")
)

func (s *singleList) Add(name string) {
	node := &node{
		name: name,
	}
	if s.head == nil {
		s.head = node
	} else {
		currentNode := s.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = node
	}
	s.len++
	return
}

func (s *singleList) Get(str string) (string, error) {
	sl := *s

	if s.head == nil {
		return "", ErrListIsEmpty
	}

	for sl.head != nil {
		if sl.head.name == str {
			return sl.head.name, nil
		}
		sl.head = sl.head.next
	}

	return "", ErrNotFound
}

func (s *singleList) Delete(str string) (string, error) {

	sl := s

	if s.head == nil {
		return "", ErrListIsEmpty
	}

	isDeleted := false
	nodeName := ""

	for sl.head != nil {
		if sl.head.next.name == str {

			nodeName += sl.head.next.name
			sl.head.next = sl.head.next.next
			isDeleted = true
			break
		}
		sl.head = sl.head.next
	}

	if isDeleted {
		msgDelete := fmt.Sprintf("Node name: %s is deleted", nodeName)
		s.len--
		return msgDelete, nil
	}
	msgNotFound := fmt.Sprintf("Node name: %s not found", str)
	return msgNotFound, nil

}

func (s *singleList) IterateList() error {
	sl := *s
	if sl.head == nil {
		return ErrListIsEmpty
	}
	currentNode := sl.head
	for currentNode != nil {
		fmt.Println(currentNode.name)
		currentNode = currentNode.next
	}
	return nil
}

func (s *singleList) Size() int {
	return s.len
}

func main() {
	list := new(singleList)

	elA, err := list.Get("A")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(elA)

	str, err := list.Delete("A")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)

	err = list.IterateList()
	if err != nil {
		fmt.Println(err)
	}

	list.Add("A")
	list.Add("B")
	list.Add("C")
	list.Add("Q")
	list.Add("Z")
	list.Add("D")

	err = list.IterateList()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Size: %d\n", list.Size())

	elA, err = list.Get("A")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(elA)

	elB, err := list.Get("B")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(elB)

	str, err = list.Delete("B")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)

	str, err = list.Delete("C")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)

	fmt.Printf("Size: %d\n", list.Size())

	elB, err = list.Get("B")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(elB)

	elC, err := list.Get("C")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(elC)

	elA, err = list.Get("A")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(elA)

	fmt.Printf("Size: %d\n", list.Size())

	err2 := list.IterateList()
	if err2 != nil {
		fmt.Println(err2)
	}

}
