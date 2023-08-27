package main

import (
	"errors"
	"fmt"
)

type Node struct {
	name int
	next *Node
}

type Queue struct {
	head, tail *Node
	len        int
	data       *[3]*Node
}

func (q *Queue) Push(val int) error {
	if q.len == len(q.data) {
		return errors.New("the queue is full")
	}

	node := &Node{name: val}

	if q.len == 0 {
		q.data = &[3]*Node{}
		q.data[0] = node
		q.len++
		q.head = q.data[0]
		return nil
	}

	if q.len > 0 {
		data := *q.data
		h := 0
		for i, v := range data {
			if i != len(data)-1 {
				q.data[i+1] = v
				h = i + 1
			}
		}
		q.data[0] = node
		q.tail = q.data[0]
		q.data[0].next = nil
		q.head = q.data[h]
		q.len++

		newData := *q.data
		for i, v := range newData {
			if v != nil && i != 0 {
				q.data[i].next = newData[i-1]
			}
		}
	}

	return nil
}

func (q *Queue) Pop() (*Node, error) {
	if q.len == 0 {
		return nil, errors.New("queue is empty")
	}

	for i, v := range q.data {
		if v == q.head && i >= 1 {
			popNode := q.data[i]
			q.data[i] = nil
			q.head = q.data[i-1]
			q.len--
			return popNode, nil
		}
		if v == q.head && i == 0 {
			popNode := q.data[i]
			q.data[i] = nil
			q.head = nil
			q.tail = nil
			q.len--
			q.data = nil
			return popNode, nil
		}
	}

	return nil, nil
}

func (q *Queue) Size() int {
	return q.len
}

func (q *Queue) Empty() bool {
	if q.head == nil && q.tail == nil {
		return true
	}
	return false
}

func main() {

	q := new(Queue)

	fmt.Printf("Queue is empty: %t\n", q.Empty())
	fmt.Println(q.Push(1))
	fmt.Println(q.data)
	fmt.Println(q.Push(2))
	fmt.Println(q.data)
	fmt.Println(q.Push(3))
	fmt.Println(q.data)
	fmt.Println(q.Push(3))
	fmt.Println(q.data)
	fmt.Println(q.tail)
	fmt.Println(q.head)
	fmt.Printf("Queue is empty: %t\n", q.Empty())
	fmt.Println(q.Size())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.tail)
	fmt.Println(q.head)
	fmt.Println(q.Size())
	fmt.Println(q.Pop())
}
