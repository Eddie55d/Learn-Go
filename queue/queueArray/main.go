package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	head, tail, len int
	data            *[3]int
}

func (q *Queue) Push(val int) error {
	if q.len == len(q.data) {
		return errors.New("the queue is full")
	}
	if q.len == 0 {
		q.data = &[3]int{}
		q.data[0] = val
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
		q.data[0] = val
		q.tail = q.data[0]
		q.head = q.data[h]
		q.len++
	}

	return nil
}

func (q *Queue) Pop() (int, error) {
	if q.Empty() {
		return 0, errors.New("queue is empty")
	}

	for i, v := range q.data {
		if v == q.head && i >= 1 {
			val := q.data[i]
			q.data[i] = 0
			q.head = q.data[i-1]
			q.len--
			return val, nil
		}
		if v == q.head && i == 0 {
			val := q.data[i]
			q.data[i] = 0
			q.head = 0
			q.tail = 0
			q.len--
			q.data = nil
			return val, nil
		}
	}

	return 0, nil
}

func (q *Queue) Size() int {
	return q.len
}

func (q *Queue) Empty() bool {
	if q.head == 0 && q.tail == 0 {
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
