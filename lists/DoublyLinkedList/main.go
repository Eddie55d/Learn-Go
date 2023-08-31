package main

import (
	"errors"
	"fmt"
)

type Node struct {
	name string
	next *Node
	prev *Node
}

type doublyList struct {
	len  int
	head *Node
}

var (
	ErrListIsEmpty = errors.New("doubly linked list is empty")

	ErrNotFound = errors.New("node not found")
)

func (d *doublyList) Add(name string) {
	if name == "" {
		return
	}

	node := &Node{
		name: name,
	}

	if d.head == nil {
		d.head = node

	} else {
		currentNode := d.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		node.prev = currentNode
		currentNode.next = node
	}
	d.len++
	return
}

func (d *doublyList) Get(str string) (*Node, error) {
	dl := *d

	if dl.head == nil {
		return nil, ErrListIsEmpty
	}

	for dl.head != nil {
		if dl.head.name == str {
			return dl.head, nil
		}
		dl.head = dl.head.next
	}

	return nil, ErrNotFound
}

func (d *doublyList) Delete(str string) (string, error) {

	dl := d

	if dl.head == nil {
		return "", ErrListIsEmpty
	}

	isDeleted := false
	nodeName := ""

	if dl.head.name == str {
		nodeName += dl.head.name
		dl.head.next.prev = nil
		dl.head = dl.head.next
		isDeleted = true
	} else {

		current := dl.head
		for current != nil {
			if current.next == nil {
				break
			}
			if current.next.name == str {

				nodeName += current.next.name
				if current.next.next != nil {
					tail := current.next.next
					current.next = tail
					tail.prev = current

					isDeleted = true
					break
				} else {
					current.next = nil
					isDeleted = true

					break
				}
			}
			current = current.next
		}
	}

	if isDeleted {
		msgDelete := fmt.Sprintf("Node name: %s is deleted", nodeName)
		dl.len--
		return msgDelete, nil
	}
	msgNotFound := fmt.Sprintf("Node name: %s not found", str)
	return msgNotFound, nil

}

func (d *doublyList) IterateList() error {
	dl := *d
	if dl.head == nil {
		return ErrListIsEmpty
	}
	currentNode := dl.head
	for currentNode != nil {
		fmt.Println(currentNode.name)
		currentNode = currentNode.next
	}
	return nil
}

func (s *doublyList) Size() int {
	return s.len
}

func main() {

	list := new(doublyList)

	fmt.Println(list.Get("A"))

	fmt.Println(list.Delete("A"))

	linkedArr := [...]string{"A", "B", "C", "Q", "Z", "D"}

	for _, v := range linkedArr {
		list.Add(v)
	}

	list.IterateList()

	fmt.Printf("Size: %d\n", list.Size())

	fmt.Println(list.Get("A"))
	fmt.Println(list.Get("B"))

	fmt.Println(list.Delete("A"))

	fmt.Printf("Size: %d\n", list.Size())

	fmt.Println(list.Get("D"))
	fmt.Println(list.Get("Z"))
	fmt.Println(list.Get("Q"))

	fmt.Println(list.Delete("Z"))

	fmt.Printf("Size: %d\n", list.Size())

	list.IterateList()

	fmt.Println(list.Get("Q"))

}
