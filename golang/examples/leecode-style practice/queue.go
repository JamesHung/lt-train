package main

import (
	"fmt"
	"errors"
)

type Queue struct {
	data []int
	capacity int
}

func NewQueue(cap int) *Queue {
	return &Queue {
		data: make([]int, 0, cap),
		capacity: cap,
	}
}

func (q *Queue) Enqueue(value int) (err error) {
	if q.Length() >= q.capacity {
		err = errors.New("out of capacity")
		return
	}

	q.data = append(q.data, value)
	return
}

func (q *Queue) Dequeue() (result int, err error) {
	if q.Length() == 0 {
		return 0, errors.New("queue is empty")
	}

	v := q.data[0]
	q.data = q.data[1:]
	return v, nil
}

func (q *Queue) Length() int {
	return len(q.data)
}


func main() {
	queue := NewQueue(100)

	for i:= 0; i < 50; i++ {
		fmt.Println("enqueue ", i)
		queue.Enqueue(i)
		if i%2 == 1 {
			if v, err := queue.Dequeue(); err == nil {
				fmt.Println("dequeue and get value ", v)
			}
		}
	}

	fmt.Println("Size of queue ", queue.Length())
}