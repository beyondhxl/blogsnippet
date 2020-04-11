package main

import (
	"blogsnippet/数据结构与算法/utils"
	"fmt"
)

type MinStack struct {
	data utils.Stack // 数据栈
	help utils.Stack // 辅助栈
}

func (m *MinStack) Push(val int) {
	top, _ := m.data.Top()
	if m.help.IsEmpty() == true || val <= top.(int) {
		m.help.Push(val)
	} else {
		m.help.Push(top.(int))
	}
	m.data.Push(val)
}

func (m *MinStack) Pop() {
	if m.data.IsEmpty() == true || m.help.IsEmpty() == true {
		return
	}
	m.data.Pop()
	m.help.Pop()
}

func (m *MinStack) Top() int {
	if m.data.IsEmpty() {
		return 0
	}
	top, _ := m.data.Top()
	return top.(int)
}

func (m *MinStack) Min() int {
	if m.help.IsEmpty() {
		return 0
	}
	min, _ := m.help.Top()
	return min.(int)
}

func main() {
	m := MinStack{nil, nil}
	m.Push(5)
	m.Push(2)
	m.Push(3)
	fmt.Print("数据栈 --> ", m.data, "\n")
	fmt.Print("辅助栈 --> ", m.help, "\n")
	fmt.Print("栈顶元素 --> ", m.Top(), "\n")
	fmt.Print("栈最小元素 --> ", m.Min(), "\n")
}
