package main

import (
	"errors"
	"fmt"
)


// =====Stack=======
type Stack struct {
	data     []int
	capacity int
}

func NewStack(cap int) *Stack {
	return &Stack{
		data:     make([]int, 0, cap),
		capacity: cap,
	}
}

func (s *Stack) Push(value int) error {
	if len(s.data) >= s.capacity {
		return errors.New("Stack is full")
	}

	s.data = append(s.data, value)
	return nil
}

func (s *Stack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("no data")
	}

	n := len(s.data)-1
	result := s.data[n]
	s.data = s.data[:n]
	return result, nil
}

func (s *Stack) Peek() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("no data")
	}
	return s.data[len(s.data)-1], nil
}

func (s *Stack) Length() int {
	return len(s.data)
}

func testStack() {

	stack := NewStack(10)

	for i := 0; i < 10; i++ {
		stack.Push(i)
	}

	for i := 0; i < 5; i++ {
		v, err := stack.Pop()
		if err != nil {
			fmt.Println("pop err: ", err)
			continue
		}

		fmt.Println("pop got ", v)
	}

	if v, err := stack.Peek(); err == nil {
		fmt.Println("peek stack ", v)
	}
	
	fmt.Println("size of stack ", stack.Length())

}


func main() {
	testStack()
}
