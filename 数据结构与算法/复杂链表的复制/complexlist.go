package main

import (
	"fmt"
)

type ComplexListNode struct {
	Val     int              // 值
	Next    *ComplexListNode // 下一个节点
	Sibling *ComplexListNode // 随机的兄弟节点
}

func CloneComplexList(head *ComplexListNode) *ComplexListNode {
	if head == nil {
		return nil
	}
	m := make(map[*ComplexListNode]*ComplexListNode)
	cur := head
	for cur != nil {
		if _, ok := m[cur]; !ok {
			m[cur] = &ComplexListNode{cur.Val, nil, nil}
		}
		if cur.Next != nil {
			if _, ok := m[cur.Next]; !ok {
				m[cur.Next] = &ComplexListNode{cur.Next.Val, nil, nil}
			}
			m[cur].Next = m[cur.Next]
		}

		if cur.Sibling != nil {
			if _, ok := m[cur.Sibling]; !ok {
				m[cur.Sibling] = &ComplexListNode{cur.Sibling.Val, nil, nil}
			}
			m[cur].Sibling = m[cur.Sibling]
		}
		cur = cur.Next
	}
	return m[head]
}

func printList(head *ComplexListNode) {
	for head != nil {
		fmt.Printf("当前节点: %d 兄弟节点: %d \t", head.Val, head.Sibling.Val)
		head = head.Next
	}
}

func main() {
	//test
	l3 := &ComplexListNode{3, nil, nil}
	l2 := &ComplexListNode{2, l3, nil}
	l1 := &ComplexListNode{1, l2, nil}
	l1.Sibling = l3
	l2.Sibling = l2
	l3.Sibling = l1
	fmt.Println("复制前 --------- ")
	printList(l1)
	fmt.Println()
	fmt.Println("复制后 --------- ")
	printList(CloneComplexList(l1))
	fmt.Println("")
}
