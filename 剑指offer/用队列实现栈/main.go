package main

import (
	"container/list"
	"fmt"
)

// 实现有问题
type Stack struct {
	que *list.List
	top interface{}
}

func NewStack() *Stack {
	return &Stack{list.New(), 0}
}

func (s *Stack) Push(val interface{}) {
	s.que.PushBack(val)
	s.top = val
}

func (s *Stack) Top() interface{} {
	return s.top
}

func (s *Stack) Pop1() int {
	if s.Empty() == true {
		return -1
	}
	len := s.que.Len()
	for len > 1 {
		e := s.que.Front()
		s.que.Remove(e)
		s.que.PushBack(e)
		len--
	}
	e := s.que.Front()
	s.que.Remove(e)
	return e.Value.(int)
}

func (s *Stack) Len() int {
	return s.que.Len()
}

func (s *Stack) Pop() interface{} {
	if s.Empty() == true {
		return -1
	}

	// len其实是移动次数
	len := s.que.Len()
	for len > 2 {
		e := s.que.Front()
		s.que.Remove(e)
		s.que.PushBack(e)
		len--
	}
	// top更新
	s.top = s.que.Front()

	// 链表头元素是新的队尾
	e := s.que.Front()
	s.que.Remove(e)
	s.que.PushBack(e)

	// 出栈元素要从链表中移除
	e = s.que.Front()
	s.que.Remove(e)

	return e.Value
}

func (s *Stack) Empty() bool {
	return s.que.Len() == 0
}

func main() {
	stack := NewStack()
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	fmt.Println(stack.Len(), stack.Top())
	stack.Pop()
	fmt.Println(stack.Len(), stack.Top())
	stack.Pop()
	fmt.Println(stack.Len(), stack.Top())
	stack.Pop()
	fmt.Println(stack.Len(), stack.Top())
	stack.Pop()
	fmt.Println(stack.Len(), stack.Top())
}
