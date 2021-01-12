package main

import (
	"container/list"
	"fmt"
)

//import (
//	"container/list"
//	"fmt"
//)
//
//type Stack struct {
//	que *list.List
//	top interface{}
//}
//
//func NewStack() *Stack {
//	return &Stack{list.New(), 0}
//}
//
//func (s *Stack) Push(val interface{}) {
//	s.que.PushBack(val)
//	s.top = val
//}
//
//func (s *Stack) Top() interface{} {
//	return s.top
//}
//
////栈先进后出 pop 弹出的是最后进入队列的
//func (s *Stack) Pop1() int {
//	if s.Empty() == true {
//		return -1
//	}
//
//	e := s.que.Back()
//	s.que.Remove(e)
//
//	s.top = e
//
//	return e.Value.(int)
//
//	//len := s.que.Len()
//	//for len > 1 {
//	//	e := s.que.Front()
//	//	s.que.Remove(e)
//	//	s.que.PushBack(e)
//	//	len--
//	//}
//	//e := s.que.Front()
//	//s.que.Remove(e)
//	//return e.Value.(int)
//}
//
//func (s *Stack) Len() int {
//	return s.que.Len()
//}
//
//func (s *Stack) Pop() interface{} {
//	if s.Empty() == true {
//		return -1
//	}
//
//	// len其实是移动次数
//	len := s.que.Len()
//	for len > 2 {
//		e := s.que.Front()
//		s.que.Remove(e)
//		s.que.PushBack(e)
//		len--
//	}
//	// top更新
//	s.top = s.que.Front()
//
//	// 链表头元素是新的队尾
//	e := s.que.Front()
//	s.que.Remove(e)
//	s.que.PushBack(e)
//
//	// 出栈元素要从链表中移除
//	e = s.que.Front()
//	s.que.Remove(e)
//
//	return e.Value
//}
//
//func (s *Stack) Empty() bool {
//	return s.que.Len() == 0
//}
//
//func main() {
//	stack := NewStack()
//	stack.Push(3)
//	stack.Push(4)
//	stack.Push(5)
//	stack.Push(6)
//	fmt.Println(stack.Len(), stack.Top())
//	stack.Pop1()
//	fmt.Println(stack.Len(), stack.Top())
//	stack.Pop1()
//	fmt.Println(stack.Len(), stack.Top())
//	stack.Pop1()
//	fmt.Println(stack.Len(), stack.Top())
//	stack.Pop1()
//	fmt.Println(stack.Len(), stack.Top())
//}

type MyStack struct {
	in  *list.List
	out *list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		in:  list.New(),
		out: list.New(),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.in.PushBack(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	for this.in.Front().Next() != nil {
		h := this.in.Front().Value
		this.in.Remove(this.in.Front())
		this.out.PushBack(h)
	}

	top := this.in.Front().Value
	this.in.Remove(this.in.Front())

	//又转储一下
	for this.out.Front() != nil {
		h := this.out.Front().Value
		this.out.Remove(this.out.Front())
		this.in.PushBack(h)
	}

	v, _ := top.(int)
	return v
}

/** Get the top element. */
func (this *MyStack) Top() int {
	for this.in.Front().Next() != nil {
		h := this.in.Front().Value
		this.in.Remove(this.in.Front())
		this.out.PushBack(h)
	}

	top := this.in.Front().Value
	this.in.Init()

	for this.out.Front() != nil {
		h := this.out.Front().Value
		this.out.Remove(this.out.Front())
		this.in.PushBack(h)
	}
	this.out.Init()
	this.in.PushBack(top)

	v, _ := top.(int)
	return v
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if this.in.Front() == nil {
		return true
	} else {
		return false
	}
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	param_2 := obj.Pop()
	param_3 := obj.Pop()
	param_4 := obj.Pop()
	param_5 := obj.Empty()

	fmt.Println(param_2, param_3, param_4, param_5)
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
